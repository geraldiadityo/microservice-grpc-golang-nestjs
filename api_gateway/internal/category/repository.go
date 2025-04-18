package category

import (
	"api_gateway/proto/category"
	"context"
)

type RepositoryCategoryImpl struct {
	client category.CategoryServiceClient
}

func NewRepositoryCategory(client category.CategoryServiceClient) *RepositoryCategoryImpl {
	return &RepositoryCategoryImpl{client: client}
}

func (repo *RepositoryCategoryImpl) Create(name string) (*category.CategoryResponse, error) {
	req := &category.CreateCategoryRequest{
		Name: name,
	}

	return repo.client.CreateCategory(context.Background(), req)
}

func (repo *RepositoryCategoryImpl) GetAll() (*category.GetCategoriesResponse, error) {
	req := &category.GetCategoriesRequest{}

	return repo.client.GetCategories(context.Background(), req)
}

func (repo *RepositoryCategoryImpl) GetByID(id int32) (*category.CategoryResponse, error) {
	req := &category.GetCategoryRequest{
		Id: id,
	}

	return repo.client.GetCategory(context.Background(), req)
}

func (repo *RepositoryCategoryImpl) Update(id int32, name string) (*category.CategoryResponse, error) {
	req := &category.UpdateCategoryRequest{
		Id:   id,
		Name: name,
	}

	return repo.client.UpdateCategory(context.Background(), req)
}

func (repo *RepositoryCategoryImpl) Delete(id int32) (*category.DeleteCategoryResponse, error) {
	req := &category.DeleteCategoryRequest{
		Id: id,
	}

	return repo.client.DeleteCategory(context.Background(), req)
}
