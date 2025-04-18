import { CreateRoleRequest, DeleteRoleRequest, DeleteRoleResponse, GetRoleRequest, GetRolesRequest, RoleResponse, RoleServiceController, RolesResponse, UpdateRoleRequest } from "src/generated/role";
import { RoleService } from "./role.service";
import { GrpcMethod } from "@nestjs/microservices";
import { Observable } from "rxjs";
import { Controller } from "@nestjs/common";

@Controller()
export class RoleController implements RoleServiceController {
    constructor(
        private service: RoleService
    ) {}

    @GrpcMethod('RoleService', 'CreateRole')
    async createRole(request: CreateRoleRequest): Promise<RoleResponse>  {
        return await this.service.create(request)
    }

    @GrpcMethod('RoleService', 'GetRoles')
    async getRoles(request: GetRolesRequest): Promise<RolesResponse> {
        return await this.service.getAll();
    }

    @GrpcMethod('RoleService', 'GetRole')
    async getRole(request: GetRoleRequest): Promise<RoleResponse> {
        return await this.service.getOne(request.id)
    }

    @GrpcMethod('RoleService', 'UpdateRole')
    async updateRole(request: UpdateRoleRequest): Promise<RoleResponse> {
        const data = {
            name: request.name
        }

        return await this.service.update(request.id, data)
    }

    @GrpcMethod('RoleService', 'DeleteRole')
    async deleteRole(request: DeleteRoleRequest): Promise<DeleteRoleResponse> {
        return await this.service.remove(request.id)
    }
}