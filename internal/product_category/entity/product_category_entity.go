package product_category

import (
	"database/sql"

	adminEntity "edtech.id/internal/admin/entity"
	"gorm.io/gorm"
)

type ProductCategory struct {
	ID          int64              `json:"id"`
	Name        string             `json:"name"`
	CreatedAt   sql.NullTime       `json:"created_at"`
	UpdatedAt   sql.NullTime       `json:"updated_at"`
	DeletedAt   gorm.DeletedAt     `json:"deleted_at"`
	CreatedBy   *adminEntity.Admin `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID int64              `json:"created_by" gorm:"column:created_by"`
	UpdatedBy   *adminEntity.Admin `gorm:"foreignKey:UpdatedByID;references:ID"`
	UpdatedByID *int64             `json:"updated_by" gorm:"column:updated_by"`
}
