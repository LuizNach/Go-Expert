package entity

import (
	"log"

	"github.com/LuizNach/Go-Expert/7-APIs/7-Testando-user/pkg/entity"
	"golang.org/x/crypto/bcrypt"
)

// VO - Value Object
type User struct {
	ID       entity.ID `json:"id"`
	Name     string    `json:"name"`
	Email    string    `json:"email"`
	Password string    `json:"-"` // MARK Não entendi bem essa parte
}

func NewUser(name string, email string, password string) (*User, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	return &User{
		ID:       entity.NewUuid(),
		Name:     name,
		Email:    email,
		Password: string(hash), // Nunca guardamos a senha de forma pura, sempre incriptamos para guardar nodb
	}, nil
}

func (u *User) ValidatePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password)) // Se bater tudo certo, retornará nil pois não ha erros
	if err != nil {
		log.Printf("Error on checking bcrypt hash validation on user: %s", u.ID)
	}
	return err == nil
}
