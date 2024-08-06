# Packaging
Packaging em Golang refere-se à organização e estruturação de código em pacotes para facilitar a reutilização e manutenção do código. Os pacotes em Golang são unidades de código que podem ser importadas e utilizadas em outros arquivos ou projetos.  
  
O `go mod` é um recurso introduzido no Go 1.11 para gerenciar dependências de forma mais eficiente. Alguns dos principais comandos do go mod são:  
  
`go mod init`: Inicializa um novo arquivo go.mod no diretório atual, que é usado para gerenciar as dependências do projeto.  
`go mod tidy`: Garante que as dependências listadas no go.mod estejam corretas e atualizadas.  
`go mod vendor`: Copia as dependências do projeto para o diretório vendor, permitindo o versionamento das dependências.  
`go mod download`: Baixa as dependências do projeto para o cache local.  
`go mod verify`: Verifica a integridade das dependências do projeto.  
As consequências de utilizar o go mod incluem uma melhor gestão de dependências, facilitando a reprodução do ambiente de desenvolvimento em diferentes máquinas, garantindo a consistência das versões das dependências e evitando problemas de compatibilidade.  
  
## GO mod
O passo-a-passo de como criar o arquivo go.mod e o que deve ser preenchido nele, bem como as opções disponíveis para preenchimento.  
  
Passo-a-passo para criar o arquivo go.mod:  
* Inicialização do módulo: No diretório raiz do seu projeto, execute o comando go mod init nome-do-modulo. Isso criará o arquivo go.mod e inicializará o módulo com o nome especificado.  
* Adição de dependências: Ao importar pacotes externos no seu código, o go mod irá automaticamente adicionar essas dependências ao arquivo go.mod.  
* Gerenciamento de dependências: Utilize comandos como go mod tidy para garantir que as dependências listadas no go.mod estejam corretas e atualizadas.  
  
O que deve ser preenchido no arquivo go.mod:  
`Módulo`: O nome do módulo é especificado na primeira linha do arquivo go.mod. Ele define o escopo do módulo e é usado para identificar o pacote.  
`Dependências`: As dependências do projeto são listadas no arquivo go.mod, juntamente com suas versões. O go mod gerencia automaticamente as versões das dependências.  
  
Opções disponíveis de keywords para preenchimento no arquivo go.mod:  
`Module`: Define o nome do módulo e o caminho do repositório.  
Exemplo: module github.com/seu-usuario/seu-projeto  
  
`Require`: Especifica as dependências do projeto e suas versões.
Exemplo: require github.com/pacote-exemplo v1.2.3  
  
`Exclude`: Permite excluir uma versão específica de uma dependência.    
Exemplo: exclude github.com/pacote-exemplo v1.2.4  
  
`Replace`: Substitui uma dependência por outra, útil para testes ou desenvolvimento.  
Exemplo: replace github.com/pacote-exemplo => ../pacote-local  

Essas são algumas das opções principais que podem ser preenchidas no arquivo go.mod. O go mod oferece um controle mais preciso sobre as dependências do projeto e facilita a gestão das versões.

## Go Get

O comando `go get` é uma ferramenta do Go que é usada para baixar e instalar pacotes e suas dependências a partir de repositórios remotos. Aqui estão alguns pontos importantes sobre o `go get`:  
  
Instalação de Pacotes: O `go get` é comumente usado para baixar e instalar pacotes Go e suas dependências a partir de repositórios remotos, como o GitHub. Por exemplo, ao executar `go get` github.com/nome-do-pacote, o pacote será baixado e instalado no diretório de trabalho.  
  
Atualização de Pacotes: O `go get` também pode ser usado para atualizar pacotes já instalados. Ao executar `go get` -u nome-do-pacote, o pacote será atualizado para a versão mais recente disponível.  
  
Instalação de Ferramentas: Além de pacotes, o `go get` também pode ser usado para instalar ferramentas e utilitários relacionados ao Go. Por exemplo, `go get` golang.org/x/tools/cmd/godoc instala a ferramenta godoc.  
  
Depreciação: Com a introdução do módulo Go (go mod), o uso do `go get` para gerenciar dependências em projetos modernos é desencorajado. O go mod oferece um método mais robusto e controlado para gerenciar dependências.  
  
Variáveis de Ambiente: O `go get` respeita as variáveis de ambiente relacionadas ao Go, como GOPATH e GOROOT, para determinar onde os pacotes devem ser instalados.  
  
É importante notar que o uso do `go get` pode afetar a estabilidade e a consistência do ambiente de desenvolvimento, especialmente em projetos que utilizam o go mod para gerenciamento de dependências.  
  
## Go mod e GOENV | GOROOT | GOPATH

O go mod não faz modificações diretas nas variáveis de ambiente `GOROOT`, `GOPATH` e `GOENV`. No entanto, o uso do go mod pode impactar indiretamente essas variáveis devido à forma como o Go e suas ferramentas são configuradas.  
  
`GOROOT`: A variável de ambiente `GOROOT` aponta para o diretório de instalação do Go. O go mod não altera o `GOROOT`, pois é responsabilidade do usuário manter essa variável apontando para a instalação correta do Go.  
  
`GOPATH`: A variável de ambiente `GOPATH` é usada para especificar o diretório raiz dos seus projetos Go e suas dependências. Com a introdução do go mod, a necessidade de definir `GOPATH` diminuiu, pois o go mod gerencia as dependências de forma mais independente. No entanto, o `GOPATH` ainda é relevante para projetos que não utilizam o go mod ou para configurações específicas.  
  
`GOENV`: A variável de ambiente `GOENV` não é uma variável padrão do Go. Ela pode ser usada para definir configurações específicas do ambiente Go, como a versão do Go a ser utilizada. O go mod não interfere diretamente na variável `GOENV`, mas pode ser afetado por configurações globais do ambiente Go.  
  
Em resumo, o go mod não faz modificações diretas nas variáveis `GOROOT`, `GOPATH` e `GOENV`, mas é importante entender como essas variáveis se relacionam com o gerenciamento de dependências em projetos Go.