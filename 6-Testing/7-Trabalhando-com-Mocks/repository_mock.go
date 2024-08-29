package tax

import (
	"github.com/stretchr/testify/mock"
)

type TaxRepositoryMock struct {
	mock.Mock
}

func (m *TaxRepositoryMock) SaveTax(tax float64) error {
	args := m.Called(tax)
	return args.Error(0) // Aqui colocamos os args para retornar Error pq a função SaveTax retorna error. Se a função retornasse int, retornariamos args.Int()
}
