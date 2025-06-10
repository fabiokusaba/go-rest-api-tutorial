package auth

import "golang.org/x/crypto/bcrypt"

func CreateHashedPassword(plainPassword string) (string, error) {
	passwordByte := []byte(plainPassword)

	hashedPassword, err := bcrypt.GenerateFromPassword(passwordByte, bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}

	return string(hashedPassword), nil
}

func ComparePassword(plainPassword, hashedPassword string) error {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plainPassword))
	if err != nil {
		return err
	}

	return nil
}
