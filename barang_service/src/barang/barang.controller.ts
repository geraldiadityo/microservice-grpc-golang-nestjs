import { Controller, Injectable } from "@nestjs/common";
import { BarangResponse, BarangServiceController, CreateBarangRequest, DeleteBarangRequest, DeleteBarangResponse, GetBarangRequest, GetBarangsRequest, GetBarangsResponse, UpdateBarangRequest } from "src/generated/barang";
import { BarangService } from "./barang.service";
import { GrpcMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";

@Controller()
export class BarangController implements BarangServiceController {
    constructor(
        private service: BarangService
    ) {}

    @GrpcMethod('BarangService', 'CreateBarang')
    async createBarang(request: CreateBarangRequest): Promise<BarangResponse> {
        const data = {
            name: request.name,
            categoryId: request.categoryId
        }

        return await this.service.create(data)
    }

    @GrpcMethod('BarangService', 'GetBarangs')
    async getBarangs(request: GetBarangsRequest): Promise<GetBarangsResponse> {
        return this.service.getAll()
    }

    @GrpcMethod('BarangService', 'GetBarang')
    async getBarang(request: GetBarangRequest): Promise<BarangResponse> {
        return this.service.getOne(request.id)
    }

    @GrpcMethod('BarangService', 'UpdateBarang')
    async updateBarang(request: UpdateBarangRequest): Promise<BarangResponse>  {
        const data = {
            name: request.name,
            categoryId: request.categoryId
        }

        return this.service.update(request.id, data)
    }

    @GrpcMethod('BarangService', 'DeleteBarang')
    async deleteBarang(request: DeleteBarangRequest): Promise<DeleteBarangResponse>  {
        return this.service.remove(request.id)
    }
}