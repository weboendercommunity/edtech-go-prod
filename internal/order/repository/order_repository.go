package order

import (
	orderEntity "edtech.id/internal/order/entity"
	"edtech.id/pkg/utils"
	"gorm.io/gorm"
)

type OrderRepository interface {
	FindAll(offset int, limit int) []orderEntity.Order
	FindById(id int) (*orderEntity.Order, error)
	Create(orderEntity orderEntity.Order) (*orderEntity.Order, error)
	Update(orderEntity orderEntity.Order) (*orderEntity.Order, error)
}

type OrderRepositoryImpl struct {
	db *gorm.DB
}

// Update implements OrderRepository
func (or *OrderRepositoryImpl) Update(orderEntity orderEntity.Order) (*orderEntity.Order, error) {
	updatedOrder := or.db.Save(&orderEntity)

	if updatedOrder.Error != nil {
		return nil, updatedOrder.Error
	}

	return &orderEntity, nil
}

// Create implements OrderRepository
func (or *OrderRepositoryImpl) Create(orderEntity orderEntity.Order) (*orderEntity.Order, error) {
	createdOrder := or.db.Create(&orderEntity)

	if createdOrder.Error != nil {
		return nil, createdOrder.Error
	}

	return &orderEntity, nil
}

// FindAll implements OrderRepository
func (or *OrderRepositoryImpl) FindAll(offset int, limit int) []orderEntity.Order {
	var orders []orderEntity.Order

	or.db.Scopes(utils.Paginate(offset, limit)).Find(&orders)

	return orders
}

// FindById implements OrderRepository
func (or *OrderRepositoryImpl) FindById(id int) (*orderEntity.Order, error) {
	var order orderEntity.Order

	dataOrder := or.db.First(&order, id)

	if dataOrder.Error != nil {
		return nil, dataOrder.Error
	}

	return &order, nil
}

func NewOrderRepository(db *gorm.DB) OrderRepository {
	return &OrderRepositoryImpl{db}
}
