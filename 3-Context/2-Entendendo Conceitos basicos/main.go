package main

import (
	"context"
	"fmt"
	"time"
)

/*
   Em Go, o pacote context é usado para facilitar o cancelamento de operações de longa duração, gerenciamento de deadlines e
   valores associados a uma determinada goroutine. Ele é muito útil em situações onde é necessário controlar o tempo de execução
   de uma operação ou propagar valores específicos através de chamadas de função.

   O contexto em Go é uma estrutura que carrega informações sobre a execução de uma determinada goroutine e permite que essas
   informações sejam propagadas através das chamadas de função. Ele é utilizado para cancelar operações de forma segura, evitando
   vazamento de recursos e garantindo uma execução controlada do programa.

   Por exemplo, ao realizar uma requisição HTTP em Go, é possível passar um contexto que contenha um timeout para garantir que a
   requisição não fique esperando indefinidamente. Além disso, o contexto também pode ser utilizado para propagar valores
   específicos, como dados de autenticação, transações de banco de dados, entre outros.

   Em resumo, o contexto em Go é uma ferramenta poderosa para controlar o fluxo de execução de um programa de forma segura e
   eficiente.

   package main

   import (
       "context"
       "fmt"
       "time"
   )

   func main() {
       // Criando um contexto de background
       ctx := context.Background()

       // Criando um contexto com cancelamento
       ctxWithCancel, cancel := context.WithCancel(ctx)

       // Chamando uma função que executa em uma goroutine
       go func(ctx context.Context) {
           for {
               select {
               case <-ctx.Done():
                   fmt.Println("Goroutine cancelada")
                   return
               default:
                   fmt.Println("Executando goroutine...")
                   time.Sleep(1 * time.Second)
               }
           }
       }(ctxWithCancel)

       // Cancelando a execução da goroutine após 3 segundos
       time.Sleep(3 * time.Second)
       cancel()

       fmt.Println("Programa finalizado")
   }

   Neste exemplo, criamos um contexto de background e um contexto com cancelamento.
   Em seguida, executamos uma goroutine que imprime uma mensagem a cada segundo.
   Após 3 segundos, cancelamos a execução da goroutine usando o contexto com cancelamento.
*/

func main() {
	fmt.Printf("Hello!\n")

	ctx := context.Background() // Contexto em branco é aquele que só é definido para ficar em background
	ctxWithTimeOut, cancel := context.WithTimeout(ctx, time.Second*3)
	/*
	   Aqui transformamos o ctx tendo a possibilidade de ser cancelado quando alguem executar a função cancel
	*/
	defer fmt.Printf("Um print para ser executado após o cancel\n")
	defer cancel() // boa pratica sempre registrar um cancel.
	bookHotel(ctxWithTimeOut)

}

func bookHotel(ctx context.Context) {
	fmt.Printf("Antes do select\n")
	select {
	case <-ctx.Done():
		fmt.Printf("Hotel booking canceled. Intervalo de tempo para agendamentos acabou!\n")
		return
	case <-time.After(1 * time.Second): // trocar de 5 segundos para 2 segundo
		fmt.Printf("A vaga no hotel foi booked...\n")
	}
	fmt.Printf("Apos do select\n")
}
