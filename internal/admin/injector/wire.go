//go:build wireinject
// +build wireinject

package admin

import (
	adminHandler "edtech.id/internal/admin/delivery/http"
	adminReposiotry "edtech.id/internal/admin/repository"
	adminUseCase "edtech.id/internal/admin/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeAdminHandler(db *gorm.DB) *adminHandler.AdminHandler {
	wire.Build(
		adminHandler.NewAdminHandler,
		adminUseCase.NewAdminUseCase,
		adminReposiotry.NewAdminRepository,
	)

	return &adminHandler.AdminHandler{}
}
