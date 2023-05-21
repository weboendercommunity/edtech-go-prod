package class_room

import (
	"errors"

	classRoomDto "edtech.id/internal/class_room/dto"
	classRoomEntity "edtech.id/internal/class_room/entity"
	classRoomRepository "edtech.id/internal/class_room/repository"
	"gorm.io/gorm"
)

type ClassRoomUsecase interface {
	Create(classRoomDto classRoomDto.ClassRoomRequestBody) (*classRoomEntity.ClassRoom, error)
	FindAllByUserID(offset int, limit int, userId int) (classRoomDto.ClassRoomListResponse, error)
	// FindOneByUserIDAndProductID(userId int, productId int) (*classRoomDto.ClassRoomResponseBody, error)
}

type ClassRoomUsecaseImpl struct {
	classRoomRepository classRoomRepository.ClassRoomRepository
}

// Create implements ClassRoomUsecase
func (cru *ClassRoomUsecaseImpl) Create(classRoomDto classRoomDto.ClassRoomRequestBody) (*classRoomEntity.ClassRoom, error) {
	dataClassRoom, err := cru.classRoomRepository.FindOneByUserIDAndProductID(
		int(classRoomDto.UserID),
		int(classRoomDto.ProductID),
	)

	// if the class room is not found, but there is an error, return the error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	// if the class room is already exists, return the error
	if dataClassRoom != nil {
		return nil, errors.New("class room already exists")
	}

	// create the class room
	classRoom := classRoomEntity.ClassRoom{
		ProductID: classRoomDto.ProductID,
		UserID:    classRoomDto.UserID}

	createdClassRoom, err := cru.classRoomRepository.Create(classRoom)

	if err != nil {
		return nil, err
	}

	return createdClassRoom, nil

}

// FindAllByUserID implements ClassRoomUsecase
func (cru *ClassRoomUsecaseImpl) FindAllByUserID(offset int, limit int, userId int) (classRoomDto.ClassRoomListResponse, error) {
	classRooms, err := cru.classRoomRepository.FindAllByUserID(offset, limit, userId)

	if err != nil {
		return classRoomDto.ClassRoomListResponse{}, err
	}

	classRoomListResponse := classRoomDto.CreateClassRoomListResponse(classRooms)

	return classRoomListResponse, nil

}

func NewClassRoomUsecase(
	classRoomRepository classRoomRepository.ClassRoomRepository,
) ClassRoomUsecase {
	return &ClassRoomUsecaseImpl{classRoomRepository}
}
