package tokens

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/hhyouke/server/models"
)

// TokenSignSecret the sign secret using the HMAC algorithm for now
const TokenSignSecret = "50TIN-APP-TOKEN-SIGN-SECRET"

// Token is the composite of access-token and refresh-token
type Token struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func GenerateTokenPair(auth *models.Auth) (*Token, error) {
	// Create access token
	claims := auth.ToClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(TokenSignSecret))
	if err != nil {
		return nil, err
	}
	// create refresh token
	refreshClaims := auth.ToRefreshClaims()
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	rt, err := refreshToken.SignedString([]byte(TokenSignSecret))
	if err != nil {
		return nil, err
	}

	return &Token{
		AccessToken:  t,
		RefreshToken: rt,
	}, nil
}
