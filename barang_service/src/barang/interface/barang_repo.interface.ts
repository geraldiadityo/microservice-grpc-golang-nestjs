import { barang } from "@prisma/client";

export interface IBarangRepository {
    create(data: { name: string, categoryId: number }): Promise<barang>;
    findAll(): Promise<barang[]>
    findOne(id: number): Promise<barang | null>;
    update(id: number, data: { name: string, categoryId: number }): Promise<barang>;
    delete(id: number): Promise<void>;
}