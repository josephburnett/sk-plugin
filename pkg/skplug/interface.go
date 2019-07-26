package skplug

import (
	"context"

	"github.com/hashicorp/go-plugin"
	"github.com/josephburnett/sk-plugin/pkg/skplug/proto"
	"google.golang.org/grpc"
)

// Handshake is a common handshake that is shared by plugin and host.
var Handshake = plugin.HandshakeConfig{
	ProtocolVersion:  1,
	MagicCookieKey:   "SKENARIO_PLUGIN",
	MagicCookieValue: "skenario",
}

// PluginMap is the map of plugins we can dispense.
var PluginMap = map[string]plugin.Plugin{
	"autoscaler": &AutoscalerPlugin{},
}

type Pod struct {
	Name           string
	State          string
	LastTransition int64
	CpuRequest     int32
}

type Cluster interface {
	ListPods() (pods []*Pod, err error)
}

type Stat struct {
	Time    int64
	PodName string
	Metric  string
	Value   int32
}

// Autoscaler is the interface that we're exposing as a plugin.
type Autoscaler interface {
	Create(yaml string, c Cluster) (key string, err error)
	Scale(key string) (rec int32, err error)
	Stat(stat []*Stat) error
	Delete(key string) error
}

// This is the implementation of plugin.Plugin so we can serve/consume this.
// We also implement GRPCPlugin so that this plugin can be served over
// gRPC.
type AutoscalerPlugin struct {
	plugin.NetRPCUnsupportedPlugin
	// Concrete implementation, written in Go. This is only used for plugins
	// that are written in Go.
	Impl Autoscaler
}

func (p *AutoscalerPlugin) GRPCServer(broker *plugin.GRPCBroker, s *grpc.Server) error {
	proto.RegisterAutoscalerServer(s, &GRPCServer{
		Impl:   p.Impl,
		broker: broker,
	})
	return nil
}

func (p *AutoscalerPlugin) GRPCClient(ctx context.Context, broker *plugin.GRPCBroker, c *grpc.ClientConn) (interface{}, error) {
	return &GRPCClient{
		client: proto.NewAutoscalerClient(c),
		broker: broker,
	}, nil
}

var _ plugin.GRPCPlugin = &AutoscalerPlugin{}
