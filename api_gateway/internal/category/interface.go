package category

import "api_gateway/proto/category"

type RepositoryCategory interface {
	Create(name string) (*category.CategoryResponse, error)
	GetAll() (*category.GetCategoriesResponse, error)
	GetByID(id int32) (*category.CategoryResponse, error)
	Update(id int32, name string) (*category.CategoryResponse, error)
	Delete(id int32) (*category.DeleteCategoryResponse, error)
}
