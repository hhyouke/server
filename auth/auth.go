package auth

import (
	"fmt"
	"strings"

	"github.com/hhyouke/server/auth/tokens"
	"github.com/hhyouke/server/logger"
	"github.com/hhyouke/server/models"
	"github.com/jinzhu/gorm"
)

// Auth authentication main struct
type Auth struct {
	*Config
	providers []Provider
}

// Config the config of auth
type Config struct {
	DB        *gorm.DB
	AuthModel interface{}
	URLPrefix string
	MachineID string
	Logger    *logger.AppLogger

	SignInHandler       func(*Context, func(*Context) (*tokens.Token, error))
	RefreshTokenHandler func(*Context, func(*Context) (*tokens.Token, error))
	SignUpHandler       func(*Context, func(*Context) (*tokens.Token, error))
	SignOutHandler      func(*Context)
}

// New create Auth inst
func New(config *Config) *Auth {
	if config == nil {
		config = &Config{}
	}
	if config.AuthModel == nil {
		config.AuthModel = &models.Auth{}
	}

	if config.URLPrefix == "" {
		config.URLPrefix = "/auth/"
	} else {
		config.URLPrefix = fmt.Sprintf("/%v/", strings.Trim(config.URLPrefix, "/"))
	}

	if config.AuthModel == nil {
		config.AuthModel = &models.Auth{}
	}

	if config.SignInHandler == nil {
		config.SignInHandler = DefaultSignInHandler
	}

	if config.RefreshTokenHandler == nil {
		config.RefreshTokenHandler = DefaultRefreshTokenHandler
	}

	if config.SignUpHandler == nil {
		config.SignUpHandler = DefaultSignUpHandler
	}

	if config.SignOutHandler == nil {
		config.SignOutHandler = DefaultSignOutHandler
	}

	auth := &Auth{Config: config}
	return auth
}

// RegisterProvider register auth provider
func (auth *Auth) RegisterProvider(provider Provider) {
	name := provider.GetName()
	for _, p := range auth.providers {
		if p.GetName() == name {
			fmt.Printf("warning: auth provider %v already registered", name)
			return
		}
	}

	provider.ConfigAuth(auth)
	auth.providers = append(auth.providers, provider)
}

// GetProvider get provider with name
func (auth *Auth) GetProvider(name string) Provider {
	for _, provider := range auth.providers {
		if provider.GetName() == name {
			return provider
		}
	}
	return nil
}

// GetProviders return registered providers
func (auth *Auth) GetProviders() (providers []Provider) {
	for _, provider := range auth.providers {
		providers = append(providers, provider)
	}
	return
}
