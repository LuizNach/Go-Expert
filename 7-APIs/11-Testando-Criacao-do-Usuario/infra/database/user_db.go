package database

import (
	"github.com/LuizNach/Go-Expert/7-APIs/11-Testando-Criacao-do-Usuario/internal/entity"
	"gorm.io/gorm"
)

type User struct { // struct User com um campo DB do tipo gorm.DB para conexão com o banco de dados
	DB *gorm.DB // conexão com o banco de dados
}

func NewUser(db *gorm.DB) *User { // nova conexão com o banco de dados, tabela User
	return &User{
		DB: db,
	}
}

func (u *User) CreateUser(user *entity.User) error { // Cria uma entidade User na tabela User
	return u.DB.Create(user).Error
}

func (u *User) FindUserByEmail(email string) (*entity.User, error) {
	var user entity.User
	err := u.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}
