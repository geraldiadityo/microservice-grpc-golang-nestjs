import { Module } from '@nestjs/common';
import { CommonModule } from './common/common.module';
import { CategoryModule } from './category/category.module';
import { BarangModule } from './barang/barang.module';


@Module({
  imports: [
    CommonModule,
    CategoryModule,
    BarangModule
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
