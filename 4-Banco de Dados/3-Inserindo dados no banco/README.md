# Executando
Para criar um serviço mysql disponível na porta 3306 pelo docker, basta rodar:
```
docker compose up -d
```
Para acessar o container com nome mysql, basta rodar:
```
docker container exec -it mysql bash
```
Uma vez dentro do terminal do container podemos acessaro sgdb do mysql com o comando:
```
mysql -u <user> -p <password> <table-name>
```
No caso da configuração do docker compose será:
```
mysql -u root -p password goexpert
```
Uma vez dentro do mysql podemos requisitar uma query SQL para listar todos os produtos cadastrados para que saibamos o que o programa está funcionando:
```
select * from products;
```
Para sair do sgdb MySql é só usar o comando `exit` da mesma maneira para sair do container com o comando `exit`.