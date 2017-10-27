package main

import (
	"fmt"
	"math"
)

func main() {

	var n1 Neurona
	n1.crearAnd(3)
	alternador(4,8)
}

type Neurona struct {
	tablaAnd [][]int
	tablaOr  [][]int
	pesos []float32
}

func calcularPotencia(valor float64) float64  {
	return math.Pow(2,valor)
}

func (t *Neurona) crearAnd(variables int)  {
	filas := int(calcularPotencia(float64(variables)))
	t.tablaAnd = make([][]int,variables+1)
	for x := 0; x < len(t.tablaAnd) ; x++  {
		t.tablaAnd[x] = make([]int,filas)
		for j := 0; j < filas ; j++  {
			t.tablaAnd[x][j] = 1
		}
	}
	fmt.Println(t.tablaAnd)


}

func (t *Neurona) crearOr(variables int)  {

}

/*Dada una logintud y una cantidad
Implementar la función que permita alternar los valores del arrreglo 0 y 1
Ejemplo:
cantidad: 2  y longitud 8

[0 0 1 1 0 0 1 1]

*/

func alternador(cantidad, longitud int) (array []int)  {
	return nil
}