package pengguna

import "api_gateway/proto/pengguna"

type PenggunaRepository interface {
	Create(username string, name string, password string, roleId int32) (*pengguna.PenggunaResponse, error)
	GetAll() (*pengguna.GetPenggunasResponse, error)
	GetByID(id int32) (*pengguna.PenggunaResponse, error)
	Update(id int32, username string, name string, password string, roleId int32) (*pengguna.PenggunaResponse, error)
	Delete(id int32) (*pengguna.DeletePenggunaResponse, error)
}
