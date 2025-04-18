import { Injectable } from "@nestjs/common";
import { RoleRepository } from "./role.repository";
import { Role } from "@prisma/client";
import { RoleResponse, RolesResponse } from "src/generated/role";

@Injectable()
export class RoleService {
    constructor(
        private readonly repo: RoleRepository
    ) {}

    private toRoleResponse(role: Role): RoleResponse {
        return {
            id: role.id,
            name: role.name
        }
    }

    async create(
        data: { name: string }
    ): Promise<RoleResponse> {
        const role = await this.repo.create(data);

        return this.toRoleResponse(role)
    }

    async getAll(): Promise<RolesResponse> {
        const listRoles = await this.repo.findAll();
        
        return {
            roles: listRoles.map((item) => this.toRoleResponse(item))
        }
    }

    async getOne(
        id: number
    ): Promise<RoleResponse> {
        const role = await this.repo.findOne(id);
        if(!role){
            throw new Error('Role not found!')
        }

        return this.toRoleResponse(role)
    }

    async update(
        id: number,
        data: { name: string }
    ): Promise<RoleResponse> {
        const role = await this.repo.update(id, data);

        return this.toRoleResponse(role)
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