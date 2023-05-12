package discount

import (
	"database/sql"

	adminEntity "edtech.id/internal/admin/entity"
	"gorm.io/gorm"
)

type Discount struct {
	ID                int64              `json:"id"`
	Name              string             `json:"name"`
	Code              string             `json:"code"`
	Quantity          int64              `json:"quantity"`
	Value             int64              `json:"value"`
	RemainingQuantity int64              `json:"remaining_quantity"`
	Type              DiscountType       `json:"type"`
	StartDate         sql.NullTime       `json:"start_date"`
	EndDate           sql.NullTime       `json:"end_date"`
	CreatedAt         sql.NullTime       `json:"created_at"`
	UpdatedAt         sql.NullTime       `json:"updated_at"`
	DeletedAt         gorm.DeletedAt     `json:"deleted_at"`
	CreatedBy         *adminEntity.Admin `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID       *int64             `json:"created_by" gorm:"column:created_by"`
	UpdatedBy         *adminEntity.Admin `gorm:"foreignKey:UpdatedByID;references:ID"`
	UpdatedByID       *int64             `json:"updated_by" gorm:"column:updated_by"`
}

type DiscountType string

const (
	Precentage DiscountType = "percentage"
	Fixed      DiscountType = "fixed"
)
