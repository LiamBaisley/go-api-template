package security

import (
	"golang.org/x/crypto/bcrypt"
)

const (
	MinCost     int = 4
	Maxcost     int = 31
	DefaultCost int = 10
)

func CompareHashAndPassword(hashedPassword []byte, password []byte) error {
	return bcrypt.CompareHashAndPassword(hashedPassword, password)
}

func GetCost(hashedPassword []byte) (int, error) {
	return bcrypt.Cost(hashedPassword)
}

func HashPassword(password []byte) ([]byte, error) {
	return bcrypt.GenerateFromPassword(password, DefaultCost)
}
