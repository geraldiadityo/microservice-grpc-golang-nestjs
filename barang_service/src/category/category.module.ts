import { Module } from "@nestjs/common";
import { CategoryRepository } from "./category.repository";
import { CategoryService } from "./category.service";
import { CategoryController } from "./category.controller";

@Module({
    providers: [CategoryRepository, CategoryService],
    controllers: [CategoryController]
})
export class CategoryModule {}