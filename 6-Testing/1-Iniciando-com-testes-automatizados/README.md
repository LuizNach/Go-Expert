# Testing

Para Executar testes que estejam dentro do repositório, devemos:  
* Construir os testes com o mesmo pacotes do pacotes que queremos testar  
* Podemos fazer o assert, avalização, verificação de um valor com o if padrão  
* Devemos nomear os arquivos de teste com _test.go no seu final  
* Devemos executar o modulo nativo do golang de teste executando: 
    * `go test .` para execução de qualquer arquivo de teste dentro do repositório local
    * `go test -v .` para execução em modo verbose, para mais detalhes da execução