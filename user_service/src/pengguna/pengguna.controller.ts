import { Controller } from "@nestjs/common";
import { PenggunaService } from "./pengguna.service";
import { GrpcMethod } from "@nestjs/microservices";
import { CreatePenggunaRequest, DeletePenggunaRequest, DeletePenggunaResponse, GetPenggunaRequest, GetPenggunasRequest, GetPenggunasResponse, PenggunaResponse, PenggunaServiceController, UpdatePenggunaRequest } from "src/generated/pengguna";


@Controller()
export class PenggunaController implements PenggunaServiceController {
    constructor(
        private service: PenggunaService
    ) {}

    @GrpcMethod('PenggunaService', 'CreatePengguna')
    async createPengguna(request: CreatePenggunaRequest): Promise<PenggunaResponse> {
        return await this.service.create(request);
    }

    @GrpcMethod('PenggunaService', 'GetPenggunas')
    async getPenggunas(request: GetPenggunasRequest): Promise<GetPenggunasResponse> {
        return await this.service.getAll();
    }

    @GrpcMethod('PenggunaService', 'GetPengguna')
    async getPengguna(request: GetPenggunaRequest): Promise<PenggunaResponse> {
        return await this.service.getOne(request.id)
    }

    @GrpcMethod('PenggunaService', 'UpdatePengguna')
    async updatePengguna(request: UpdatePenggunaRequest): Promise<PenggunaResponse> {
        const data = {
            username: request.username,
            name: request.name,
            password: request.password,
            roleId: request.roleId
        }

        return await this.service.update(request.id, data);
    }

    @GrpcMethod('PenggunaService', 'DeletePengguna')
    async deletePengguna(request: DeletePenggunaRequest): Promise<DeletePenggunaResponse> {
        return await this.service.remove(request.id)
    }
}