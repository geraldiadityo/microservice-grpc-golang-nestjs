import { Injectable } from "@nestjs/common";
import { CategoryRepository } from "./category.repository";
import { category } from "@prisma/client";
import { CategoryResponse, GetCategoriesResponse } from "src/generated/category";

@Injectable()
export class CategoryService {
    constructor(
        private readonly repo: CategoryRepository
    ) {}

    private toCategoryResponse(category: category): CategoryResponse {
        return {
            id: category.id,
            name: category.name
        }
    }

    async create(
        data: { name: string }
    ): Promise<CategoryResponse> {
        const category = await this.repo.create(data);

        return this.toCategoryResponse(category);
    }

    async getAll(): Promise<GetCategoriesResponse> {
        const listCategory = await this.repo.findAll();

        return {
            categories: listCategory.map((item) => this.toCategoryResponse(item))
        }
    }

    async getOne(
        id: number
    ): Promise<CategoryResponse> {
        const category = await this.repo.findOne(id);
        if (!category){
            throw new Error('Category not found!')
        }

        return this.toCategoryResponse(category);
    }

    async update(
        id: number,
        data: { name: string }
    ): Promise<CategoryResponse> {
        const category = await this.repo.update(id, data);

        return this.toCategoryResponse(category);
    }

    async remove(
        id: number
    ): Promise<{ success: boolean }> {
        await this.repo.delete(id);

        return {
            success: true
        }
    }
}