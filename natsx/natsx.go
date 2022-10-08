package natsx

import (
	"context"

	"github.com/junaozun/gogopkg/config"
	"github.com/junaozun/gogopkg/logrusx"
)

type NatsxServer struct {
	ServerName string
	*Server
}

func New(natsCfg *config.NatsConfig, serverName string) *NatsxServer {
	connEnc, err := NewNatsPBEnc(natsCfg.Server) // nats.MaxReconnects(int(natsCfg.MaxReconnects)),
	// nats.ReconnectWait(time.Duration(natsCfg.ReconnectWait)),
	// nats.Timeout(time.Duration(natsCfg.RequestTimeout)),

	if err != nil {
		panic(err)
	}
	server, err := NewServer(connEnc)
	if err != nil {
		panic(err)
	}
	return &NatsxServer{
		Server:     server,
		ServerName: serverName,
	}
}

func (n *NatsxServer) Start(ctx context.Context) error {
	logrusx.Log.WithFields(logrusx.Fields{
		"serverName": n.ServerName,
	}).Info("[NatsxServer] Start success......")
	select {
	case <-ctx.Done():
		return nil
	}
	return nil
}

func (n *NatsxServer) Stop(ctx context.Context) error {
	return n.Close(ctx)
}
