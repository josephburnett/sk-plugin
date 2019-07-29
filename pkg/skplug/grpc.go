package skplug

import (
	"context"
	"fmt"

	"github.com/hashicorp/go-plugin"
	"github.com/josephburnett/sk-plugin/pkg/skplug/proto"
	"google.golang.org/grpc"
)

var _ Autoscaler = &GRPCClient{}

// GRPCClient is an implementation of Autoscaler that talks over RPC.
type GRPCClient struct {
	broker *plugin.GRPCBroker
	client proto.AutoscalerClient
}

func (m *GRPCClient) Create(yaml string, c Cluster) (key string, err error) {
	clusterServer := &GRPCClusterServer{Impl: c}

	var s *grpc.Server
	serverFunc := func(opts []grpc.ServerOption) *grpc.Server {
		s = grpc.NewServer(opts...)
		proto.RegisterClusterServer(s, clusterServer)

		return s
	}

	brokerID := m.broker.NextId()
	go m.broker.AcceptAndServe(brokerID, serverFunc)
	defer s.Stop()

	resp, err := m.client.Create(context.Background(), &proto.CreateRequest{
		Yaml:          yaml,
		ClusterServer: brokerID,
	})
	if err != nil {
		return "", err
	}
	fmt.Printf("+%v", resp)
	if resp.Err != "" {
		// TODO: why do I get a nil point dereference here?
		return "", fmt.Errorf(resp.Err)
	}
	return resp.Key, nil
}

func (m *GRPCClient) Scale(key string) (rec int32, err error) {
	resp, err := m.client.Scale(context.Background(), &proto.ScaleRequest{
		Key: key,
	})
	if err != nil {
		return 0, err
	}
	if resp.Err != "" {
		err = fmt.Errorf(resp.Err)
		return
	}
	return resp.Rec, nil
}

func (m *GRPCClient) Stat(stats []*Stat) error {
	protoStats := make([]*proto.Stat, len(stats))
	for i, s := range stats {
		protoStats[i] = &proto.Stat{
			Time:    s.Time,
			PodName: s.PodName,
			Metric:  s.Metric,
			Value:   s.Value,
		}
	}
	resp, err := m.client.Stat(context.Background(), &proto.StatRequest{
		Key:  "",
		Stat: protoStats,
	})
	if err != nil {
		return err
	}
	if resp.Err != "" {
		return fmt.Errorf(resp.Err)
	}
	return nil
}

func (m *GRPCClient) Delete(key string) error {
	resp, err := m.client.Delete(context.Background(), &proto.DeleteRequest{
		Key: key,
	})
	if err != nil {
		return err
	}
	if resp.Err != "" {
		return fmt.Errorf(resp.Err)
	}
	return nil
}

// GRPCServer is the gRPC server that the GRPCClient talks to.
type GRPCServer struct {
	// This is the real implementation
	Impl Autoscaler

	broker *plugin.GRPCBroker
}

func (m *GRPCServer) Create(ctx context.Context, req *proto.CreateRequest) (*proto.CreateResponse, error) {
	conn, err := m.broker.Dial(req.ClusterServer)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	c := &GRPCClusterClient{proto.NewClusterClient(conn)}
	key, err := m.Impl.Create(req.Yaml, c)
	if err != nil {
		return &proto.CreateResponse{
			Err: err.Error(),
		}, nil
	}
	return &proto.CreateResponse{
		Key: key,
	}, nil
}

func (m *GRPCServer) Scale(ctx context.Context, req *proto.ScaleRequest) (*proto.ScaleResponse, error) {
	rec, err := m.Impl.Scale(req.Key)
	if err != nil {
		return &proto.ScaleResponse{
			Err: err.Error(),
		}, nil
	}
	return &proto.ScaleResponse{
		Rec: rec,
	}, nil
}

func (m *GRPCServer) Stat(ctx context.Context, req *proto.StatRequest) (*proto.StatResponse, error) {
	stats := make([]*Stat, len(req.Stat))
	for i, s := range req.Stat {
		stats[i] = &Stat{
			Time:    s.Time,
			PodName: s.PodName,
			Metric:  s.Metric,
			Value:   s.Value,
		}
	}
	err := m.Impl.Stat(stats)
	if err != nil {
		return &proto.StatResponse{
			Err: err.Error(),
		}, nil
	}
	return &proto.StatResponse{}, nil
}

func (m *GRPCServer) Delete(ctx context.Context, req *proto.DeleteRequest) (*proto.DeleteResponse, error) {
	err := m.Impl.Delete(req.Key)
	if err != nil {
		return &proto.DeleteResponse{
			Err: err.Error(),
		}, nil
	}
	return &proto.DeleteResponse{}, nil
}

// GRPCClusterClient is an implementation of Cluster that talks over RPC.
type GRPCClusterClient struct {
	client proto.ClusterClient
}

func (m *GRPCClusterClient) ListPods() (pods []*Pod, err error) {
	resp, err := m.client.ListPods(context.Background(), &proto.ListPodsRequest{})
	if err != nil {
		return nil, err
	}
	if resp.Err != "" {
		return nil, fmt.Errorf(resp.Err)
	}
	pods = make([]*Pod, len(resp.Pod))
	for i, p := range resp.Pod {
		pods[i] = &Pod{
			Name:           p.Name,
			State:          p.State,
			LastTransition: p.LastTransition,
			CpuRequest:     p.CpuRequest,
		}
	}
	return pods, nil
}

// GRPCClusterServer is the gRPC server that GRPCClusterClient talks to.
type GRPCClusterServer struct {
	// This is the real implementation
	Impl Cluster
}

func (m *GRPCClusterServer) ListPods(ctx context.Context, req *proto.ListPodsRequest) (*proto.ListPodsResponse, error) {
	pods, err := m.Impl.ListPods()
	if err != nil {
		return &proto.ListPodsResponse{
			Err: err.Error(),
		}, nil
	}
	protoPods := make([]*proto.Pod, len(pods))
	for i, p := range pods {
		protoPods[i] = &proto.Pod{
			Name:           p.Name,
			State:          p.State,
			LastTransition: p.LastTransition,
			CpuRequest:     p.CpuRequest,
		}
	}
	return &proto.ListPodsResponse{
		Pod: protoPods,
	}, nil
}
