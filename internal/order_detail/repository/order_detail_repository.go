package order_detail

import (
	orderDetailEntity "edtech.id/internal/order_detail/entity"
	"gorm.io/gorm"
)

type OrderDetailRepository interface {
	Create(orderDetailEntity orderDetailEntity.OrderDetail) (*orderDetailEntity.OrderDetail, error)
}

type OrderDetailRepositoryImpl struct {
	db *gorm.DB
}

// Create implements OrderDetail
func (or *OrderDetailRepositoryImpl) Create(orderDetailEntity orderDetailEntity.OrderDetail) (*orderDetailEntity.OrderDetail, error) {
	createdOrderDetail := or.db.Create(&orderDetailEntity)

	if createdOrderDetail.Error != nil {
		return nil, createdOrderDetail.Error
	}

	return &orderDetailEntity, nil
}

func NewOrderDetailRepository(db *gorm.DB) OrderDetailRepository {
	return &OrderDetailRepositoryImpl{db}
}
