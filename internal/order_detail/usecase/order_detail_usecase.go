package order_detail

import (
	orderDetailDto "edtech.id/internal/order_detail/dto"
	orderDetailEntity "edtech.id/internal/order_detail/entity"
	orderDetailRepository "edtech.id/internal/order_detail/repository"
)

type OrderDetailUsecase interface {
	Create(orderDetailDto orderDetailDto.OrderDetailRequestBody) (*orderDetailEntity.OrderDetail, error)
}

type OrderDetailUsecaseImpl struct {
	orderDetailRepository orderDetailRepository.OrderDetailRepository
}

// Create implements OrderDetailUsecase
func (odu *OrderDetailUsecaseImpl) Create(orderDetailDto orderDetailDto.OrderDetailRequestBody) (*orderDetailEntity.OrderDetail, error) {
	orderDetailEntity := orderDetailEntity.OrderDetail{
		Price:       orderDetailDto.Price,
		ProductID:   orderDetailDto.ProductID,
		OrderID:     orderDetailDto.OrderID,
		CreatedByID: orderDetailDto.CreatedBy,
	}

	createdOrderDetail, err := odu.orderDetailRepository.Create(orderDetailEntity)

	if err != nil {
		return nil, err
	}

	return createdOrderDetail, nil
}

func NewOrderDetailUsecase(orderDetailRepository orderDetailRepository.OrderDetailRepository) OrderDetailUsecase {
	return &OrderDetailUsecaseImpl{orderDetailRepository}
}
