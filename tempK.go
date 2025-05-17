// declara o pacote principal
package main

// importar a função fmt
import "fmt"

// declaracao da variavel CONST do ponto de ebulicao em F
const ebulicaok = 373.15

// função principal
func main() {

	tempK := ebulicaok
	tempC := (tempK - 273)

	// APARECER O RESULTADO
	fmt.Printf("A temperatura de ebulição da água em ºK = %g / A temperatura de ebulição da água em ºC = %g.", tempK, tempC)

}
