import { Module } from "@nestjs/common";
import { PenggunaRepository } from "./pengguna.repository";
import { PenggunaService } from "./pengguna.service";
import { PenggunaController } from "./pengguna.controller";

@Module({
    providers: [PenggunaRepository, PenggunaService],
    controllers: [PenggunaController]
})
export class PenggunaModule {}