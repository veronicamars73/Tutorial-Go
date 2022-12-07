package main

import "fmt"

func main() {
	var choice int = -1
	for choice != 0 {
		menu()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			fmt.Println("1")

		case 2:
			fmt.Println("2")

		case 3:
			fmt.Println("3")

		case 4:
			fmt.Println("4")

		case 5:
			fmt.Println("5")

		case 6:
			fmt.Println("6")

		case 7:
			fmt.Println("7")

		case 8:
			fmt.Println("8")

		case 0:
			fmt.Println("0")

		default:
			fmt.Println("Opção inválida")
		}
	}
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
