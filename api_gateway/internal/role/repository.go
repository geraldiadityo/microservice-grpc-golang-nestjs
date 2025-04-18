package role

import (
	"api_gateway/proto/role"
	"context"
)

type RepositoryRoleImpl struct {
	client role.RoleServiceClient
}

func NewRepositoryRole(client role.RoleServiceClient) *RepositoryRoleImpl {
	return &RepositoryRoleImpl{
		client: client,
	}
}

func (r *RepositoryRoleImpl) Create(name string) (*role.RoleResponse, error) {
	req := &role.CreateRoleRequest{
		Name: name,
	}

	return r.client.CreateRole(context.Background(), req)
}

func (r *RepositoryRoleImpl) GetAll() (*role.RolesResponse, error) {
	req := &role.GetRolesRequest{}

	return r.client.GetRoles(context.Background(), req)
}

func (r *RepositoryRoleImpl) GetByID(id int32) (*role.RoleResponse, error) {
	req := &role.GetRoleRequest{
		Id: id,
	}

	return r.client.GetRole(context.Background(), req)
}

func (r *RepositoryRoleImpl) Update(id int32, name string) (*role.RoleResponse, error) {
	req := &role.UpdateRoleRequest{
		Id:   id,
		Name: name,
	}

	return r.client.UpdateRole(context.Background(), req)
}

func (r *RepositoryRoleImpl) Delete(id int32) (*role.DeleteRoleResponse, error) {
	req := &role.DeleteRoleRequest{
		Id: id,
	}

	return r.client.DeleteRole(context.Background(), req)
}
