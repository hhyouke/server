package password

import (
	"github.com/hhyouke/server/auth"
	encryptor "github.com/hhyouke/server/auth/providers/password/encrypt"
	"github.com/hhyouke/server/auth/providers/password/encrypt/bcrypts"
	"github.com/hhyouke/server/auth/tokens"
)

// Config password authentication provider config
type Config struct {
	signUphandler       func(*auth.Context) (*tokens.Token, error)
	authenticateHandler func(*auth.Context) (*tokens.Token, error)
	refreshTokenHandler func(*auth.Context) (*tokens.Token, error)
	Encryptor           encryptor.Encryptor
}

// Provider is the password way to implements *auth.Provider interface
type Provider struct {
	*Config
}

// New initialize password authentication provider
func New(config *Config) *Provider {

	if config == nil {
		config = &Config{}
	}

	if config.Encryptor == nil {
		config.Encryptor = bcrypts.New(&bcrypts.Config{})
	}

	provider := &Provider{Config: config}

	if config.authenticateHandler == nil {
		config.authenticateHandler = PasswordAuthenticateHandler
	}

	if config.refreshTokenHandler == nil {
		config.refreshTokenHandler = PasswordRefreshTokenHandler
	}

	if config.signUphandler == nil {
		config.signUphandler = PasswordSignUpHandler
	}

	return provider
}

// GetName return the password provider name
func (p Provider) GetName() string {
	return "pwd"
}

// ConfigAuth config auth
func (p Provider) ConfigAuth(auth *auth.Auth) {
}

// SignIn implemented sign-in with password provider
func (p Provider) SignIn(context *auth.Context) {
	context.Auth.SignInHandler(context, p.authenticateHandler)
}

// RefreshToken implemented refreshToken method
func (p Provider) RefreshToken(context *auth.Context) {
	context.Auth.RefreshTokenHandler(context, p.refreshTokenHandler)
}

// SignOut implemented sign-out with password provider
func (p Provider) SignOut(context *auth.Context) {

}

// SignUp implemented sign-up with password provider
func (p Provider) SignUp(context *auth.Context) {
	context.Auth.SignUpHandler(context, p.signUphandler)
}

// Callback implement Callback with password provider
func (p Provider) Callback(context *auth.Context) {

}

// ServeHTTP implement ServeHTTP with password provider
func (p Provider) ServeHTTP(context *auth.Context) {

}
