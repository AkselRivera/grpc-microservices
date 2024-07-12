package service

import (
	"context"
	"log"
	"strings"

	pb "github.com/akselrivera/grpc-order-microservice/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"

	"github.com/google/uuid"
)

type OrderService struct {
	pb.UnimplementedOrderServiceServer
}

var productClient pb.ProductServiceClient

var InternalError error = status.Errorf(codes.Internal, "product microservice unavailable")

func init() {
	connection, err := grpc.NewClient("product-service:50052", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Println("Error creating connection: ", err)
	}

	productClient = pb.NewProductServiceClient(connection)
}

var orders []*pb.Order = []*pb.Order{}

// GetOrder implements proto.OrderServiceServer.
func (o OrderService) GetOrder(c context.Context, getOrderRequest *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	orderId := getOrderRequest.Id

	products := []*pb.SelectedProduct{}
	completeOrder := pb.CompleteOrder{}

	for _, order := range orders {
		if order.Id == orderId {
			products = order.Products
			break
		}
	}

	if len(products) == 0 {
		return &pb.GetOrderResponse{}, status.Errorf(codes.NotFound, "order not found")
	}

	for _, product := range products {
		getProductResponse, err := productClient.GetProduct(c, &pb.GetProductRequest{Id: product.Id})

		if err != nil {
			log.Println("Error getting product: ", err)
			if strings.Contains(err.Error(), "Unavailable") {
				return &pb.GetOrderResponse{}, InternalError
			}
			return &pb.GetOrderResponse{}, status.Errorf(codes.NotFound, "product not found")
		}

		getProductResponse.Product.Quantity = product.Quantity
		completeOrder.Products = append(completeOrder.Products, getProductResponse.Product)
		completeOrder.Total += float64(getProductResponse.Product.Price) * float64(product.Quantity)
	}

	order := pb.GetOrderResponse{
		Order: &pb.CompleteOrder{
			Id:       orderId,
			Products: completeOrder.Products,
			Total:    completeOrder.Total,
			Status:   completeOrder.Status,
		},
	}
	return &order, nil
}

// ListOrders implements proto.OrderServiceServer.
// Subtle: this method shadows the method (UnimplementedOrderServiceServer).ListOrders of OrderService.UnimplementedOrderServiceServer.
func (o OrderService) ListOrders(c context.Context, listOrdersRequest *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	completeOrders := pb.ListOrdersResponse{}

	for _, order := range orders {
		var products []*pb.Product
		var total float64

		for _, product := range order.Products {
			getProductResponse, err := productClient.GetProduct(c, &pb.GetProductRequest{Id: product.Id})
			if err != nil {
				if strings.Contains(err.Error(), "Unavailable") {
					return &pb.ListOrdersResponse{}, InternalError
				}
				log.Println("Error getting product: ", err)
				return &pb.ListOrdersResponse{}, status.Errorf(codes.NotFound, "product not found")
			}

			getProductResponse.Product.Quantity = product.Quantity
			products = append(products, getProductResponse.Product)
			total += float64(getProductResponse.Product.Price) * float64(product.Quantity)
		}

		completeOrders.Orders = append(completeOrders.Orders, &pb.CompleteOrder{
			Id:       order.Id,
			Products: products,
			Total:    total,
			Status:   order.Status,
		})

	}
	return &completeOrders, nil
}

