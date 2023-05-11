package product

import (
	"database/sql"

	adminEntity "edtech.id/internal/admin/entity"
	productEntity "edtech.id/internal/product/entity"
	productCategoryEntity "edtech.id/internal/product_category/entity"
	"gorm.io/gorm"
)

type ProductResponseBody struct {
	ID              int64                                  `json:"id"`
	ProductCategory *productCategoryEntity.ProductCategory `json:"product_category"`
	Title           string                                 `json:"title"`
	Image           string                                 `json:"image"`
	Video           string                                 `json:"video"`
	Description     string                                 `json:"description"`
	Price           int64                                  `json:"price"`
	CreatedBy       *adminEntity.Admin                     `json:"created_by"`
	UpdatedBy       *adminEntity.Admin                     `json:"updated_by"`
	CreatedAt       sql.NullTime                           `json:"created_at"`
	UpdatedAt       sql.NullTime                           `json:"updated_at"`
	DeletedAt       gorm.DeletedAt                         `json:"deleted_at"`
}

func CreateProductResponse(productEntity productEntity.Product) ProductResponseBody {
	return ProductResponseBody{
		ProductCategory: productEntity.ProductCategory,
		Title:           productEntity.Title,
		Image:           *productEntity.Image,
		Video:           *productEntity.Video,
		Description:     productEntity.Description,
		Price:           productEntity.Price,
		CreatedBy:       productEntity.CreatedBy,
		UpdatedBy:       productEntity.UpdatedBy,
		DeletedAt:       gorm.DeletedAt{},
	}
}

type ProductListResponse []ProductResponseBody

func CreateProductListResponse(products []productEntity.Product) []ProductResponseBody {
	productResponse := ProductListResponse{}

	for _, p := range products {
		product := CreateProductResponse(p)
		productResponse = append(productResponse, product)
	}

	return productResponse
}
