package user

import (
	userEntity "edtech.id/internal/user/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	FindAll(offset int, limit int) []userEntity.User
	FindById(id int64) (*userEntity.User, error)
	FindByEmail(email string) (*userEntity.User, error)
	Create(user userEntity.User) (*userEntity.User, error)
	Update(user userEntity.User) (*userEntity.User, error)
	Delete(user userEntity.User) error
}

type UserRepositoryImpl struct {
	db *gorm.DB
}

// FindAll implements UserRepository
func (ur *UserRepositoryImpl) FindAll(offset int, limit int) []userEntity.User {
	var users []userEntity.User

	ur.db.Find(&users)

	return users
}

// FindByEmail implements UserRepository
func (ur *UserRepositoryImpl) FindByEmail(email string) (*userEntity.User, error) {
	var user userEntity.User

	err := ur.db.Where("email = ?", email).First(&user)

	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

// FindById implements UserRepository
func (ur *UserRepositoryImpl) FindById(id int64) (*userEntity.User, error) {
	var user userEntity.User

	err := ur.db.First(&user, id)

	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

// Create implements UserRepository
func (ur *UserRepositoryImpl) Create(user userEntity.User) (*userEntity.User, error) {
	err := ur.db.Create(&user)

	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

// Update implements UserRepository
func (ur *UserRepositoryImpl) Update(user userEntity.User) (*userEntity.User, error) {
	err := ur.db.Save(&user)

	if err.Error != nil {
		return nil, err.Error
	}

	return &user, nil
}

// Delete implements UserRepository
func (ur *UserRepositoryImpl) Delete(user userEntity.User) error {
	err := ur.db.Save(&user)

	if err.Error != nil {
		return err.Error
	}

	return nil
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &UserRepositoryImpl{db}
}
