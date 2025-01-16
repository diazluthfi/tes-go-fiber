package utils

import "golang.org/x/crypto/bcrypt"

// HashingPassword hashes a given password using bcrypt.
func HashingPassword(password string) (string, error) {
	// Hash the password with bcrypt and a cost factor of 14
	hashedByte, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		return "", err // Return an empty string and the error if hashing fails
	}
	return string(hashedByte), nil // Return the hashed password as a string
}

func CheckPasswordHash(password, hashedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))

	return err == nil
}
