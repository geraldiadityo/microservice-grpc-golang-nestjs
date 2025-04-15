import { Controller } from "@nestjs/common";
import { CategoryResponse, CategoryServiceController, CreateCategoryRequest, DeleteCategoryRequest, DeleteCategoryResponse, GetCategoriesRequest, GetCategoriesResponse, GetCategoryRequest, UpdateCategoryRequest } from "src/generated/category";
import { CategoryService } from "./category.service";
import { GrpcMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";

@Controller()
export class CategoryController implements CategoryServiceController {
    constructor(
        private service: CategoryService
    ) {}

    @GrpcMethod('CategoryService', 'CreateCategory')
    async createCategory(request: CreateCategoryRequest): Promise<CategoryResponse>  {
        return await this.service.create(request);
    }

    @GrpcMethod('CategoryService', 'GetCategories')
    async getCategories(request: GetCategoriesRequest): Promise<GetCategoriesResponse>  {
        return await this.service.getAll();
    }

    @GrpcMethod('CategoryService', 'GetCategory')
    async getCategory(request: GetCategoryRequest): Promise<CategoryResponse>  {
        return await this.service.getOne(request.id)
    }

    @GrpcMethod('CategoryService', 'UpdateCategory')
    async updateCategory(request: UpdateCategoryRequest): Promise<CategoryResponse>  {
        const data = {
            name: request.name
        }
        return await this.service.update(request.id, data)
    }

    @GrpcMethod('CategoryService', 'DeleteCategory')
    async deleteCategory(request: DeleteCategoryRequest): Promise<DeleteCategoryResponse>  {
        return await this.service.remove(request.id)
    }
}