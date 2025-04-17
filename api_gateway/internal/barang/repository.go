package barang

import (
	"api_gateway/proto/barang"
	"context"
)

type RepositoryBarangImpl struct {
	client *Client
}

func NewRepositoryBarang(client *Client) *RepositoryBarangImpl {
	return &RepositoryBarangImpl{
		client: client,
	}
}

func (r *RepositoryBarangImpl) Create(name string, categoryId int32) (*barang.BarangResponse, error) {
	req := &barang.CreateBarangRequest{
		Name:       name,
		CategoryId: categoryId,
	}

	return r.client.service.CreateBarang(context.Background(), req)
}

func (r *RepositoryBarangImpl) GetAll() (*barang.GetBarangsResponse, error) {
	req := &barang.GetBarangsRequest{}

	return r.client.service.GetBarangs(context.Background(), req)
}

func (r *RepositoryBarangImpl) GetByID(id int32) (*barang.BarangResponse, error) {
	req := &barang.GetBarangRequest{
		Id: id,
	}

	return r.client.service.GetBarang(context.Background(), req)
}

func (r *RepositoryBarangImpl) Update(id int32, name string, categoryId int32) (*barang.BarangResponse, error) {
	req := &barang.UpdateBarangRequest{
		Id:         id,
		Name:       name,
		CategoryId: categoryId,
	}

	return r.client.service.UpdateBarang(context.Background(), req)
}

func (r *RepositoryBarangImpl) Delete(id int32) (*barang.DeleteBarangResponse, error) {
	req := &barang.DeleteBarangRequest{
		Id: id,
	}

	return r.client.service.DeleteBarang(context.Background(), req)
}

func (r *RepositoryBarangImpl) Close() {
	r.client.Close()
}
