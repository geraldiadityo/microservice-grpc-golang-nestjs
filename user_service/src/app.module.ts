import { Module } from '@nestjs/common';
import { CommonModule } from './common/common.module';
import { RoleModule } from './role/role.module';
import { PenggunaModule } from './pengguna/pengguna.module';


@Module({
  imports: [
    CommonModule,
    RoleModule,
    PenggunaModule
  ],
  controllers: [],
  providers: [],
})
export class AppModule {}
