import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { MicroserviceOptions, Transport } from '@nestjs/microservices';
import { join } from 'path';
import { WINSTON_MODULE_NEST_PROVIDER } from 'nest-winston';

async function bootstrap() {
  const port = process.env.PORT;
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(AppModule, {
    transport: Transport.GRPC,
    options: {
      url: `localhost:${Number(port)}`,
      package: ['role', 'pengguna'],
      protoPath: [
        join(__dirname, '..', 'src/proto', 'role.proto'),
        join(__dirname, '..', 'src/proto', 'pengguna.proto')
      ]
    }
  });
  const logger = app.get(WINSTON_MODULE_NEST_PROVIDER);
  app.useLogger(logger);
  app.enableShutdownHooks();
  await app.listen();
  console.log(`User Service running on port: ${port}`)
}
bootstrap();
