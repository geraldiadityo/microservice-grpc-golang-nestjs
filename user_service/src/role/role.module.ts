import { Module } from "@nestjs/common";
import { RoleRepository } from "./role.repository";
import { RoleService } from "./role.service";
import { RoleController } from "./role.controller";

@Module({
    providers: [RoleRepository, RoleService],
    controllers: [RoleController]
})
export class RoleModule {}