//go:build wireinject
// +build wireinject

package product

import (
	productHandler "edtech.id/internal/product/delivery/http"
	productRepository "edtech.id/internal/product/repository"
	productUsecase "edtech.id/internal/product/usecase"
	fileUpload "edtech.id/pkg/fileupload/cloudinary"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializedService(db *gorm.DB) *productHandler.ProductHandler {
	wire.Build(
		productHandler.NewProductHandler,
		productUsecase.NewProductUsecase,
		productRepository.NewProductRepository,
		fileUpload.NewFileUpload,
	)

	return &productHandler.ProductHandler{}

}
