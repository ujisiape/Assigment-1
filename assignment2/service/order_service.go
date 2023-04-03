package service

import (
	"assignment-2/dto"
	"assignment-2/entity"
	"assignment-2/pkg/errs"
	"assignment-2/repository/order_repository"
	"fmt"
	"net/http"
)

type OrderService interface {
	CreateOrder(payload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr)
	GetAllOrders() (*dto.GetAllOrdersResponse, errs.MessageErr)
	GetOrderByID(orderID uint) (*dto.GetOrderByIDResponse, errs.MessageErr)
	UpdateOrderByID(orderID uint, payload dto.NewOrderRequest) (*dto.GetOrderByIDResponse, errs.MessageErr)
	DeleteOrderByID(orderID uint) (*dto.DeleteOrderByIDResponse, errs.MessageErr)
}

type orderService struct {
	orderRepo order_repository.OrderRepository
}

func NewOrderService(orderRepo order_repository.OrderRepository) OrderService {
	return &orderService{
		orderRepo: orderRepo,
	}
}

func (o *orderService) CreateOrder(payload dto.NewOrderRequest) (*dto.NewOrderResponse, errs.MessageErr) {
	orderPayload := payload.OrderRequestToEntity()

	itemsPayload := []entity.Item{}
	for _, eachItem := range payload.Items {
		item := eachItem.ItemRequestToEntity()

		itemsPayload = append(itemsPayload, *item)
	}

	newOrder, err := o.orderRepo.CreateOrder(orderPayload, itemsPayload)
	if err != nil {
		return nil, err
	}

	items := []dto.NewItemRequest{}
	for _, eachItem := range newOrder.Items {
		item := dto.NewItemRequest{
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
		}

		items = append(items, item)
	}

	response := &dto.NewOrderResponse{
		Message:    fmt.Sprintf("Order with id %d has been created", newOrder.ID),
		StatusCode: http.StatusCreated,
		Data: dto.NewOrderRequest{
			OrderedAt:    newOrder.OrderedAt,
			CustomerName: newOrder.CustomerName,
			Items:        items,
		},
	}

	return response, nil
}

func (o *orderService) GetAllOrders() (*dto.GetAllOrdersResponse, errs.MessageErr) {
	orders, err := o.orderRepo.GetAllOrders()
	if err != nil {
		return nil, err
	}

	data := []dto.OrderData{}

	for _, eachOrder := range orders {
		items := []dto.ItemData{}

		for _, eachItem := range eachOrder.Items {
			item := dto.ItemData{
				ID:          eachItem.ID,
				CreatedAt:   eachItem.CreatedAt,
				UpdatedAt:   eachItem.UpdatedAt,
				ItemCode:    eachItem.ItemCode,
				Description: eachItem.Description,
				Quantity:    eachItem.Quantity,
				OrderID:     eachItem.OrderID,
			}

			items = append(items, item)
		}

		order := dto.OrderData{
			ID:           eachOrder.ID,
			CreatedAt:    eachOrder.CreatedAt,
			UpdatedAt:    eachOrder.UpdatedAt,
			CustomerName: eachOrder.CustomerName,
			Items:        items,
		}

		data = append(data, order)
	}

	response := &dto.GetAllOrdersResponse{
		Message:    "success",
		StatusCode: http.StatusOK,
		Data:       data,
	}

	return response, nil
}

func (o *orderService) GetOrderByID(orderID uint) (*dto.GetOrderByIDResponse, errs.MessageErr) {
	order, err := o.orderRepo.GetOrderByID(orderID)
	if err != nil {
		return nil, err
	}

	items := []dto.ItemData{}
	for _, eachItem := range order.Items {
		item := dto.ItemData{
			ID:          eachItem.ID,
			CreatedAt:   eachItem.CreatedAt,
			UpdatedAt:   eachItem.UpdatedAt,
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
			OrderID:     eachItem.OrderID,
		}

		items = append(items, item)
	}

	response := &dto.GetOrderByIDResponse{
		StatusCode: http.StatusOK,
		Message:    "success",
		Data: dto.OrderData{
			ID:           order.ID,
			CreatedAt:    order.CreatedAt,
			UpdatedAt:    order.UpdatedAt,
			CustomerName: order.CustomerName,
			Items:        items,
		},
	}

	return response, nil
}

func (o *orderService) UpdateOrderByID(orderID uint, payload dto.NewOrderRequest) (*dto.GetOrderByIDResponse, errs.MessageErr) {
	orderPayload := payload.OrderRequestToEntity()

	itemsPayload := []entity.Item{}
	for _, eachItem := range payload.Items {
		item := eachItem.ItemRequestToEntity()

		itemsPayload = append(itemsPayload, *item)
	}

	updatedOrder, err := o.orderRepo.UpdateOrderByID(orderID, orderPayload, itemsPayload)
	if err != nil {
		return nil, err
	}

	items := []dto.ItemData{}
	for _, eachItem := range updatedOrder.Items {
		item := dto.ItemData{
			ID:          eachItem.ID,
			CreatedAt:   eachItem.CreatedAt,
			UpdatedAt:   eachItem.UpdatedAt,
			ItemCode:    eachItem.ItemCode,
			Description: eachItem.Description,
			Quantity:    eachItem.Quantity,
			OrderID:     eachItem.OrderID,
		}

		items = append(items, item)
	}

	response := &dto.GetOrderByIDResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Order with id %d has been updated", orderID),
		Data: dto.OrderData{
			ID:           updatedOrder.ID,
			CreatedAt:    updatedOrder.CreatedAt,
			UpdatedAt:    updatedOrder.UpdatedAt,
			CustomerName: updatedOrder.CustomerName,
			Items:        items,
		},
	}

	return response, nil
}

func (o *orderService) DeleteOrderByID(orderID uint) (*dto.DeleteOrderByIDResponse, errs.MessageErr) {
	if err := o.orderRepo.DeleteOrderByID(orderID); err != nil {
		return nil, err
	}

	response := &dto.DeleteOrderByIDResponse{
		StatusCode: http.StatusOK,
		Message:    fmt.Sprintf("Order with id %d has been deleted", orderID),
	}

	return response, nil
}
