package category

import "github.com/google/wire"

var ProviderSet = wire.NewSet(
	NewClientCategory,
	NewRepositoryCategory,
	NewHandlerCategory,
	wire.Bind(new(RepositoryCategory), new(*RepositoryCategoryImpl)),
)
