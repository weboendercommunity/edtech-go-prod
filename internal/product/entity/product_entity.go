package product

import (
	"database/sql"

	adminEntity "edtech.id/internal/admin/entity"
	productCategoryEntity "edtech.id/internal/product_category/entity"
	"gorm.io/gorm"
)

type Product struct {
	ID                int64                                  `json:"id"`
	ProductCategory   *productCategoryEntity.ProductCategory `gorm:"foreignKey:ProductCategoryID;references:ID"`
	ProductCategoryID int64                                  `json:"product_category_id" gorm:"column:product_category_id"`
	Title             string                                 `json:"title"`
	Image             *string                                `json:"image"`
	Video             *string                                `json:"video"`
	Description       string                                 `json:"description"`
	Price             int64                                  `json:"price"`
	CreatedAt         sql.NullTime                           `json:"created_at"`
	UpdatedAt         sql.NullTime                           `json:"updated_at"`
	DeletedAt         gorm.DeletedAt                         `json:"deleted_at"`
	CreatedBy         *adminEntity.Admin                     `gorm:"foreignKey:CreatedByID;references:ID"`
	CreatedByID       int64                                  `json:"created_by" gorm:"column:created_by"`
	UpdatedBy         *adminEntity.Admin                     `gorm:"foreignKey:UpdatedByID;references:ID"`
	UpdatedByID       *int64                                 `json:"updated_by" gorm:"column:updated_by"`
}
