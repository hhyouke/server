package encryptor

// Encryptor the encryptor interface
type Encryptor interface {
	Digest(raw string) (string, error)
	Compare(raw string, encrypted string) error
}
