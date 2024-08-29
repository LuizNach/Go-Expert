package tax

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCalculateTax(t *testing.T) {
	tax, err := CalculateTax(1000.00) // Fazendo que a funcao possa retornar o valor e o erro podemos analisar de maneiras diferentes ambos
	assert.Nil(t, err)                // Testamos que o retorno da função não exibiu o campo de erro
	assert.Equal(t, tax, 10.0)        // Testamos que retorno do valor da função retorne o valor correto

	tax, err = CalculateTax(0) // Novo teste para caso que o erro não é nulo
	// assert.NotNil(t, err)                                // Garantimos que o erro não é nulo
	assert.Error(t, err, "mount must be greater than 0") // Verificamos que o erro é diferente de nil

	// Reference https://pkg.go.dev/github.com/stretchr/testify@v1.9.0/assert#Error
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "amount must be greater than 0.0") // Para verificar a mensagem precisamos criar uma verificação manual
	}
	assert.Contains(t, err.Error(), "greater than 0") // Verificamos que o erro contém uma certa string dentro da mensagem
	assert.Equal(t, tax, 0.0)                         // Por fim, verificamos que o valor retornado é do tipo float64(diferente se comparassemos somente com 0) com valor correto
}

func TestCalculateTaxAndSave(t *testing.T) {

	repository := &TaxRepositoryMock{}
	// ARRANGE
	repository.On("SaveTax", 10.0).Return(nil) // Quando o metodo "SaveTax" for chamado com o argumento amount=10.0 retorne error=nil

	// ACT
	err := CalculateTaxAndSave(1000.0, repository)

	// ASSERT
	assert.Nil(t, err)
	repository.AssertExpectations(t) // Aqui fazemos a garantia de que os metodos mockados em ARRANGE foram chamados. Caso o contrário ocorrerá erro no teste.

	// ----------------------------------------------------------------------------------------------------------------------------------

	// ARRANGE
	repository.On("SaveTax", 0.0).Return(errors.New("error saving tax")) // Quando o "SaveTax" for chamado com parametro amount 0.0 lançará error

	// ACT
	err = CalculateTaxAndSave(0.0, repository)

	// ASSERT
	// Verificamos que o erro foi emitido
	if assert.Error(t, err) {
		assert.Contains(t, err.Error(), "error saving tax")
	}
	repository.AssertExpectations(t)
	repository.AssertNumberOfCalls(t, "SaveTax", 2) // Garanto nessa execução do teste todo, a função "SaveTax" foi chamada exatamente 2 vezes
}
