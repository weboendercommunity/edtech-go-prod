package class_room

import (
	"database/sql"

	classRoomEntity "edtech.id/internal/class_room/entity"
	productEntity "edtech.id/internal/product/entity"
	userEntity "edtech.id/internal/user/entity"
	"gorm.io/gorm"
)

type ClassRoomResponseBody struct {
	ID        int64                 `json:"id"`
	User      userEntity.User       `json:"user"`
	Product   productEntity.Product `json:"product"`
	CreatedAt sql.NullTime          `json:"created_at"`
	UpdatedAt sql.NullTime          `json:"updated_at"`
	DeletedAt gorm.DeletedAt        `json:"deleted_at"`
	CreatedBy userEntity.User       `json:"created_by"`
	UpdatedBy userEntity.User       `json:"updated_by"`
}

func CreateClassRoomResponse(classRoomEntity *classRoomEntity.ClassRoom) ClassRoomResponseBody {

	return ClassRoomResponseBody{
		ID:        classRoomEntity.ID,
		User:      classRoomEntity.User,
		Product:   classRoomEntity.Product,
		CreatedAt: classRoomEntity.CreatedAt,
		UpdatedAt: classRoomEntity.UpdatedAt,
		DeletedAt: classRoomEntity.DeletedAt,
		CreatedBy: classRoomEntity.CreatedBy,
		UpdatedBy: classRoomEntity.UpdatedBy,
	}
}

type ClassRoomListResponse []ClassRoomResponseBody

func CreateClassRoomListResponse(classRoomEntities []classRoomEntity.ClassRoom) ClassRoomListResponse {

	var classRoomListResponse ClassRoomListResponse

	for _, c := range classRoomEntities {

		c.Product.VideoUrl = c.Product.Video

		classRoomListResponse = append(classRoomListResponse, CreateClassRoomResponse(&c))
	}

	return classRoomListResponse
}
