package main

import (
	"fmt"
)

func main() {
	fmt.Printf("Hello\n")

	var minha_var interface{} = "Luiz" // variável declarada com interface vazia.
	println("var minha_var interface{} = \"Luiz\" // variável declarada com interface vazia")

	fmt.Printf("Minha car do tipo %T valor %v\n", minha_var, minha_var)

	println("println interpreta uma variavel de tipo interface vazia somente como uma referencia de memoria")
	println(minha_var)
	println("Para solucionar isso e necessario um cast fazendo .(string)")
	println("minha_var.(string)", minha_var.(string))
	println("estou afirmando pro sistema que consigo converter para string, logo tomar cuidado com o tipo para que converteremos")

	println("--------------------------------")

	println("Agora podemos pensar em como garantimos que a conversao da certo sem interromper o programa com um panic")

	println("Como por exemplo: println(minha_var.(float32)) // panic: interface conversion: interface {} is string, not float32")

	println("Para solucionar isso temos o type assertion")
	println("Conversoes type assertion retornam dois resultados: convertido e um ok (true para sucesso|false para falha na conversao)")

	resultado, ok := minha_var.(float32)
	println("resultado, ok := minha_var.(float32)")

	if ok == false {
		println("Nao foi possivel fazer a conversao de", minha_var, "para float32")
	} else {
		println("Temos o valor convertido em float32 como", resultado)
	}

	println("--------------------------------")

	println("Só é possível fazer a conversão por type assertion com variáveis declardas como interfaces vazias")
	println("nova_var := 10")
	println("resultado, ok = nova_var.(string) // invalid operation: nova_var (variable of type int) is not an interface. compiler(InvalidAssert)")

	println("--------------------------------")

}
