# Fuzz Testing
  
## O que é Fuzzing ou Fuzz Testing?
Fuzzing é uma técnica de teste que envolve fornecer dados de entrada gerados aleatoriamente para uma aplicação para descobrir possíveis falhas ou comportamentos inesperados. Em Go, isso pode ser usado para testar funções e encontrar bugs que podem não ser detectados por testes unitários tradicionais.  
  
## Como o Fuzzing Funciona?
No fuzzing, você especifica a função que deseja testar, e o fuzzing engine (motor de fuzzing) gera uma variedade de entradas aleatórias para essa função. O objetivo é explorar o maior número possível de caminhos de execução, tentando causar falhas, panics, ou comportamentos inesperados.  
  
## Introdução ao Fuzzing em Go
Desde a versão 1.18, o Go suporta fuzzing integrado.  
  
## Exemplo Simples de Fuzzing
Imagine que você tem uma função simples que inverte uma string:
  
```golang
package main

import (
	"strings"
)

// Função que inverte uma string
func Reverse(s string) string {
	var sb strings.Builder
	for i := len(s) - 1; i >= 0; i-- {
		sb.WriteByte(s[i])
	}
	return sb.String()
}
```
Agora, queremos testar essa função para garantir que ela sempre se comporta conforme o esperado, independentemente da entrada.  
  
## Criando um Teste de Fuzzing
Crie um arquivo `reverse_test.go`:  
  
```go
package main

import (
	"testing"
)

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // Adiciona casos de teste conhecidos
	}
	
	f.Fuzz(func(t *testing.T, orig string) {
		rev := Reverse(orig)
		doubleRev := Reverse(rev)
		if orig != doubleRev {
			t.Errorf("Original: %q, Dupla Reversão: %q", orig, doubleRev)
		}

		if strings.ContainsAny(rev, "\x00") {
			t.Errorf("Reversão contém byte nulo: %q", rev)
		}
	})
}

```
  
## Explicação do Código
1. **Função FuzzReverse**:  
    * f *testing.F: Recebe o objeto testing.F para configurar o fuzzing.  
    * f.Add(tc): Adiciona alguns casos de teste conhecidos. Isso ajuda a garantir que o fuzzing comece com entradas válidas e comuns.  
    * f.Fuzz(func(t *testing.T, orig string)): Esta é a função principal onde o fuzzing ocorre. A função recebe uma string gerada aleatoriamente e aplica a função Reverse a ela.  
2. **Dupla Reversão**:  
    * Se você reverter uma string duas vezes, ela deve ser igual à original. Este é um teste simples para verificar a correção da função.  
3. **Verificação de Byte Nulo**:  
    * Aqui verificamos se a reversão introduz caracteres nulos inesperados.  

## Rodando o Fuzzing  
Para executar o fuzzing, você pode usar o seguinte comando:  
```sh
go test -fuzz=FuzzReverse
```
Para rodar todos os testes Fuzz do repositório atual:
```sh
go test -run=^# -fuzz=.
```  
A flag `-run=` recebe um regex para saber quais testes deve rodar em conjunto e se passarmos um regex de nome inválido garantimos que nenhum outro teste fora os Fuzz serão rodados.  
  
O fuzzing continuará a gerar dados até que você o interrompa ou até que encontre um bug.  
  
Após a execução dos dados adicionados em f.Add, a massa principal executada, se o fuzz testing capturar algum erro ele gerará um diretório de testdata. Esse diretório conterá o relatório fuzz em relação a função executada e os casos  de teste que falharam com os inputs utilizados. O próprio CLIdo `go test` irá mostrar o nome do arquivo de teste que falhou, o arquivo criado que contem o input data do teste e o comando para re rodar esses mesmo input data:
```sh
fuzz: elapsed: 0s, gathering baseline coverage: 0/7 completed
fuzz: elapsed: 0s, gathering baseline coverage: 7/7 completed, now fuzzing with 6 workers
fuzz: elapsed: 0s, execs: 77 (1573/sec), new interesting: 0 (total: 7)
--- FAIL: FuzzCalculateTax (0.05s)
    --- FAIL: FuzzCalculateTax (0.00s)
        tax_test.go:79: Received 20.000000 but expected 10
    
    Failing input written to testdata/fuzz/FuzzCalculateTax/0a5de64eb02dc213
    To re-run:
    go test -run=FuzzCalculateTax/0a5de64eb02dc213
FAIL
exit status 1
FAIL    github.com/LuizNach/Go-Expert/6-Testing/5-Fuzzing       0.051s
```
  
## Casos de Uso do Fuzzing
  
1. **Teste de Robustez**: Ideal para testar a robustez de funções que processam entradas complexas, como parsers ou funções de manipulação de strings.  
2. **Descoberta de Vulnerabilidades**: Pode ser usado para encontrar vulnerabilidades de segurança, como estouro de buffer, desreferenciamento de ponteiro nulo, etc.  
3. **Melhoria de Testes Existentes**: Fuzzing pode complementar testes unitários, explorando caminhos que os testes escritos manualmente podem não cobrir.  
4. **Testes de Integração**: Fuzzing pode ser aplicado em testes de integração para testar a interação de vários componentes.  

## Vantagens do Fuzzing
* **Cobertura Extensa**: Pode explorar partes do código que normalmente não são cobertas por testes tradicionais.
* **Automático**: O processo de geração de entradas é automático, o que significa menos trabalho manual.
* **Eficiência**: Pode encontrar bugs que são difíceis de detectar usando técnicas de teste convencionais.

## Limitações do Fuzzing
* **Falso Positivo/Negativo**: Como qualquer teste automático, o fuzzing pode produzir falsos positivos ou negativos.
* **Consome Tempo**: Pode ser demorado, especialmente se for aplicado a funções que levam muito tempo para processar entradas.
* **Reprodução de Erros**: Algumas entradas que causam erros podem ser difíceis de reproduzir manualmente.
  
## Conclusão
O fuzzing é uma ferramenta poderosa para melhorar a qualidade e a segurança do código. No contexto do Go, ele pode ser facilmente integrado aos testes existentes, proporcionando uma camada adicional de proteção contra bugs e vulnerabilidades. O exemplo que vimos mostra como ele pode ser usado para testar uma simples função de reversão de string, mas a técnica pode ser aplicada a qualquer função ou sistema que precise ser robusto contra entradas inesperadas.  
