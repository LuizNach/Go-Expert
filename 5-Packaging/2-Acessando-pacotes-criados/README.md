# Acessando pacotes criados

## Importação de Pacotes no Go
No Go, a importação de pacotes é feita utilizando a palavra-chave import seguida do caminho do pacote que você deseja importar. Aqui está um exemplo básico:

```go
package main

import (
    "fmt"
    "math"
)

func main() {
    fmt.Println("Hello, World!")
    fmt.Println("The square root of 4 is:", math.Sqrt(4))
}
```

Neste exemplo, estamos importando os pacotes fmt e math. O pacote fmt é usado para formatação de strings e saída, enquanto o pacote math fornece funções matemáticas básicas.

## Importância do Endereço do Módulo como Link para um Repositório
Quando você está desenvolvendo um projeto em Go, é comum que você crie seus próprios pacotes e módulos. Para facilitar a reutilização e a colaboração, é uma boa prática hospedar esses pacotes em um repositório, como GitHub, GitLab ou Bitbucket. Aqui estão algumas razões para isso:  
  
* Facilidade de Acesso: Outros desenvolvedores podem facilmente acessar e usar seu código. Eles podem simplesmente importar seu pacote usando o caminho do repositório.  
  
* Controle de Versão: Hospedar seu código em um repositório permite que você utilize sistemas de controle de versão como Git. Isso facilita o gerenciamento de diferentes versões do seu código e a colaboração com outros desenvolvedores.  
  
* Distribuição: Quando seu código está em um repositório público, ele pode ser facilmente distribuído e utilizado por outros projetos. Isso é especialmente útil para bibliotecas e ferramentas que você deseja compartilhar com a comunidade.  
  
* Documentação e Exemplos: Repositórios geralmente permitem que você adicione documentação e exemplos de uso, o que pode ser extremamente útil para outros desenvolvedores que desejam utilizar seu código.  
  
Aqui está um exemplo de como você pode importar um pacote de um repositório GitHub:  

```go
package main

import (
    "fmt"
    "github.com/username/reponame"
)

func main() {
    reponame.SomeFunction()
    fmt.Println("Imported a custom package from GitHub!")
}
```

Neste exemplo, github.com/username/reponame é o caminho do pacote que está hospedado no GitHub.