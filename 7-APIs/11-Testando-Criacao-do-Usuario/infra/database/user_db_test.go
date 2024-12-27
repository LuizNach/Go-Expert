package database

import (
	"testing"

	"github.com/LuizNach/Go-Expert/7-APIs/11-Testando-Criacao-do-Usuario/internal/entity"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func TestCreateNewuser(t *testing.T) {

	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{})
	if err != nil {
		t.Error(err)
	}
	db.AutoMigrate(&entity.User{}) // create the table: User from user_db.go

	// Create a new user
	user, err := entity.NewUser("John Doe", "email@email.com", "123")
	if err != nil {
		t.Error(err)
	}

	userDB := NewUser(db) // struct User com um campo DB do tipo gorm.DB para conexão com o banco de dados

	err = userDB.CreateUser(user)
	// if err != nil {
	// 	t.Error(err)
	// }
	assert.Nil(t, err)

	var userFound entity.User
	err = db.First(&userFound).Where("id = ?", user.ID).Error
	assert.Nil(t, err)
	assert.Equal(t, user.ID, userFound.ID)
	assert.Equal(t, user.Name, userFound.Name)
	assert.Equal(t, user.Email, userFound.Email)
	assert.NotNil(t, userFound.Password)

}
