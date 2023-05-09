package oauth

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	oauthDto "edtech.id/internal/oauth/dto"
	oauthEntity "edtech.id/internal/oauth/entity"
	oauthRepository "edtech.id/internal/oauth/repository"
	userUseCase "edtech.id/internal/user/usecase"

	utils "edtech.id/pkg/utils"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

type OauthUseCase interface {
	Login(loginRequestBody oauthDto.LoginRequestBody) (*oauthDto.LoginResponseBody, error)
	Refresh(refreshTokenRequestBody oauthDto.RefreshTokenRequestBody) (*oauthDto.LoginResponseBody, error)
}

type OauthUseCaseImpl struct {
	oauthClientRepository       oauthRepository.OauthClientRepository
	oauthAccessTokenRepository  oauthRepository.OauthAccessTokenRepository
	oauthRefreshTokenRepository oauthRepository.OauthRefreshTokenRepository
	userUseCase                 userUseCase.UserUseCase
}

// Login implements OauthUseCase
func (ou *OauthUseCaseImpl) Login(loginRequestBody oauthDto.LoginRequestBody) (*oauthDto.LoginResponseBody, error) {

	oauthClient, err := ou.oauthClientRepository.FindByClientIdAndClientSecret(loginRequestBody.ClientID, loginRequestBody.ClientSecret)

	if err != nil {
		return nil, err
	}

	var user oauthDto.UserResponseBody

	// Login user

	dataUser, err := ou.userUseCase.FindByEmail(loginRequestBody.Email)

	if err != nil {
		return nil, errors.New("invalid email")
	}

	user.ID = dataUser.ID
	user.Email = dataUser.Email
	user.Name = dataUser.Name
	user.Password = dataUser.Password

	jwtKey := []byte(os.Getenv("JWT_SECRET"))

	bcryptErr := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginRequestBody.Password))

	if bcryptErr != nil {
		return nil, errors.New("invalid password")
	}

	expirationTime := time.Now().Add(24 * 365 * time.Hour)

	claims := &oauthDto.ClaimsResponseBody{
		ID:      user.ID,
		Name:    user.Name,
		Email:   user.Email,
		IsAdmin: false,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtKey)

	if err != nil {
		return nil, err
	}

	// Insert to oauth_access_token

	dataOauthAccessToken := oauthEntity.OauthAccessToken{
		OauthClientID: &oauthClient.ID,
		UserID:        user.ID,
		AccessToken:   tokenString,
		Scope:         "*",
		ExpiredAt: sql.NullTime{
			Time: expirationTime,
		},
	}

	oauthAccessToken, err := ou.oauthAccessTokenRepository.Create(dataOauthAccessToken)

	if err != nil {
		return nil, err
	}

	// Insert to oauth_refresh_token

	dataOauthRefreshToken := oauthEntity.OauthRefreshToken{
		OauthAccessTokenID: &oauthAccessToken.ID,
		UserID:             user.ID,
		RefreshToken:       utils.RandString(128),
		ExpiredAt: sql.NullTime{
			Time: time.Now().Add(24 * 366 * time.Hour),
		},
	}

	oauthRefreshToken, err := ou.oauthRefreshTokenRepository.Create(dataOauthRefreshToken)

	if err != nil {
		fmt.Println("this is error", err)
		return nil, err
	}

	return &oauthDto.LoginResponseBody{
		AccessToken:  tokenString,
		RefreshToken: oauthRefreshToken.RefreshToken,
		Type:         "Bearer",
		ExpiredAt:    expirationTime.Format(time.RFC3339),
		Scope:        "*",
	}, nil

}

// Refresh implements OauthUseCase
func (*OauthUseCaseImpl) Refresh(refreshTokenRequestBody oauthDto.RefreshTokenRequestBody) (*oauthDto.LoginResponseBody, error) {
	panic("unimplemented")
}

func NewOauthUseCase(
	oauthClientRepository oauthRepository.OauthClientRepository,
	oauthAccessTokenRepository oauthRepository.OauthAccessTokenRepository,
	oauthRefreshTokenRepository oauthRepository.OauthRefreshTokenRepository,
	userUseCase userUseCase.UserUseCase) OauthUseCase {
	return &OauthUseCaseImpl{
		oauthClientRepository,
		oauthAccessTokenRepository,
		oauthRefreshTokenRepository,
		userUseCase,
	}
}
