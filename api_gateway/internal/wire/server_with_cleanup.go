package wire

import (
	"api_gateway/internal/client"
	"api_gateway/internal/server"
)

type ServerWithCleanup struct {
	Server      *server.Server
	CleanupFunc func()
}

func newServerWithCleanup(s *server.Server, grpcClient *client.GrpcClient) *ServerWithCleanup {
	return &ServerWithCleanup{
		Server:      s,
		CleanupFunc: grpcClient.Close,
	}
}
