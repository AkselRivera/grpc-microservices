import { NestFactory } from "@nestjs/core";
import { Transport, MicroserviceOptions } from "@nestjs/microservices";
import { AppModule } from "./app.module";

async function bootstrap() {
  const app = await NestFactory.createMicroservice<MicroserviceOptions>(
    AppModule,
    {
      transport: Transport.GRPC,
      options: {
        // url: 'http://localhost:6379',
        url: "0.0.0.0:50052",
        protoPath: "./proto/product.proto",
        package: "product",
      },
    },
  );
  await app.listen();
}
bootstrap();
