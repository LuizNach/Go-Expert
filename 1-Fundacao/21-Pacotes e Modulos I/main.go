package main

//"fmt"
//"matematica" // could not import matematica (cannot find package "matematica" in GOROOT or GOPATH)compilerBrokenImport

//"./matematica" // main.go:6:2: "./matematica" is relative, but relative import paths are not supported in module mode

// A principio a funcao Soma do pacote matematica nao e encontrada pois estao em pacotes diferentes
// Para que possamos utilizar a funcao devemos importar o pacote ou somente a funcao do pacote
// import puro "matematica" nao pode ser achado por que ele esta buscando no GOROOT ou no GOPATH
// Para resolver isso, resolvemos atraves do go modules

func main() {
	//fmt.Printf("O resultado de  2.5 + 3.5 é: %v\n", matematica.Soma(2.5, 3.5))
}
