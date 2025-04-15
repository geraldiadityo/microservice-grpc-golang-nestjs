import { category } from "@prisma/client";

export interface ICategoryRepository {
    create(data: { name: string }): Promise<category>;
    findAll(): Promise<category[]>;
    findOne(id: number): Promise<category | null>;
    update(id: number, data: { name: string }): Promise<category>;
    delete(id: number): Promise<void>;
}