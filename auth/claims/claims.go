package claims

import (
	"github.com/dgrijalva/jwt-go"
)

// Claims auth claims
type Claims struct {
	Provider string `json:"provider,omitempty"`
	SignOrg  string `json:"org,omitempty"`
	// UserID   string `json:"userid,omitempty"`
	jwt.StandardClaims
}

// Claimer the claimer interface
type Claimer interface {
	ToClaims() *Claims
}

// ToClaims Claims implments the claimer interface
func (c *Claims) ToClaims() *Claims {
	return c
}
