import { Injectable } from "@nestjs/common";
import { IRoleRepository } from "./interface/role_repo.interface";
import { PrismaService } from "src/common/prisma.service";
import { Role } from "@prisma/client";

@Injectable()
export class RoleRepository implements IRoleRepository {
    constructor(
        private prismaService: PrismaService
    ) {}

    async create(data: { name: string; }): Promise<Role> {
        return await this.prismaService.role.create({
            data: data
        });
    }

    async findAll(): Promise<Role[]> {
        return await this.prismaService.role.findMany();
    }

    async findOne(id: number): Promise<Role | null> {
        return await this.prismaService.role.findUnique({ where: { id } });
    }

    async update(id: number, data: { name: string; }): Promise<Role> {
        return await this.prismaService.role.update({
            where: {
                id: id
            },
            data: data
        })
    }

    async delete(id: number): Promise<void> {
        await this.prismaService.role.delete({
            where: {
                id: id
            }
        })
    }
}