package main

import "fmt"

func main() {
	var choice int
	menu()
	fmt.Scan(&choice)
}

func menu() {
	fmt.Println("Calculadora de Geometria Plana e Espacial")
	fmt.Println("(1) Triângulo equilátero")
	fmt.Println("(2) Retângulo")
	fmt.Println("(3) Quadrado")
	fmt.Println("(4) Cîrculo")
	fmt.Println("(5) Pirâmide com base quadrangular")
	fmt.Println("(6) Cubo")
	fmt.Println("(7) Paralelepîpedo")
	fmt.Println("(8) Esfera")
	fmt.Println("(0) Sair")
	fmt.Println("")
	fmt.Println("Digite a sua opção:")
}
