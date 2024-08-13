
# Workspaces  
  
Podemos eliminar problemas de conexão dos pacotes criando um workspaces que ligará os pacotes faltantes com os que temos localmente:  
```sh
go work init ./math ./sistema
```  
O go conseguirá rodar os arquivos .go em ambos os pacotes caso um dependa do outro.  
  
## Go workspaces e go mod tidy  
Digamos que `sistema/main.go` depende de `math` mas também depende de `github.com/google/uuid`. Logo somente o workspaces não soluciona todos os problemas dado que precisamos instalar o pacotes `uuid` para que `sistema/main.go` funcione. Logo poderiamos fazer a instalação por meio do `go get`. Porém para o caso de muitas instalações podemoriamos utilizar a comodidade do `go mod tidy` pois já limpa dependencias não utilizadas.  
  
Para utilizar o workspaces e ao mesmo para utilizar o `go mod tidy`, podemos utilizar: `go mod tidy -e` nos repos que tem a dependencia externa. Assim vamos ignorar pacotes não publicados e poderemos utilizar o  o workspaces para modulos que ainda não foram publicados mas temos acesso local.  
  
Ou seja:  
```bash
go work init ./math ./sistema
cd sistema
go mod tidy -e
cd ..
go run sisetma/main.go
```

Reference: https://aprendagolang.com.br/como-utilizar-go-workspaces/