// CreateOrder implements proto.OrderServiceServer.
func (o OrderService) CreateOrder(c context.Context, createOrderRequest *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {

	if len(createOrderRequest.Products) == 0 {
		return &pb.CreateOrderResponse{}, status.Errorf(codes.InvalidArgument, "order must have at least one product")
	}

	completeOrder := pb.CompleteOrder{}
	completeOrder.Total = 0
	completeOrder.Id = uuid.NewString()
	completeOrder.Status = pb.EnumStatus_PENDING

	for _, selectedProduct := range createOrderRequest.Products {

		getProductResponse, err := productClient.GetProduct(c, &pb.GetProductRequest{Id: selectedProduct.Id})

		if err != nil {
			if strings.Contains(err.Error(), "Unavailable") {
				return &pb.CreateOrderResponse{}, InternalError
			}

			return &pb.CreateOrderResponse{}, status.Errorf(codes.InvalidArgument, "product not found")
		}

		checkStockResponse, err := productClient.HandleProductStock(c, &pb.HandleProductStockRequest{Id: selectedProduct.Id, Quantity: selectedProduct.Quantity, Action: pb.EnumAction_SUBTRACT})

		if err != nil {
			if strings.Contains(err.Error(), "Unavailable") {
				return &pb.CreateOrderResponse{}, InternalError
			}

			log.Println("Error checking stock: ", err)
			if strings.Contains(err.Error(), "Product is inactive") {
				return &pb.CreateOrderResponse{}, status.Errorf(codes.InvalidArgument, "product is inactive")
			}

			return &pb.CreateOrderResponse{}, status.Errorf(codes.InvalidArgument, "product out of stock")
		}

		if checkStockResponse.Success {
			getProductResponse.Product.Quantity = selectedProduct.Quantity
			completeOrder.Products = append(completeOrder.Products, getProductResponse.Product)
			completeOrder.Total += float64(getProductResponse.Product.Price) * float64(selectedProduct.Quantity)
		}

		orders = append(orders, &pb.Order{
			Id:       completeOrder.Id,
			Products: []*pb.SelectedProduct{selectedProduct},
			Status:   completeOrder.Status,
		})

	}

	return &pb.CreateOrderResponse{Order: &completeOrder}, nil

}

// ChangeOrderStatus implements proto.OrderServiceServer.
func (o OrderService) ChangeOrderStatus(c context.Context, changeOrderStatusRequest *pb.ChangeOrderStatusRequest) (*pb.ChangeOrderStatusResponse, error) {

	if changeOrderStatusRequest.Status == pb.EnumStatus_CANCELLED {
		return &pb.ChangeOrderStatusResponse{}, status.Errorf(codes.InvalidArgument, "order cannot be cancelled, please use cancel order instead")
	}

	orderId := changeOrderStatusRequest.Id
	newOrder := pb.ChangeOrderStatusResponse{
		Order: &pb.CompleteOrder{},
	}
	var products []*pb.SelectedProduct
	for _, order := range orders {
		if order.Id == orderId {
			if order.Status == pb.EnumStatus_CANCELLED {
				return &pb.ChangeOrderStatusResponse{}, status.Errorf(codes.InvalidArgument, "order already cancelled")
			}
			newOrder.Order.Id = order.Id
			products = order.Products
			order.Status = changeOrderStatusRequest.Status
			newOrder.Order.Status = changeOrderStatusRequest.Status
			break
		}
	}

	if newOrder.Order.Id == "" {
		return &pb.ChangeOrderStatusResponse{}, status.Errorf(codes.NotFound, "order not found")
	}

	for _, product := range products {
		getProductResponse, err := productClient.GetProduct(c, &pb.GetProductRequest{Id: product.Id})

		if err != nil {
			if strings.Contains(err.Error(), "Unavailable") {
				return &pb.ChangeOrderStatusResponse{}, InternalError
			}
			return &pb.ChangeOrderStatusResponse{}, status.Errorf(codes.NotFound, "product not found")
		}

		getProductResponse.Product.Quantity = product.Quantity
		newOrder.Order.Products = append(newOrder.Order.Products, getProductResponse.Product)
		newOrder.Order.Total += float64(getProductResponse.Product.Price) * float64(product.Quantity)

	}

	return &newOrder, nil
}

// CancelOrder implements proto.OrderServiceServer.
func (o OrderService) CancelOrder(c context.Context, cancelOrderRequest *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	var foundOrder bool
	for _, order := range orders {
		if order.Id == cancelOrderRequest.Id {
			foundOrder = true
			if order.Status == pb.EnumStatus_CANCELLED {
				return &pb.CancelOrderResponse{}, status.Errorf(codes.InvalidArgument, "order already cancelled")
			}
			order.Status = pb.EnumStatus_CANCELLED
			break
		}
	}

	if !foundOrder {
		return &pb.CancelOrderResponse{}, status.Errorf(codes.NotFound, "order not found")
	}

	getOrderResponse, err := o.GetOrder(c, &pb.GetOrderRequest{Id: cancelOrderRequest.Id})

	if err != nil {
		if strings.Contains(err.Error(), "Unavailable") {
			return &pb.CancelOrderResponse{}, InternalError
		}

		log.Println("Error getting order to cancel: ", err)
		return &pb.CancelOrderResponse{}, status.Errorf(codes.NotFound, "order not found")
	}

	for _, product := range getOrderResponse.Order.Products {

		handleProductStock, err := productClient.HandleProductStock(c, &pb.HandleProductStockRequest{Id: product.Id, Quantity: product.Quantity, Action: pb.EnumAction_ADD})

		if err != nil || !handleProductStock.Success {
			if strings.Contains(err.Error(), "Unavailable") {
				return &pb.CancelOrderResponse{}, InternalError
			}

			log.Println("Error checking stock to cancel order: ", err)
			return &pb.CancelOrderResponse{}, status.Errorf(codes.Internal, "product not found")
		}

	}

	return &pb.CancelOrderResponse{Success: true}, nil
}
