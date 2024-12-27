package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	p, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, p)
	assert.NotEmpty(t, p)
	assert.Equal(t, "Product 1", p.Name)
	assert.Equal(t, 10, p.Price)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	p, err := NewProduct("", 10)
	assert.NotNil(t, err)
	assert.Nil(t, p)

	assert.EqualError(t, err, ErrNameIsRequired.Error()) // detalhe argumentos: error, string
	// OR
	assert.Equal(t, ErrNameIsRequired, err) // detalhe argumentos: error, error
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	p, err := NewProduct("Product 2", 0)
	assert.NotNil(t, err)
	assert.Nil(t, p)

	assert.EqualError(t, err, ErrPriceIsRequired.Error())
	// OR
	assert.Equal(t, ErrPriceIsRequired, err)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	p, err := NewProduct("Product 2", -10)
	assert.NotNil(t, err)
	assert.Nil(t, p)

	assert.EqualError(t, err, ErrInvalidPrice.Error())
	// OR
	assert.Equal(t, ErrInvalidPrice, err)
}
