import { Module } from "@nestjs/common";
import { BarangRepository } from "./barang.repository";
import { BarangService } from "./barang.service";
import { BarangController } from "./barang.controller";

@Module({
    providers: [BarangRepository, BarangService],
    controllers: [BarangController]
})
export class BarangModule {}