//go:build wireinject
// +build wireinject

package wire

import (
	"api_gateway/config"
	"api_gateway/internal/barang"
	"api_gateway/internal/category"
	"api_gateway/internal/server"

	"github.com/google/wire"
)

func InitializeServer() (*server.Server, error) {
	wire.Build(
		// provider
		config.ProvideConfig,

		// category
		category.ProviderSet,

		// barang
		barang.ProviderSet,

		// server
		server.NewServer,
		wire.Struct(new(server.Handlers), "*"),
	)

	return &server.Server{}, nil
}
