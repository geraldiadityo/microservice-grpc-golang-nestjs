package role

import "api_gateway/proto/role"

type RepositoryRole interface {
	Create(name string) (*role.RoleResponse, error)
	GetAll() (*role.RolesResponse, error)
	GetByID(id int32) (*role.RoleResponse, error)
	Update(id int32, name string) (*role.RoleResponse, error)
	Delete(id int32) (*role.DeleteRoleResponse, error)
}
