//go: build wireinject
//go:build wireinject
// +build wireinject

package class_room

import (
	classRoomHandler "edtech.id/internal/class_room/delivery/http"
	classRoomRepository "edtech.id/internal/class_room/repository"
	classRoomUsecase "edtech.id/internal/class_room/usecase"
	"github.com/google/wire"
	"gorm.io/gorm"
)

func InitializeService(db *gorm.DB) *classRoomHandler.ClassRoomHandler {
	wire.Build(
		classRoomHandler.NewClassRoomHandler,
		classRoomUsecase.NewClassRoomUsecase,
		classRoomRepository.NewClassRoomRepository,
	)

	return &classRoomHandler.ClassRoomHandler{}
}
