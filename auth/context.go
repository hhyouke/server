package auth

import (
	"net/http"

	"github.com/hhyouke/server/auth/claims"
)

// Context the auth context
type Context struct {
	*Auth
	Claims   *claims.Claims
	Provider Provider
	Request  *http.Request
	Writer   http.ResponseWriter
}
