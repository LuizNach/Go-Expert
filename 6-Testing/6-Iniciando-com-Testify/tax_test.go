package tax

import (
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
		assert.Equal(t, err.Error(), "amount must be greater than 0.0") // Para verificar a mensagem precisamos criar uma verificação manual
	}
	assert.Contains(t, err.Error(), "greater than 0") // Verificamos que o erro contém uma certa string dentro da mensagem
	assert.Equal(t, tax, 0.0)                         // Por fim, verificamos que o valor retornado é do tipo float64(diferente se comparassemos somente com 0) com valor correto
}
