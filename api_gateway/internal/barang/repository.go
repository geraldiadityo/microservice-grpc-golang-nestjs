package barang

import (
	"api_gateway/proto/barang"
	"context"
)

type RepositoryBarangImpl struct {
	client barang.BarangServiceClient
}

func NewRepositoryBarang(client barang.BarangServiceClient) *RepositoryBarangImpl {
	return &RepositoryBarangImpl{
		client: client,
	}
}

func (r *RepositoryBarangImpl) Create(name string, categoryId int32) (*barang.BarangResponse, error) {
	req := &barang.CreateBarangRequest{
		Name:       name,
		CategoryId: categoryId,
	}

	return r.client.CreateBarang(context.Background(), req)
}

func (r *RepositoryBarangImpl) GetAll() (*barang.GetBarangsResponse, error) {
	req := &barang.GetBarangsRequest{}

	return r.client.GetBarangs(context.Background(), req)
}

func (r *RepositoryBarangImpl) GetByID(id int32) (*barang.BarangResponse, error) {
	req := &barang.GetBarangRequest{
		Id: id,
	}

	return r.client.GetBarang(context.Background(), req)
}

func (r *RepositoryBarangImpl) Update(id int32, name string, categoryId int32) (*barang.BarangResponse, error) {
	req := &barang.UpdateBarangRequest{
		Id:         id,
		Name:       name,
		CategoryId: categoryId,
	}

	return r.client.UpdateBarang(context.Background(), req)
}

func (r *RepositoryBarangImpl) Delete(id int32) (*barang.DeleteBarangResponse, error) {
	req := &barang.DeleteBarangRequest{
		Id: id,
	}

	return r.client.DeleteBarang(context.Background(), req)
}

// func (r *RepositoryBarangImpl) Close() {
// 	r.client.Close()
// }
