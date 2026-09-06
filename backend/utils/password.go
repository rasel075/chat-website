package utils

import (
	"golang.org/x/crypto/bcrypt"
)

// HashPassword converts a plain text password into a bcrypt hash
func HashPassword(password string) (string, error) {
	// Cost determines how much CPU time is needed to compute the hash
	// Higher = more secure but slower (10 is standard)
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// CheckPassword compares a plain text password with its hash
func CheckPassword(hashedPassword string, password string) bool {
	// Returns nil if password matches, error if it doesn't
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	return err == nil
}

/*Explanation:

Function	Explanation
HashPassword()	Takes a plain password, returns secure hash
bcrypt.GenerateFromPassword()	Creates the hash
CheckPassword()	Compares a plain password to a hash
bcrypt.CompareHashAndPassword()	Returns nil if match, error if not
*/