package oauth

import (
	oauth "edtech.id/internal/oauth/entity"
	"gorm.io/gorm"
)

type OauthClientRepository interface {
	FindByClientIdAndClientSecret(clientId string, clientSecret string) (*oauth.OauthClient, error)
}

type OauthClientRepositoryImpl struct {
	db *gorm.DB
}

// FindByClientIdAndClientSecret implements OauthClientRepository
func (oc *OauthClientRepositoryImpl) FindByClientIdAndClientSecret(clientId string, clientSecret string) (*oauth.OauthClient, error) {
	var oauthClient oauth.OauthClient

	client := oc.db.Where("client_id = ?", clientId).Where("client_secret = ?", clientSecret).First(&oauthClient)

	if client.Error != nil {
		return nil, client.Error
	}

	return &oauthClient, nil
}

func NewOauthClientRepository(db *gorm.DB) OauthClientRepository {
	return &OauthClientRepositoryImpl{db}
}
