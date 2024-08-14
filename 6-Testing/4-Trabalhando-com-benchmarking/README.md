# Trabalhando com benchmarking

Analisando os dados exibidos de um benchmark: Ao rodar o comando `go test -bench=.` vamos receber várias informações sobre o teste de benchmark.

```sh
goos: linux
goarch: amd64
pkg: github.com/LuizNach/Go-Expert/6-Testing/4-Trabalhando-com-benchmarking
cpu: 11th Gen Intel(R) Core(TM) i5-11600K @ 3.90GHz
BenchmarkCalculateTax-6         1000000000               0.2156 ns/op
PASS
ok      github.com/LuizNach/Go-Expert/6-Testing/4-Trabalhando-com-benchmarking  0.247s
```

Sistema operacional: linux
Arquitetura do Sistema: amd64
Pacote sendo testado: github.com/LuizNach/Go-Expert/6-Testing/4-Trabalhando-com-benchmarking
CPU do computador: cpu: 11th Gen Intel(R) Core(TM) i5-11600K @ 3.90GHz
Funcao executada para o benchmark: BenchmarkCalculateTax
Numero de processadores utilizados: 6
Quantidade de operações que conseguiu executar nessa função: 1000000000
Quanto tempo por operação: 0.2156 ns/op
