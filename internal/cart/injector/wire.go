//go:build wireinject
// +build wireinject

package cart

import (
	cartHandler "edtech.id/internal/cart/delivery/http"
	cartRepository "edtech.id/internal/cart/repository"
	cartUseCase "edtech.id/internal/cart/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *cartHandler.CartHandler {
	wire.Build(
		cartHandler.NewCartHandler,
		cartUseCase.NewCartUsecase,
		cartRepository.NewCartRepository,
	)

	return &cartHandler.CartHandler{}
}
