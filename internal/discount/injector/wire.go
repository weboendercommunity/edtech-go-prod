// go:build wireinject
//go:build wireinject
// +build wireinject

package discount

import (
	discountHandler "edtech.id/internal/discount/delivery/http"
	discountRepository "edtech.id/internal/discount/repository"
	discountUsecase "edtech.id/internal/discount/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *discountHandler.DiscountHandler {
	wire.Build(
		discountHandler.NewDiscountHandler,
		discountUsecase.NewDiscountUsecase,
		discountRepository.NewDiscountRepository,
	)

	return &discountHandler.DiscountHandler{}
}
