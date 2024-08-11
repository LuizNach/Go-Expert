# Replace

Para o caso que o nosso projeto possui outros módulos dentro, podemos importá-los com o mesmo path do mesmo módulo.  
Para uma nova situação, se nosso projeto `sistema` utiliza um outro programa independente `math` de maneira que `sistema` possui o próprio go mod e `math` possui o próprio go mod.  
Nessa situação se `math` estiver publicado no github na versão correta que `sistema` está utilizando, será instalado quando rodarmos `go mod tidy`. Porém se a versão que `sitema` estiver utilizando ainda não estiver publicada não será possível baixá-la apesar de estarmos utilizando a versão localmente.  
Logo para que seja possível utilizar a versão local para desenvolvimento ou build podemos utilizar a flag `replace` do go mod da seguinte maneira:  
```
go mod edit -replace github.com/LuizNach/Go-Expert/5-Packaging/5-Trabalhando-com-go-mod-replace/math=../math
go mod tidy
```
Essa conexão fará que utilizamos a pasta local ao invés da versão publicada online.