package helper

import "golang.org/x/crypto/bcrypt"

func ComparePassword(hashedPassword, password string) bool {
	pass := []byte(password)
	hash := []byte(hashedPassword)

	err := bcrypt.CompareHashAndPassword(hash, pass)
	return err != nil
}
