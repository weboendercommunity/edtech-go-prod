package oauth

import "github.com/golang-jwt/jwt/v5"

type LoginResponseBody struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	Type         string `json:"type"`
	ExpiredAt    string `json:"expired_at"`
	Scope        string `json:"scope"`
}

type UserResponseBody struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type ClaimsResponseBody struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	IsAdmin bool   `json:"is_admin,omitempty"`
	jwt.RegisteredClaims
}

type MapClaimsResponseBody struct {
	ID            int64  `json:"id"`
	Name          string `json:"name"`
	Email         string `json:"email"`
	IsAdmin       bool   `json:"is_admin,omitempty"`
	jwt.MapClaims `json:"omitempty"`
}
