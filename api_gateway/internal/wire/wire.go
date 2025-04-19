//go:build wireinject
// +build wireinject

package wire

import (
	"api_gateway/config"
	"api_gateway/internal/barang"
	"api_gateway/internal/category"
	"api_gateway/internal/client"
	"api_gateway/internal/pengguna"
	"api_gateway/internal/role"
	"api_gateway/internal/server"

	"github.com/google/wire"
)

func InitializeServer() (*ServerWithCleanup, error) {
	wire.Build(
		// provider
		config.ProvideConfig,

		// grpc client
		client.ProvideSet,

		// service
		category.ProviderSet,
		barang.ProviderSet,
		role.ProviderSet,
		pengguna.ProviderSet,

		// server
		wire.Struct(new(server.Handlers), "*"),
		server.NewServer,
		newServerWithCleanup,
	)

	return &ServerWithCleanup{}, nil
}
