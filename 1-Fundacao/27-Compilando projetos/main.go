package main

/*
Queremos compilar esse codigo para varias plataformas diferentes.
A execuçpao do go possui na maquina um programa Run time que executa seu codigo go toda vez que rodamos o go run main.go.
Compilarmos para uma plataforma especifica é o mesmo que juntar todo codigo do run time e adicionar o código do projeto e
transformar para o binário da plataforma especifica.

Para verificar qual plataforma vc pode compilar paa rodar na sua maquina ie só verificar em qual GOOS seu run time esta rodando.
Para fazer isso basta rodar go env GOOS.
Para descobrir a arquitetura do seu runtime basta rodar go env GOARCH.

Para listar todos os possiveis combinações de sistemas operacionais e arquiteturas possíveis de serem compiladas pelo run time: go tool dist list
Reference: https://www.digitalocean.com/community/tutorials/building-go-applications-for-different-operating-systems-and-architectures

Para por exemplo, compilarmos um build para um executavel para windows arquitetura amd64
GOOS=windows GOARCH=amd64 go build main.go
GOOS=windows GOARCH=amr go build main.go
Isso gerara um .exe pronto para windows mesmo que estejamos programando em um linux.

*/

func main() {

}
