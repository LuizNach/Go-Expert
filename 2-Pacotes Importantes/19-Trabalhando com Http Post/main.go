package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

/*
   O método Post da biblioteca http da linguagem Go é utilizado para enviar uma requisição HTTP do tipo POST para um servidor.
   Ele é comumente utilizado para enviar dados para o servidor, como por exemplo, ao submeter um formulário em uma aplicação web.
   O funcionamento detalhado do método Post envolve a criação de uma requisição do tipo POST, onde os dados a serem enviados são
   incluídos no corpo da requisição. Esses dados podem ser enviados em diferentes formatos, como JSON, formulários HTML, entre outros.
   Após a criação da requisição, ela é enviada para o servidor especificado, que irá processar os dados recebidos e retornar uma
   resposta ao cliente. É importante ressaltar que o método Post é uma operação síncrona, ou seja, o cliente aguarda a resposta do
   servidor antes de continuar a execução do código.
   Caso haja necessidade de mais detalhes específicos sobre o funcionamento do método Post da biblioteca http da linguagem Go,
   recomendo consultar a documentação oficial da linguagem ou buscar por exemplos práticos de uso deste método.


   O método http.Client.Post da linguagem Go é utilizado para enviar uma requisição HTTP do tipo POST para um servidor.
   Ele possui a seguinte assinatura:

       func (c *Client) Post(url string, bodyType string, body io.Reader) (resp *Response, err error)


   Aqui está a explicação detalhada dos parâmetros utilizados no método http.Client.Post:

       url string: Este parâmetro representa a URL para a qual a requisição POST será enviada.
           É uma string que deve conter o endereço completo, incluindo o protocolo (http:// ou https://) e, se necessário, a porta.

       bodyType string: Este parâmetro especifica o tipo de conteúdo que está sendo enviado no corpo da requisição.
           Por exemplo, se estiver enviando dados no formato JSON, você pode especificar "application/json".
           Se estiver enviando um formulário HTML, pode ser "application/x-www-form-urlencoded".

       body io.Reader: Este parâmetro representa o corpo da requisição, ou seja, os dados que serão enviados para o servidor.
           Ele deve ser do tipo io.Reader, o que significa que pode ser qualquer tipo que implemente a interface Reader, como
           por exemplo um bytes.Buffer ou um strings.Reader.

   Após chamar o método Post com os parâmetros adequados, ele irá criar uma requisição do tipo POST com a URL, tipo de conteúdo
   e corpo especificados, e enviar essa requisição para o servidor. O método retornará um objeto Response contendo a resposta do
   servidor e um possível erro, caso ocorra algum problema durante a requisição.

   É importante ressaltar que o método http.Client.Post é uma forma conveniente de enviar requisições POST, porém, para cenários
   mais complexos, pode ser necessário utilizar outras funcionalidades da biblioteca http padrão da linguagem Go.
*/

func main() {
	fmt.Printf("Hello\n")

	// Instanciamos um http client
	c := http.Client{Timeout: time.Duration(2) * time.Second}

	varSliceDeBytes := []byte("{This is the body for the post request}")

	bufferJson := bytes.NewBuffer(varSliceDeBytes) // Criamos um buffer a partir do slice de bytes que temos.

	/*
	   Buffer em Go é uma estrutura de dados flexível e eficiente que permite armazenar e manipular dados em memória de forma dinâmica,
	   facilitando diversas operações de manipulação de dados em programas Go
	*/

	// Detalhe aqui que o post recebe a string da url, string do tipo de dados que vamos transmitir e um buffer que contem o que queremos transmitir
	resp, err := c.Post("http://google.com", "application/json", bufferJson)

	if err != nil {
		panic("Deu erro na requisicao http")
	}

	defer resp.Body.Close()

	// body, err := io.ReadAll(resp.Body)
	// if err != nil {
	// 	panic("Nao foi possivel resuperar o body da requisicao http")
	// }
	// fmt.Printf("O body da response: %s\n", body)

	// Um exemplo de utilização de buffer: podemos copiar a resposta do post, que fica no resp.Body
	fmt.Printf("O que vem do copy buffer: \n")
	io.CopyBuffer(os.Stdout, resp.Body, nil)

	/*
		Um detalhe importante aqui é que se fizermos io.ReadAll no resp.Body estaremos limpando todo o buffer do
		resp.Body e passando tudo para a variavel body. Logo se fizermos o io.CopyBuffer enviando tudo do resp.Body
		para o io.Stdout não haverá mais nada sobrando pois tudo já foi enviado para variável body.
		Para testar basta descomentar a lionha 75~79
	*/

}
