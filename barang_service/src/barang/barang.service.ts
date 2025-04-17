import { Injectable } from "@nestjs/common";
import { BarangRepository } from "./barang.repository";
import { barang } from "@prisma/client";
import { BarangResponse, GetBarangsResponse } from "src/generated/barang";

@Injectable()
export class BarangService {
    constructor(
        private readonly repo: BarangRepository
    ) {}

    private toBarangResponse(barang: any): BarangResponse {
        return {
            id: barang.id,
            name: barang.name,
            category: {
                id: barang.category.id,
                name: barang.category.name
            }
        }
    }

    async create(
        data: { name: string, categoryId: number }
    ): Promise<BarangResponse> {
        const barang = await this.repo.create(data);

        return this.toBarangResponse(barang);
    }

    async getAll(): Promise<GetBarangsResponse> {
        const listBarang = await this.repo.findAll()
        
        return {
            barangs: listBarang.map((item) => this.toBarangResponse(item))
        }
    }

    async getOne(
        id: number
    ): Promise<BarangResponse> {
        const barang = await this.repo.findOne(id)
        if (!barang){
            throw new Error('Barang not found!')
        }

        return this.toBarangResponse(barang)
    }

    async update(
        id: number,
        data: { name: string, categoryId: number }
    ): Promise<BarangResponse> {
        const barang = await this.repo.update(id, data);

        return this.toBarangResponse(barang);
    }

    async remove(
        id: number
    ): Promise<{ success: boolean }> {
        await this.repo.delete(id)

        return {
            success: true
        }
    }
}