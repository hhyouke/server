package bcrypts

import "golang.org/x/crypto/bcrypt"

// Config BcryptEncryptor config
type Config struct {
	Cost int
}

// Encryptor BcryptEncryptor struct
type Encryptor struct {
	Config *Config
}

// New initalize BcryptEncryptor
func New(config *Config) *Encryptor {
	if config == nil {
		config = &Config{}
	}

	if config.Cost == 0 {
		config.Cost = bcrypt.DefaultCost
	}

	return &Encryptor{
		Config: config,
	}
}

// Digest generate encrypted password
func (bcryptEncryptor *Encryptor) Digest(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(hashedPassword), err
}

// Compare check hashed password
func (bcryptEncryptor *Encryptor) Compare(raw string, enc string) error {
	return bcrypt.CompareHashAndPassword([]byte(enc), []byte(raw))
}
