package order_detail

import (
	"database/sql"

	orderEntity "edtech.id/internal/order/entity"
	productEntity "edtech.id/internal/product/entity"
	userEntity "edtech.id/internal/user/entity"

	"gorm.io/gorm"
)

type OrderDetail struct {
	ID          int64                  `json:"id"`
	Price       int64                  `json:"price"`
	ProductID   int64                  `json:"product_id"`
	Product     *productEntity.Product `gorm:"foreignKey:ProductID;references:ID"`
	OrderID     int64                  `json:"order_id"`
	Order       *orderEntity.Order     `gorm:"foreignKey:OrderID;references:ID"`
	CreatedAt   sql.NullTime           `json:"created_at"`
	UpdatedAt   sql.NullTime           `json:"updated_at"`
	DeletedAt   gorm.DeletedAt         `json:"deleted_at"`
	CreatedBy   *userEntity.User       `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID int64                  `json:"created_by" gorm:"column:created_by"`
	UpdatedBy   *userEntity.User       `gorm:"foreignKey:UpdatedByID;references:ID"`
	UpdatedByID *int64                 `json:"updated_by" gorm:"column:updated_by"`
}
