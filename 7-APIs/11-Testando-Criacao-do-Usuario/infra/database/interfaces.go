package database

import "github.com/LuizNach/Go-Expert/7-APIs/11-Testando-Criacao-do-Usuario/internal/entity"

type UserInterface interface {
	CreateUser(user *entity.User) error
	FindUserByEmail(email string) (*entity.User, error)
}
