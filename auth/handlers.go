package auth

import (
	"net/http"

	"github.com/hhyouke/server/auth/claims"
	"github.com/hhyouke/server/auth/tokens"
	"github.com/hhyouke/server/models"
)

func respondAfterLogged(claims *claims.Claims, context *Context) {
}

// DefaultSignInHandler default sign-in behavior
var DefaultSignInHandler = func(context *Context, authenticateHandler func(*Context) (*tokens.Token, error)) {
	var (
		// req         = context.Request
		w          = context.Writer
		token, err = authenticateHandler(context)
	)
	// if everything is ok
	if err == nil && token != nil {
		apiResult := models.NewAPIResult(models.ErrNil, token, "sign-in succeeded")
		models.SendJSON(w, apiResult)
	} else {
		// if it's our app scoped error
		if _, ok := models.ErrText[err.Error()]; ok {
			apiError := models.NewAPIError(err.Error(), "sign-in failed with error: %v", models.ErrText[err.Error()])
			models.SendJSON(w, apiError)
		} else {
			// this is api call failed, internal server error or something something basic
			apiError := models.NewHTTPError(http.StatusInternalServerError, models.ErrInternal, "internal server error: %v", err.Error())
			models.SendJSON(w, apiError)
		}
	}
}

// DefaultRefreshTokenHandler default refresh token handler
var DefaultRefreshTokenHandler = func(context *Context, refreshTokenHandler func(*Context) (*tokens.Token, error)) {
	var (
		w          = context.Writer
		token, err = refreshTokenHandler(context)
	)
	// if everything is ok
	if err == nil && token != nil {
		apiResult := models.NewAPIResult(models.ErrNil, token, "refresh-token succeeded")
		models.SendJSON(w, apiResult)
	} else {
		// if it's our app scoped error
		if _, ok := models.ErrText[err.Error()]; ok {
			apiError := models.NewAPIError(err.Error(), "refresh-token failed with error: %v", models.ErrText[err.Error()])
			models.SendJSON(w, apiError)
		} else {
			// this is api call failed, internal server error or something something basic
			apiError := models.NewHTTPError(http.StatusInternalServerError, models.ErrInternal, "internal server error: %v", err.Error())
			models.SendJSON(w, apiError)
		}
	}
}

// DefaultSignUpHandler default sign-up behavior
var DefaultSignUpHandler = func(context *Context, registerHandler func(*Context) (*tokens.Token, error)) {
	var (
		w          = context.Writer
		token, err = registerHandler(context)
	)
	if err == nil && token != nil {
		apiResult := models.NewAPIResult(models.ErrNil, token, "sign-up succeeded!")
		models.SendJSON(w, apiResult)
	} else {
		// if it's our app scoped error
		if _, ok := models.ErrText[err.Error()]; ok {
			apiError := models.NewAPIError(err.Error(), "sign-up failed with error: %v", models.ErrText[err.Error()])
			models.SendJSON(w, apiError)
		} else {
			// this is api call failed, internal server error or something something basic
			apiError := models.NewHTTPError(http.StatusInternalServerError, models.ErrInternal, "internal server error: %v", err.Error())
			models.SendJSON(w, apiError)
		}
	}
}

// DefaultSignOutHandler default sign-out behavior
var DefaultSignOutHandler = func(context *Context) {
}
