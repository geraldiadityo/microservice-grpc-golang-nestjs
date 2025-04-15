import { Injectable } from "@nestjs/common";
import { ICategoryRepository } from "./interface/category_repo.interface";
import { PrismaService } from "src/common/prisma.service";
import { category } from "@prisma/client";

@Injectable()
export class CategoryRepository implements ICategoryRepository {
    constructor(
        private prismaService: PrismaService
    ) {}

    async create(data: { name: string; }): Promise<category> {
        return await this.prismaService.category.create({
            data: data
        });
    }
    
    async findAll(): Promise<category[]> {
        return await this.prismaService.category.findMany();
    }

    async findOne(id: number): Promise<category | null> {
        return await this.prismaService.category.findUnique({
            where: {
                id: id
            }
        });
    }

    async update(id: number, data: { name: string; }): Promise<category> {
        return await this.prismaService.category.update({
            where: {
                id: id
            },
            data: data
        })
    };

    async delete(id: number): Promise<void> {
        await this.prismaService.category.delete({
            where: {
                id: id
            }
        })
    }
}