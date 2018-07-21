package utils

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// Crypt ...
type Crypt struct {
	cost int
}

// NewCrypt : Get a new crypt struct for encryption and checking
func NewCrypt() *Crypt {
	return &Crypt{cost: bcrypt.DefaultCost}
}

// Hash : Encrypt a string
func (c Crypt) Hash(s string) (string, error) {
	hashed, err := bcrypt.GenerateFromPassword([]byte(s), c.cost)

	if err != nil {
		return "", err
	}

	return string(hashed), nil
}

// Check : check if a string matches an encrypted hash value
func (c Crypt) Check(hashed, pwd string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashed), []byte(pwd))

	fmt.Println(err)
	if err != nil {
		return false
	}

	return true
}
