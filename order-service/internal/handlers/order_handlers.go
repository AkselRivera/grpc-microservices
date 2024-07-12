package handlers

import (
	"context"

	"github.com/akselrivera/grpc-order-microservice/internal/service"
	pb "github.com/akselrivera/grpc-order-microservice/proto"
)

type OrderController struct {
	pb.UnimplementedOrderServiceServer
}

var orderService service.OrderService

// GetOrder implements proto.OrderControllerServer.
func (o OrderController) GetOrder(c context.Context, getOrderRequest *pb.GetOrderRequest) (*pb.GetOrderResponse, error) {
	getOrderResponse, err := orderService.GetOrder(c, getOrderRequest)

	if err != nil || getOrderResponse.Order.Id == "" {
		return &pb.GetOrderResponse{}, err
	}

	return getOrderResponse, nil
}

// ListOrders implements proto.OrderControllerServer.
// Subtle: this method shadows the method (UnimplementedOrderControllerServer).ListOrders of OrderController.UnimplementedOrderControllerServer.
func (o OrderController) ListOrders(c context.Context, listOrdersRequest *pb.ListOrdersRequest) (*pb.ListOrdersResponse, error) {
	listOrdersResponse, err := orderService.ListOrders(c, listOrdersRequest)

	if err != nil {
		return &pb.ListOrdersResponse{}, err
	}

	return listOrdersResponse, nil
}

// CreateOrder implements proto.OrderControllerServer.
func (o OrderController) CreateOrder(c context.Context, createOrderRequest *pb.CreateOrderRequest) (*pb.CreateOrderResponse, error) {
	createOrderRequestResponse, err := orderService.CreateOrder(c, createOrderRequest)

	if err != nil {
		return &pb.CreateOrderResponse{}, err
	}

	return createOrderRequestResponse, nil
}

// ChangeOrderStatus implements proto.OrderControllerServer.
func (o OrderController) ChangeOrderStatus(c context.Context, changeOrderStatusRequest *pb.ChangeOrderStatusRequest) (*pb.ChangeOrderStatusResponse, error) {
	changeOrderStatusResponse, err := orderService.ChangeOrderStatus(c, changeOrderStatusRequest)

	if err != nil {
		return &pb.ChangeOrderStatusResponse{}, err
	}

	return changeOrderStatusResponse, nil
}

// CancelOrder implements proto.OrderControllerServer.
func (o OrderController) CancelOrder(c context.Context, cancelOrderRequest *pb.CancelOrderRequest) (*pb.CancelOrderResponse, error) {
	cancelOrderResponse, err := orderService.CancelOrder(c, cancelOrderRequest)

	if err != nil {
		return &pb.CancelOrderResponse{}, err
	}

	return cancelOrderResponse, nil
}
