package category

import (
	"api_gateway/proto/category"
	"context"
)

type RepositoryCategoryImpl struct {
	client *Client
}

func NewRepositoryCategory(client *Client) *RepositoryCategoryImpl {
	return &RepositoryCategoryImpl{client: client}
}

func (repo *RepositoryCategoryImpl) Create(name string) (*category.CategoryResponse, error) {
	req := &category.CreateCategoryRequest{
		Name: name,
	}

	return repo.client.service.CreateCategory(context.Background(), req)
}

func (repo *RepositoryCategoryImpl) GetAll() (*category.GetCategoriesResponse, error) {
	req := &category.GetCategoriesRequest{}

	return repo.client.service.GetCategories(context.Background(), req)
}

func (repo *RepositoryCategoryImpl) GetByID(id int32) (*category.CategoryResponse, error) {
	req := &category.GetCategoryRequest{
		Id: id,
	}

	return repo.client.service.GetCategory(context.Background(), req)
}

func (repo *RepositoryCategoryImpl) Update(id int32, name string) (*category.CategoryResponse, error) {
	req := &category.UpdateCategoryRequest{
		Id:   id,
		Name: name,
	}

	return repo.client.service.UpdateCategory(context.Background(), req)
}

func (repo *RepositoryCategoryImpl) Delete(id int32) (*category.DeleteCategoryResponse, error) {
	req := &category.DeleteCategoryRequest{
		Id: id,
	}

	return repo.client.service.DeleteCategory(context.Background(), req)
}

func (repo *RepositoryCategoryImpl) Close() {
	repo.client.Close()
}
