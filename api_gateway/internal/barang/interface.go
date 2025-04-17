package barang

import "api_gateway/proto/barang"

type RepositoryBarang interface {
	Create(name string, categoryId int32) (*barang.BarangResponse, error)
	GetAll() (*barang.GetBarangsResponse, error)
	GetByID(id int32) (*barang.BarangResponse, error)
	Update(id int32, name string, categoryId int32) (*barang.BarangResponse, error)
	Delete(id int32) (*barang.DeleteBarangResponse, error)
	Close()
}
