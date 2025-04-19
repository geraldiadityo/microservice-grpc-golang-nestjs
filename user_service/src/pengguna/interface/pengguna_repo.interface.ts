import { Pengguna } from "@prisma/client";

export interface IPenggunaRepository {
    create(data: { username: string, name: string, password: string, roleId: number }): Promise<Pengguna>;
    findAll(): Promise<Pengguna[]>;
    findOne(id: number): Promise<Pengguna | null>;
    update(id: number, data: { username: string, name: string, password: string, roleId: number }): Promise<Pengguna>;
    delete(id: number): Promise<void>;
}