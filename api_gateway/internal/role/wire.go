package role

import (
	"api_gateway/internal/client"
	"api_gateway/proto/role"

	"github.com/google/wire"
)

func ProvideRoleClient(grpcClient *client.GrpcClient) role.RoleServiceClient {
	return grpcClient.RoleClient()
}

var ProviderSet = wire.NewSet(
	ProvideRoleClient,
	NewRepositoryRole,
	NewHandlerRole,
	wire.Bind(new(RepositoryRole), new(*RepositoryRoleImpl)),
)
