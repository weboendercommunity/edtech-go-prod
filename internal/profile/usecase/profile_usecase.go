package profile

import (
	profileDto "edtech.id/internal/profile/dto"
	userUsecase "edtech.id/internal/user/usecase"
)

type ProfileUseCase interface {
	GetProfile(id int64) (*profileDto.ProfileResponseBody, error)
}

type ProfileUseCaseImpl struct {
	userUsecase userUsecase.UserUseCase
}

// GetProfile implements ProfileUseCase
func (profileUseCase *ProfileUseCaseImpl) GetProfile(id int64) (*profileDto.ProfileResponseBody, error) {
	// TODO: Get Profile

	user, err := profileUseCase.userUsecase.FindById(id)

	if err != nil {
		return nil, err
	}

	userResponse := profileDto.CreateProfileResponse(*user)

	return &userResponse, nil
}

func NewProfileUseCase(userUsecase userUsecase.UserUseCase) ProfileUseCase {
	return &ProfileUseCaseImpl{userUsecase}
}
