package barang

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewClientBarang,
	NewRepositoryBarang,
	NewHandlerBarang,
	wire.Bind(new(RepositoryBarang), new(*RepositoryBarangImpl)),
)
