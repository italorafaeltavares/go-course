package main

import "fmt"

const Name05 = "constante"

func main() {
	var name string
	name = "utilizadndo o var"

	var name02 string = "iniciando o valor com var"

	var name03 = "inferencia de tipo"

	name04 := "declaracao curta"

	fmt.Println(name)
	fmt.Println(name02)
	fmt.Println(name03)
	fmt.Println(name04)
	fmt.Println(Name05)

	var age int = 38
	var pi float32 = 3.14  //6-7 digitos de precisao
	var pi2 float64 = 3.14 //15-16 digitos de precisao

	fmt.Println(age)
	fmt.Println(pi)
	fmt.Println(pi2)

	var isEnabmed bool = true
	fmt.Println(isEnabmed)

	var a, b, c = "a", 50, true
	fmt.Println(a, b, c)
}
