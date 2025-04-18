import { Role } from "@prisma/client";

export interface IRoleRepository {
    create(data: { name: string }): Promise<Role>;
    findAll(): Promise<Role[]>;
    findOne(id: number): Promise<Role | null>;
    update(id: number, data: { name: string }): Promise<Role>;
    delete(id: number): Promise<void>;
}