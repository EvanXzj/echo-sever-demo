package utils

import "golang.org/x/crypto/bcrypt"

type crypt struct {
	cost int
}

// NewCrypt : Get a new crypt struct for encryption and checking
func NewCrypt() *crypt {
	return &crypt{cost: bcrypt.DefaultCost}
}

// Hash : Encrypt a string
func (c crypt) Hash(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), c.cost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// Check: check if a string matches an encrypted hash value
func (c crypt) Check(hashed, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd))

	if err != nil {
		return false
	}

	return true
}
