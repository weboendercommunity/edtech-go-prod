package class_room

import (
	classRoomEntity "edtech.id/internal/class_room/entity"
	"edtech.id/pkg/utils"
	"gorm.io/gorm"
)

type ClassRoomRepository interface {
	FindAllByUserID(offset int, limit int, userId int) ([]classRoomEntity.ClassRoom, error)
	FindOneByUserIDAndProductID(userId int, productId int) (*classRoomEntity.ClassRoom, error)
	Create(classRoomEntity classRoomEntity.ClassRoom) (*classRoomEntity.ClassRoom, error)
}

type ClassRoomRepositoryImpl struct {
	db *gorm.DB
}

// Create implements ClassRoomRepository
func (crr *ClassRoomRepositoryImpl) Create(classRoomEntity classRoomEntity.ClassRoom) (*classRoomEntity.ClassRoom, error) {
	createdClassRoom := crr.db.Create(&classRoomEntity)

	if createdClassRoom.Error != nil {
		return nil, createdClassRoom.Error
	}

	return &classRoomEntity, nil
}

// FindAllByUserID implements ClassRoomRepository
func (crr *ClassRoomRepositoryImpl) FindAllByUserID(offset int, limit int, userId int) ([]classRoomEntity.ClassRoom, error) {
	var classRooms []classRoomEntity.ClassRoom

	dataClassRooms := crr.db.
		Scopes(utils.Paginate(offset, limit)).
		Preload("Product").
		Preload("User").
		Where("user_id = ?", userId).Find(&classRooms)

	if dataClassRooms.Error != nil {
		return nil, dataClassRooms.Error
	}

	return classRooms, nil
}

// FindOneByUserIDAndProductID implements ClassRoomRepository
func (crr *ClassRoomRepositoryImpl) FindOneByUserIDAndProductID(userId int, productId int) (*classRoomEntity.ClassRoom, error) {
	var classRoom classRoomEntity.ClassRoom

	dataClassRoom := crr.db.
		Preload("Product").
		Preload("User").
		Where("user_id = ? AND product_id = ?", userId, productId).First(&classRoom)

	if dataClassRoom.Error != nil {
		return nil, dataClassRoom.Error
	}

	return &classRoom, nil
}

func NewClassRoomRepository(db *gorm.DB) ClassRoomRepository {
	return &ClassRoomRepositoryImpl{db}
}
