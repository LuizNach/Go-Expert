package database

import "github.com/LuizNach/Go-Expert/7-APIs/10-Criando-UserDB/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindUser(email string) (*entity.User, error)
}
