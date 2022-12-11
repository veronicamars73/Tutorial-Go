package main

import (
	"fmt"
	"math"
)

func main() {
	var choice int = -1
	for choice != 0 {
		menu()
		fmt.Scan(&choice)
		switch choice {
		case 1:
			triangulo()

		case 2:
			retangulo()

		case 3:
			quadrado()

		case 4:
			circulo()

		case 5:
			piramide()

		case 6:
			cubo()

		case 7:
			paralelepipedo()

		case 8:
			esfera()

		case 0:
			fmt.Println("Adeus!")
			break

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

func triangulo() {
	var altura, base float32
	fmt.Println("Digite a base do triângulo: ")
	fmt.Scan(&base)
	fmt.Println("Digite a altura do triângulo: ")
	fmt.Scan(&altura)
	var area float32 = (base * altura) / 2
	var perimetro float32 = base * 3
	fmt.Printf("Área: %.2f\n", area)
	fmt.Printf("Perímetro: %.2f\n", perimetro)
}

func retangulo() {
	var altura, base float32
	fmt.Println("Digite a base do retângulo: ")
	fmt.Scan(&base)
	fmt.Println("Digite a altura do triângulo: ")
	fmt.Scan(&altura)
	var area float32 = (base * altura)
	var perimetro float32 = (base + altura) * 2
	fmt.Printf("Área: %.2f\n", area)
	fmt.Printf("Perímetro: %.2f\n", perimetro)
}

func quadrado() {
	var lado float64
	fmt.Println("Digite o lado do quadrado: ")
	fmt.Scan(&lado)
	var area float64 = math.Pow(lado, 2)
	var perimetro float64 = lado * 4
	fmt.Printf("Área: %.2f\n", area)
	fmt.Printf("Perímetro: %.2f\n", perimetro)
}

func circulo() {
	var raio float64
	fmt.Println("Digite o raio do círculo: ")
	fmt.Scan(&raio)
	var area float64 = math.Pow(raio, 2) * math.Pi
	var perimetro float64 = raio * 2 * math.Pi
	fmt.Printf("Área: %.2f\n", area)
	fmt.Printf("Perímetro: %.2f\n", perimetro)
}

func piramide() {
	var aresta_base, altura_lateral, altura_pir float64
	fmt.Println("Digite o tamanho da aresta do quadrado da base: ")
	fmt.Scan(&aresta_base)
	fmt.Println("Digite a altura do triângulo da lateral: ")
	fmt.Scan(&altura_lateral)
	fmt.Println("Digite a altura da pirâmide: ")
	fmt.Scan(&altura_pir)
	var area float64 = math.Pow(aresta_base, 2) + ((aresta_base * altura_lateral) / 2)
	var volume float64 = ((math.Pow(aresta_base, 2)) * altura_pir) / 3
	fmt.Printf("Área: %.2f\n", area)
	fmt.Printf("Volume: %.2f\n", volume)
}

func cubo() {
	var lado float64
	fmt.Println("Digite o tamanho da aresta do cupo: ")
	fmt.Scan(&lado)
	var area float64 = math.Pow(lado, 2) * 6
	var volume float64 = math.Pow(lado, 3)
	fmt.Printf("Área: %.2f\n", area)
	fmt.Printf("Volume: %.2f\n", volume)
}

func paralelepipedo() {
	var aresta_1, aresta_2, aresta_3 float32
	fmt.Println("Digite o tamanho da primeira aresta: ")
	fmt.Scan(&aresta_1)
	fmt.Println("Digite o tamanho da segunda aresta: ")
	fmt.Scan(&aresta_2)
	fmt.Println("Digite o tamanho da terceira aresta: ")
	fmt.Scan(&aresta_3)
	var area float32 = (2 * aresta_1 * aresta_2) + (2 * aresta_1 * aresta_3) + (2 * aresta_2 * aresta_3)
	var volume float32 = aresta_1 * aresta_2 * aresta_3
	fmt.Printf("Área: %.2f\n", area)
	fmt.Printf("Volume: %.2f\n", volume)
}

func esfera() {
	var raio float64
	fmt.Println("Digite o raio da esfera: ")
	fmt.Scan(&raio)
	var area float64 = math.Pow(raio, 2) * math.Pi * 4
	var volume float64 = math.Pow(raio, 3) * math.Pi * (4 / 3)
	fmt.Printf("Área: %.2f\n", area)
	fmt.Printf("Volume: %.2f\n", volume)
}
