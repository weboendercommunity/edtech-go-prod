package oauth

import (
	oauthEntity "edtech.id/internal/oauth/entity"
	"gorm.io/gorm"
)

type OauthAccessTokenRepository interface {
	Create(oauthAccessToken oauthEntity.OauthAccessToken) (*oauthEntity.OauthAccessToken, error)
	Delete(id int) error
}

type OauthAccessTokenRepositoryImpl struct {
	db *gorm.DB
}

// Create implements OauthAccessTokenRepository
func (oa *OauthAccessTokenRepositoryImpl) Create(oauthAccessToken oauthEntity.OauthAccessToken) (*oauthEntity.OauthAccessToken, error) {
	createToken := oa.db.Create(&oauthAccessToken)

	if createToken.Error != nil {
		return nil, createToken.Error
	}
	return &oauthAccessToken, nil
}

// Delete implements OauthAccessTokenRepository
func (oa *OauthAccessTokenRepositoryImpl) Delete(id int) error {
	var oauthAccessToken oauthEntity.OauthAccessToken

	deleteToken := oa.db.Delete(&oauthAccessToken, id)

	if deleteToken.Error != nil {
		return deleteToken.Error
	}

	return nil

}

func NewOauthAccessTokenRepository(db *gorm.DB) OauthAccessTokenRepository {
	return &OauthAccessTokenRepositoryImpl{db}
}
