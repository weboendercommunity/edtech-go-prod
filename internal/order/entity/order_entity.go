package order

import (
	"database/sql"

	discountEntity "edtech.id/internal/discount/entity"
	userEntity "edtech.id/internal/user/entity"
	"gorm.io/gorm"
)

type Order struct {
	ID           int64                    `json:"id"`
	CheckoutLink string                   `json:"checkout_link"`
	Price        int64                    `json:"price"`
	TotalPrice   int64                    `json:"total_price"`
	ExternalID   string                   `json:"external_link"`
	Status       string                   `json:"status"`
	DiscountID   *int64                   `json:"discount_id"`
	Discount     *discountEntity.Discount `gorm:"foreignKey:DiscountID;references:ID"`
	UserID       int64                    `json:"user_id"`
	User         *userEntity.User         `gorm:"foreignKey:UserID;references:ID"`
	CreatedAt    sql.NullTime             `json:"created_at"`
	UpdatedAt    sql.NullTime             `json:"updated_at"`
	DeletedAt    gorm.DeletedAt           `json:"deleted_at"`
	CreatedBy    *userEntity.User         `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID  *int64                   `json:"created_by" gorm:"column:created_by"`
	UpdatedBy    *userEntity.User         `gorm:"foreignKey:UpdatedByID;references:ID"`
	UpdatedByID  *int64                   `json:"updated_by" gorm:"column:updated_by"`
}
