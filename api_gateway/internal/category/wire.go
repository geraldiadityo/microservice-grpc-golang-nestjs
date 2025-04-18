package category

import (
	"api_gateway/internal/client"
	"api_gateway/proto/category"

	"github.com/google/wire"
)

func ProvideCategoryClient(grpcClient *client.GrpcClient) category.CategoryServiceClient {
	return grpcClient.CategoryClient()
}

var ProviderSet = wire.NewSet(
	ProvideCategoryClient,
	NewRepositoryCategory,
	NewHandlerCategory,
	wire.Bind(new(RepositoryCategory), new(*RepositoryCategoryImpl)),
)
