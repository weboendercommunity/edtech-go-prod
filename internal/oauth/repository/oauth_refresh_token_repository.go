package oauth

import (
	oauthEntity "edtech.id/internal/oauth/entity"
	"gorm.io/gorm"
)

type OauthRefreshTokenRepository interface {
	Create(oauthRefreshToken oauthEntity.OauthRefreshToken) (*oauthEntity.OauthRefreshToken, error)
	FindOneByToken(token string) (*oauthEntity.OauthRefreshToken, error)
	Delete(id int) error
}

type OauthRefreshTokenRepositoryImpl struct {
	db *gorm.DB
}

// Create implements OauthRefreshTokenRepository
func (or *OauthRefreshTokenRepositoryImpl) Create(oauthRefreshToken oauthEntity.OauthRefreshToken) (*oauthEntity.OauthRefreshToken, error) {
	refreshToken := or.db.Create(&oauthRefreshToken)

	if refreshToken.Error != nil {
		return nil, refreshToken.Error
	}

	return &oauthRefreshToken, nil
}

// Delete implements OauthRefreshTokenRepository
func (or *OauthRefreshTokenRepositoryImpl) Delete(id int) error {

	var oauthRefreshToken oauthEntity.OauthRefreshToken

	deleteToken := or.db.Delete(&oauthRefreshToken, id)

	if deleteToken != nil {
		return deleteToken.Error
	}

	return nil
}

// FindOneByToken implements OauthRefreshTokenRepository
func (or *OauthRefreshTokenRepositoryImpl) FindOneByToken(token string) (*oauthEntity.OauthRefreshToken, error) {
	var oauthRefreshToken oauthEntity.OauthRefreshToken

	refreshToken := or.db.Where("token = ?", token).First(&oauthRefreshToken)

	if refreshToken != nil {
		return nil, refreshToken.Error
	}

	return &oauthRefreshToken, nil
}

func NewOauthRefreshTokenRepository(db *gorm.DB) OauthRefreshTokenRepository {
	return &OauthRefreshTokenRepositoryImpl{db}
}
