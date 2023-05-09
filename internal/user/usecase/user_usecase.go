package user

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	userDto "edtech.id/internal/user/dto"
	userEntity "edtech.id/internal/user/entity"
	userRepository "edtech.id/internal/user/repository"
	utils "edtech.id/pkg/utils"
)

type UserUseCase interface {
	FindAll(offset int, limit int) []userEntity.User
	FindById(id int64) (*userEntity.User, error)
	FindByEmail(email string) (*userEntity.User, error)
	Create(userDto userDto.UserRequestBody) (*userEntity.User, error)
	Update(userDto userDto.UserRequestBody) (*userEntity.User, error)
	Delete(id int) error
}

type UserUseCaseImpl struct {
	userRepository userRepository.UserRepository
}

func NewUserUseCase(ur userRepository.UserRepository) UserUseCase {
	return &UserUseCaseImpl{ur}
}

// FindAll implements UserUseCase
func (*UserUseCaseImpl) FindAll(offset int, limit int) []userEntity.User {
	panic("unimplemented")
}

// FindByEmail implements UserUseCase
func (uu *UserUseCaseImpl) FindByEmail(email string) (*userEntity.User, error) {
	return uu.userRepository.FindByEmail(email)
}

// FindById implements UserUseCase
func (uu *UserUseCaseImpl) FindById(id int64) (*userEntity.User, error) {
	return uu.userRepository.FindById(id)
}

// Create implements UserUseCase
func (uu *UserUseCaseImpl) Create(userDto userDto.UserRequestBody) (*userEntity.User, error) {
	checkUser, err := uu.userRepository.FindByEmail(*userDto.Email)

	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, err
	}

	if checkUser != nil {
		return nil, errors.New("email already exist")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(*userDto.Password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	user := userEntity.User{
		Name:         *userDto.Name,
		Email:        *userDto.Email,
		Password:     string(hashedPassword),
		CodeVerified: utils.RandString(32),
	}

	dataUser, err := uu.userRepository.Create(user)

	if err != nil {
		return nil, err
	}

	return dataUser, nil
}

// Update implements UserUseCase
func (*UserUseCaseImpl) Update(userDto userDto.UserRequestBody) (*userEntity.User, error) {
	panic("unimplemented")
}

// Delete implements UserUseCase
func (*UserUseCaseImpl) Delete(id int) error {
	panic("unimplemented")
}
