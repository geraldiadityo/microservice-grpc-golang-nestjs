import { Injectable } from "@nestjs/common";
import { PenggunaRepository } from "./pengguna.repository";
import { CreatePenggunaRequest, GetPenggunasResponse, PenggunaResponse } from "src/generated/pengguna";
import * as bcrypt from 'bcrypt';

@Injectable()
export class PenggunaService {
    constructor(
        private readonly repo: PenggunaRepository
    ) {}

    private toPenggunaResponse(pengguna: any): PenggunaResponse {
        return {
            id: pengguna.id,
            username: pengguna.username,
            name: pengguna.name,
            role: {
                id: pengguna.role.id,
                name: pengguna.role.name
            }
        }
    }

    async create(
        data: { username: string, name: string, password: string, roleId: number}
    ): Promise<PenggunaResponse> {
        const request: CreatePenggunaRequest = {
            username: data.username,
            name: data.name,
            password: await bcrypt.hash(data.password, 10),
            roleId: data.roleId
        }

        const pengguna = await this.repo.create(request);

        return this.toPenggunaResponse(pengguna)
    }

    async getAll(): Promise<GetPenggunasResponse> {
        const listPengguna = await this.repo.findAll();

        return {
            penggunas: listPengguna.map((item) => this.toPenggunaResponse(item))
        }
    }

    async getOne(
        id: number
    ): Promise<PenggunaResponse> {
        const pengguna = await this.repo.findOne(id);
        if (!pengguna) {
            throw new Error('Pengguna not found!')
        }

        return this.toPenggunaResponse(pengguna)
    }

    async update(
        id: number,
        data: { username: string, name: string, password: string, roleId: number}
    ): Promise<PenggunaResponse> {
        const request: CreatePenggunaRequest = {
            username: data.username,
            name: data.name,
            password: await bcrypt.hash(data.password, 10),
            roleId: data.roleId
        }

        const pengguna = await this.repo.update(id, request);

        return this.toPenggunaResponse(pengguna);
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