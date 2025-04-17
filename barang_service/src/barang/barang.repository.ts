import { Injectable } from "@nestjs/common";
import { IBarangRepository } from "./interface/barang_repo.interface";
import { PrismaService } from "src/common/prisma.service";
import { barang } from "@prisma/client";

@Injectable()
export class BarangRepository implements IBarangRepository {
    constructor(
        private prismaService: PrismaService
    ) {}

    async create(data: { name: string; categoryId: number; }): Promise<barang> {
        return this.prismaService.barang.create({
            data: data,
            include: {
                category: true
            }
        });
    }

    async findAll(): Promise<barang[]> {
        return this.prismaService.barang.findMany({
            include: {
                category: true
            }
        })
    }

    async findOne(id: number): Promise<barang | null> {
        return this.prismaService.barang.findUnique({
            where: {
                id: id
            },
            include: {
                category: true
            }
        })
    }

    async update(id: number, data: { name: string; categoryId: number; }): Promise<barang> {
        return this.prismaService.barang.update({
            where: {
                id: id
            },
            data: data,
            include: {
                category: true
            }
        })
    }

    async delete(id: number): Promise<void> {
        await this.prismaService.barang.delete({
            where: {
                id: id
            }
        })
    }
}