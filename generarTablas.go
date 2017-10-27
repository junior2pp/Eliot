package main

import (
	"fmt"
	"math"
)

func main() {

	var t Tabla
	t.crearAnd(3)
	alternador(4,8)
}

type Tabla struct {
	and [][]int
	or  [][]int
}

type Pesos struct {
	pesos [][]float32
}

func calcularPotencia(valor float64) float64  {
	return math.Pow(2,valor)
}

func (t *Tabla) crearAnd(variables int)  {
	filas := int(calcularPotencia(float64(variables)))
	t.and = make([][]int,variables+1)
	for x := 0; x < len(t.and) ; x++  {
		t.and[x] = make([]int,filas)
		for j := 0; j < filas ; j++  {
			t.and[x][j] = 1
		}
	}
	fmt.Println(t.and)


}

func (t Tabla) crearOr(variables int)  {

}

/*Dada una logintud y una cantidad
Implementar la función que permita alternar los valores del arrreglo 0 y 1
Ejemplo:
cantidad: 2  y longitud 8

[0 0 1 1 0 0 1 1]

*/
func alternador(cantidad, longitud int) (array []int)  {

}