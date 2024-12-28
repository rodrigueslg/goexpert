package entity

import (
	"golang.org/x/crypto/bcrypt"

	pkgEntity "github.com/rodrigueslg/fc-goexpert/api/pkg/entity"
)

type User struct {
	ID       pkgEntity.ID `json:"id"`
	Name     string       `json:"name"`
	Email    string       `json:"email"`
	Password string       `json:"-"`
}

func NewUser(name, email, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	return &User{
		ID:       pkgEntity.NewID(),
		Name:     name,
		Email:    email,
		Password: string(hash),
	}, nil
}

func (u *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
