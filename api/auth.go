package api

import (
	"net/http"
	"strings"

	"github.com/hhyouke/server/auth"
	"github.com/hhyouke/server/auth/claims"
)

// Auth endpoint
func (a *API) Auth(rw http.ResponseWriter, req *http.Request) error {

	auth0 := a.Authentication

	var (
		claims  *claims.Claims
		reqPath = strings.TrimPrefix(req.URL.Path, auth0.URLPrefix)
		paths   = strings.Split(reqPath, "/")
		context = &auth.Context{Auth: auth0, Claims: claims, Request: req, Writer: rw}
	)

	if len(paths) == 2 {
		// eg: /phone/login
		if provider := auth0.GetProvider(paths[0]); provider != nil {
			context.Provider = provider
			// serve mux
			switch paths[1] {
			case "login":
				provider.SignIn(context)
			case "refreshToken":
				provider.RefreshToken(context)
			case "logout":
				provider.SignOut(context)
			case "signUp":
				provider.SignUp(context)
			case "callback":
				provider.Callback(context)
			default:
				// provider.ServeHTTP(context)
				http.NotFound(rw, req)
			}
			return nil
		}
	}

	http.NotFound(rw, req)
	return nil
}
