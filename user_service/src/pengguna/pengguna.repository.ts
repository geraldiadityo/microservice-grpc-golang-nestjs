import { PrismaService } from "src/common/prisma.service";
import { IPenggunaRepository } from "./interface/pengguna_repo.interface";
import { Pengguna } from "@prisma/client";
import { Injectable } from "@nestjs/common";

@Injectable()
export class PenggunaRepository implements IPenggunaRepository {
    constructor(
        private prisma: PrismaService
    ) {}

    async create(data: { username: string; name: string; password: string; roleId: number; }): Promise<Pengguna> {
        return await this.prisma.pengguna.create({
            data: data,
            include: {
                role: true
            }
        });
    }

    async findAll(): Promise<Pengguna[]> {
        return await this.prisma.pengguna.findMany({
            include: {
                role: true
            }
        });
    }

    async findOne(id: number): Promise<Pengguna | null> {
        return await this.prisma.pengguna.findUnique({
            where: {
                id: id
            },
            include: {
                role: true
            }
        })
    }

    async update(id: number, data: { username: string; name: string; password: string; roleId: number; }): Promise<Pengguna> {
        return await this.prisma.pengguna.update({
            where: {
                id: id
            },
            data: data,
            include: {
                role: true
            }
        })
    }

    async delete(id: number): Promise<void> {
        await this.prisma.pengguna.delete({
            where: {
                id: id
            }
        })
    }
}