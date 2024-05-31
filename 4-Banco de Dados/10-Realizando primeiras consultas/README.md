# GORM

Para levantar o container com o banco para se conectar:
```
docker compose up
```

Pela CLI para entrar no container basta fazer um exec bash:
```
docker container exec -it bash mysql
```
Uma vez dentro do container, podemos entrar no banco `goexpert` do mysql com:
```
mysql -uroot -ppassword goexpert
```
Ao entrar dentro do sgdb do banco `goexpert` podemos executar queries como:
```
show tables;
desc products;
select * from products;
drop table products;
```