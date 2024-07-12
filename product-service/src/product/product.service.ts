import { status } from "@grpc/grpc-js";
import { Injectable } from "@nestjs/common";
import { RpcException } from "@nestjs/microservices";
import { randomUUID } from "crypto";
import {
  CreateProductRequest,
  EnumAction,
  EnumProductIsActive,
  GetProductsResponse,
  HandleProductStockRequest,
  HandleProductStockResponse,
  Product,
  UpdateProductRequest,
} from "src/proto/product";

@Injectable()
export class ProductService {
  private products: Product[] = [];

  createProduct(request: CreateProductRequest): Product {
    const product: Product = {
      id: randomUUID(),
      ...request,
      isActive: EnumProductIsActive.ACTIVE,
    };
    this.products.push(product);
    return product;
  }

  getProduct(id: string): Product {
    const product = this.products.find((product) => product.id === id);

    if (!product) {
      throw new RpcException({
        code: status.NOT_FOUND,
        message: "Product not found",
      });
    }
    return this.products.find((product) => product.id === id);
  }

  updateProduct(request: UpdateProductRequest): Product {
    const index = this.products.findIndex(
      (product) => product.id === request.product.id,
    );
    if (index !== -1) {
      this.products[index] = request.product;
      return this.products[index];
    }

    throw new RpcException({
      code: status.NOT_FOUND,
      message: "Product not found",
    });
  }

  deleteProduct(id: string): boolean {
    const index = this.products.findIndex((product) => product.id === id);
    if (index !== -1) {
      this.products[index].isActive = EnumProductIsActive.INACTIVE;
      return true;
    }
    throw new RpcException({
      code: status.NOT_FOUND,
      message: "Product not found",
    });
  }

  listProducts(): GetProductsResponse {
    return {
      products: this.products.filter(
        (product) => product.isActive === EnumProductIsActive.ACTIVE,
      ),
    };
  }
  handleProductStock(
    request: HandleProductStockRequest,
  ): HandleProductStockResponse {
    const index = this.products.findIndex(
      (product) => product.id === request.id,
    );
    if (index === -1) {
      throw new RpcException({
        code: status.NOT_FOUND,
        message: "Product not found",
      });
    }

    if (this.products[index].isActive !== EnumProductIsActive.ACTIVE) {
      throw new RpcException({
        code: status.INVALID_ARGUMENT,
        message: "Product is inactive",
      });
    }

    switch (request.action) {
      case EnumAction.ADD:
        this.products[index].quantity += request.quantity;
        break;
      case EnumAction.SUBTRACT:
        if (this.products[index].quantity < request.quantity) {
          throw new RpcException({
            code: status.INVALID_ARGUMENT,
            message: "Not enough stock",
          });
        }

        this.products[index].quantity -= request.quantity;
        break;
      case EnumAction.CLEAR:
        this.products[index].quantity = 0;
        break;
      default:
        throw new RpcException({
          code: status.INVALID_ARGUMENT,
          message: "Invalid action",
        });
    }
    return { success: true };
  }
}
