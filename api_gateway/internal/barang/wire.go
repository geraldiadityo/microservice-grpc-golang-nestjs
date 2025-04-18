package barang

import (
	"api_gateway/internal/client"
	"api_gateway/proto/barang"

	"github.com/google/wire"
)

func ProvideBarangClient(grpcClient *client.GrpcClient) barang.BarangServiceClient {
	return grpcClient.BarangClient()
}

var ProviderSet = wire.NewSet(
	ProvideBarangClient,
	NewRepositoryBarang,
	NewHandlerBarang,
	wire.Bind(new(RepositoryBarang), new(*RepositoryBarangImpl)),
)
