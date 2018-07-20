package utils

import "golang.org/x/crypto/bcrypt"

type crypt struct {
	cost int
}

// Get a new crypt struct for encryption and checking
func newCrypt() *crypt {
	return &crypt{cost: bcrypt.DefaultCost}
}

// Encrypt a string
func (c crypt) hash(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), c.cost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// Check if a string matches an encrypted hash value
func (c crypt) check(hashed, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd))

	if err != nil {
		return false
	}

	return true
}
