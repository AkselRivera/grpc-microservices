import { Controller } from "@nestjs/common";
import { GrpcMethod } from "@nestjs/microservices";
import { ProductService } from "./product.service";
import {
  CreateProductRequest,
  GetProductRequest,
  UpdateProductRequest,
  DeleteProductRequest,
  GetProductsResponse,
  GetProductResponse,
  CreateProductResponse,
  UpdateProductResponse,
  DeleteProductResponse,
  HandleProductStockRequest,
  HandleProductStockResponse,
} from "src/proto/product";

@Controller("product")
export class ProductController {
  constructor(private readonly productService: ProductService) {}

  @GrpcMethod("ProductService", "CreateProduct")
  createProduct(request: CreateProductRequest): CreateProductResponse {
    return { product: this.productService.createProduct(request) };
  }

  @GrpcMethod("ProductService", "GetProduct")
  getProduct(request: GetProductRequest): GetProductResponse {
    return { product: this.productService.getProduct(request.id) };
  }

  @GrpcMethod("ProductService", "UpdateProduct")
  updateProduct(request: UpdateProductRequest): UpdateProductResponse {
    return { product: this.productService.updateProduct(request) };
  }

  @GrpcMethod("ProductService", "DeleteProduct")
  deleteProduct(request: DeleteProductRequest): DeleteProductResponse {
    return { success: this.productService.deleteProduct(request.id) };
  }

  @GrpcMethod("ProductService", "ListProducts")
  listProducts(): GetProductsResponse {
    return this.productService.listProducts();
  }

  @GrpcMethod("ProductService", "HandleProductStock")
  handleProductStock(
    request: HandleProductStockRequest,
  ): HandleProductStockResponse {
    return this.productService.handleProductStock(request);
  }
}
