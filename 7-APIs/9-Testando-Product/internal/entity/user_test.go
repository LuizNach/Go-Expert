package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewuser(t *testing.T) {
	user, err := NewUser("Some New User Name", "new.user@email.com", "123456")
	assert.Nil(t, err)
	assert.NotNil(t, user)
	assert.NotEmpty(t, user.ID)
	assert.NotEmpty(t, user.Email)
	assert.NotEmpty(t, user.Password)

	assert.Equal(t, user.Name, "Some New User Name")
	assert.Equal(t, user.Email, "new.user@email.com")
}

func TestUserValidatePassword(t *testing.T) {
	user, err := NewUser("John Doe", "a@a.com", "123456")
	assert.Nil(t, err)
	assert.NotEqual(t, user.Password, "123456") // Garantimos que o que está sendo guardado nunca é o password puro
	assert.True(t, user.ValidatePassword("123456"))
	assert.False(t, user.ValidatePassword("1234567"))
}
