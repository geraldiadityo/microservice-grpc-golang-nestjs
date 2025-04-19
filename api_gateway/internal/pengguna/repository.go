package pengguna

import (
	"api_gateway/proto/pengguna"
	"context"
)

type RepositoryPenggunaImpl struct {
	client pengguna.PenggunaServiceClient
}

func NewRepositoryPengguna(client pengguna.PenggunaServiceClient) *RepositoryPenggunaImpl {
	return &RepositoryPenggunaImpl{
		client: client,
	}
}

func (r *RepositoryPenggunaImpl) Create(username string, name string, password string, roleId int32) (*pengguna.PenggunaResponse, error) {
	req := &pengguna.CreatePenggunaRequest{
		Username: username,
		Name:     name,
		Password: password,
		RoleId:   roleId,
	}

	return r.client.CreatePengguna(context.Background(), req)
}

func (r *RepositoryPenggunaImpl) GetAll() (*pengguna.GetPenggunasResponse, error) {
	req := &pengguna.GetPenggunasRequest{}

	return r.client.GetPenggunas(context.Background(), req)
}

func (r *RepositoryPenggunaImpl) GetByID(id int32) (*pengguna.PenggunaResponse, error) {
	req := &pengguna.GetPenggunaRequest{
		Id: id,
	}

	return r.client.GetPengguna(context.Background(), req)
}

func (r *RepositoryPenggunaImpl) Update(id int32, username string, name string, password string, roleId int32) (*pengguna.PenggunaResponse, error) {
	req := &pengguna.UpdatePenggunaRequest{
		Id:       id,
		Username: username,
		Name:     name,
		Password: password,
		RoleId:   roleId,
	}

	return r.client.UpdatePengguna(context.Background(), req)
}

func (r *RepositoryPenggunaImpl) Delete(id int32) (*pengguna.DeletePenggunaResponse, error) {
	req := &pengguna.DeletePenggunaRequest{
		Id: id,
	}

	return r.client.DeletePengguna(context.Background(), req)
}
