import { Inject, Injectable, OnModuleDestroy, OnModuleInit } from "@nestjs/common";
import { PrismaClient } from "@prisma/client";
import { WINSTON_MODULE_PROVIDER } from "nest-winston";
import { Logger } from "winston";

@Injectable()
export class PrismaService extends PrismaClient implements OnModuleInit, OnModuleDestroy {
    constructor(
        @Inject(WINSTON_MODULE_PROVIDER) private readonly logger: Logger
    ) {
        super();
        this.$use(async (params, next) => {
            const start = Date.now();
            const result = await next(params);
            const duration = Date.now() - start;
            this.logger.info({
                message: `Query ${params.model}.${params.action} took ${duration}ms`,
                meta: { params, duration }
            });

            return result;
        })
    }

    async onModuleInit() {
        await this.$connect();
        this.logger.info('Connected to ORM');
    }

    async onModuleDestroy() {
        this.logger.info('Disconnect to ORM');
        await this.$disconnect();
    }
}