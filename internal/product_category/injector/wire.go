//go:build wireinject
// +build wireinject

package product_category

import (
	productCategoryHandler "edtech.id/internal/product_category/delivery/http"
	productCategoryRepository "edtech.id/internal/product_category/repository"
	productCategoryUseCase "edtech.id/internal/product_category/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *productCategoryHandler.ProductCategoryHandler {
	wire.Build(
		productCategoryHandler.NewProductCategoryHandler,
		productCategoryUseCase.NewProductCategoryUseCase,
		productCategoryRepository.NewProductCategoryRepository,
	)

	return &productCategoryHandler.ProductCategoryHandler{}
}
