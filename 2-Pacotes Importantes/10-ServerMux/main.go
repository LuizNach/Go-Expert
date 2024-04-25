package main

import (
	"fmt"
	"net/http"
)

/*
   Um http.ServeMux em Go é um multiplexador de solicitações HTTP, que é usado para rotear solicitações HTTP
   para funções handler específicas com base nos padrões de URL. Ele atua como um roteador que direciona as
   solicitações recebidas para os manipuladores apropriados com base nos caminhos das URLs. O http.ServeMux
   é uma parte fundamental da criação de aplicativos web em Go, permitindo o roteamento eficiente das
   solicitações HTTP para as funções handler corretas.

   Basicamente é um componente que attachamos as rotas criadas nele.
*/

/*
   Exemplo de main sem utilizar o server mux:

   func handler(w http.ResponseWriter, r *http.Request) {
       fmt.Fprintf(w, "Bem-vindo ao servidor sem ServeMux!")
   }

   func main() {
       http.HandleFunc("/", handler)

       http.ListenAndServe(":8080", nil)
   }


   Vantagens e desvantagens:
       Com http.ServeMux:
           Vantagens:

               Permite registrar múltiplos manipuladores para diferentes rotas.
               Facilita a organização e manutenção do código, especialmente em aplicações maiores.
               Flexibilidade para adicionar middleware e funcionalidades extras.

           Desvantagens:

               Pode adicionar complexidade desnecessária em aplicações simples.
               Requer um entendimento adicional sobre como funciona o roteamento em Go.

       Sem http.ServeMux:
           Vantagens:

               Simplicidade e facilidade de implementação em aplicações pequenas e simples.
               Menos código para lidar com rotas básicas.

           Desvantagens:

               Dificuldade em lidar com múltiplas rotas e funcionalidades mais avançadas.
               Menos flexibilidade para adicionar middleware e funcionalidades extras.
*/

// func handler(w http.ResponseWriter, r *http.Request) {
//     fmt.Fprintf(w, "Bem-vindo ao servidor sem ServeMux!")
// }
// func main() {
//     http.HandleFunc("/", handler)
//     http.ListenAndServe(":8080", nil)
// }

// func main() {
// 	mux := http.NewServeMux()
// 	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
// 		fmt.Fprintf(w, "Bem-vindo ao servidor com ServeMux!")
// 	})
// 	http.ListenAndServe(":8080", mux)
// }

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Bem-vindo ao servidor com ServeMux!")
	})
	mux.Handle("/blog", Blog{title: "Luizera"})
	http.ListenAndServe(":8080", mux)

	mux2 := http.NewServeMux()
	mux2.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Olha o mux 2 rolando!")
	})
	http.ListenAndServe(":8081", mux2)
}

// Um terceiro jeito de se usar o mux é passando não só funcoes mas utilizando o .Handle
// podemos passar qualquer estrutura desde que ela implemente ServeHTTP(w http.ResponseWriter, r *http.Request)
type Blog struct {
	title string
}

func (b Blog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Blogzera:" + b.title))
}
