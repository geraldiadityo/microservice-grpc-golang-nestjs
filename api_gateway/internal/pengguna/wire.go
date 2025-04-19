package pengguna

import (
	"api_gateway/internal/client"
	"api_gateway/proto/pengguna"

	"github.com/google/wire"
)

func ProvidePenggunaClient(grpcClient *client.GrpcClient) pengguna.PenggunaServiceClient {
	return grpcClient.PenggunaClient()
}

var ProviderSet = wire.NewSet(
	ProvidePenggunaClient,
	NewRepositoryPengguna,
	NewHandlerPengguna,
	wire.Bind(new(PenggunaRepository), new(*RepositoryPenggunaImpl)),
)
