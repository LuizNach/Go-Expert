# Testing Coverage

Para descobrir a cobertura de testes podemos executar os testes com a flag de coverage rodando: `go test -v -coverprofile=coverage.out . `  
Isso exibirá no CLI o % de code coverage que conseguimos testar. Facilitando a analise da métrica.
Podmeos agora tentar descobrir quais partes do código não foram percorridas dentro dos nossos casos de teste.  
Excutamos o `go tool cover -html=coverage.out -o coverage.html` e assim temos o relatório para abrirmos no browser destacando quais locais ainda não percorremos com os testes. Segue a screenshot:

![coverage](./Screenshot-from-2024-08-13%2003-32-44.png)