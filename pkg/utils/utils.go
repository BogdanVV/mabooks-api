package utils

import (
	"crypto/sha256"
	"fmt"
	"os"
)

func CreateHashPassword(password string) string {
	hashedPassword := sha256.New()
	hashedPassword.Write([]byte(password))
	return fmt.Sprintf("%x", hashedPassword.Sum([]byte(os.Getenv("HASH_SALT"))))
}
