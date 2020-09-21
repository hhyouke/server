package auth

// Provider authentication provider interface
type Provider interface {
	GetName() string

	ConfigAuth(*Auth)
	SignIn(*Context)
	RefreshToken(*Context)
	SignOut(*Context)
	SignUp(*Context)
	Callback(*Context)
	ServeHTTP(*Context)
}
