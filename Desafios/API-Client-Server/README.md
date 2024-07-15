# Descrição

Neste desafio vamos aplicar o que aprendemos sobre webserver http, contextos,
banco de dados e manipulação de arquivos com Go.
 
Você precisará nos entregar dois sistemas em Go:
- client.go
- server.go
 
Os requisitos para cumprir este desafio são:
 
O client.go deverá realizar uma requisição HTTP no server.go solicitando a cotação do dólar.
 
O server.go deverá consumir a API contendo o câmbio de Dólar e Real no endereço: https://economia.awesomeapi.com.br/json/last/USD-BRL e em seguida deverá retornar no formato JSON o resultado para o cliente.
 
Usando o package "context", o server.go deverá registrar no banco de dados SQLite cada cotação recebida, sendo que o timeout máximo para chamar a API de cotação do dólar deverá ser de 200ms e o timeout máximo para conseguir persistir os dados no banco deverá ser de 10ms.
 
O client.go precisará receber do server.go apenas o valor atual do câmbio (campo "bid" do JSON). Utilizando o package "context", o client.go terá um timeout máximo de 300ms para receber o resultado do server.go.
 
Os 3 contextos deverão retornar erro nos logs caso o tempo de execução seja insuficiente.
 
O client.go terá que salvar a cotação atual em um arquivo "cotacao.txt" no formato: Dólar: {valor}
 
O endpoint necessário gerado pelo server.go para este desafio será: /cotacao e a porta a ser utilizada pelo servidor HTTP será a 8080.
 
Ao finalizar, envie o link do repositório para correção.

# Como executar?
Para facilitar sem necessidade de alterações no código utilizaremos [docker-compose](https://docs.docker.com/compose/)

```bash
cd Server
docker compose up -d
cd ../Client
go run client.go
```
Validação no caso de execução sem problemas teremos o arquivo `cotacao.txt` na pasta Client.
Também é possível verificar o banco de dados MySql no docker container mysql executando os seguintes comandos em um novo terminal:
```bash
docker container ls # Para visualizar que os containers servidor e mysql estão rodando
docker container exec -it bash mysql # Fazer o attach em um processo iterativo bash no container nomeado mysql
mysql -u root -ppassword desafio-api-client-server # Acessar o banco desafio-api-client-server com o mysql
show tables; # Dentro do sgbd podemos listas todas as tabelas para verificarmos que a tabela cotação foi criada corretamente
select * from cotacoes; # Para visualizar todos os dados criados na tabela cotacoes
exit # Para sair do mysql
exit # Para sair do container mysql 
```
