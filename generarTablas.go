package main

import (
	"fmt"
	"math"
	"time"
	"math/rand"
)

func main() {

	var n1 Neurona
	n1.crearTablas(2)
	fmt.Println("Tabla And: ",n1.tablaAnd)
	fmt.Println("Tabla Or: ",n1.tablaOr)
	fmt.Println("Pesos And: ",n1.pesosAnd)
	fmt.Println("Pesos Or: ",n1.pesosOr)
	fmt.Println(n1.calcularSalidaReal())
}

type Neurona struct {
	tablaAnd [][]int
	tablaOr  [][]int
	pesosAnd []float32 //Pesos para las entradas
	pesosOr []float32  //la ultima posición del arreglo pesos sera la del umbral
}

func calcularPotencia(valor float64) float64  {
	return math.Pow(2,valor)
}

func (n *Neurona) crearTablas(variables int)  {
	filas := int(calcularPotencia(float64(variables)))
	n.tablaAnd = make([][]int,variables+1)
	v := 0
	for x := 0; x < len(n.tablaAnd) ; x++  {
		n.tablaAnd[x] = make([]int,filas)
		for j := 0; j < filas ; j++  {
			n.tablaAnd[x][j] = v
			v++
		}
	}

	n.tablaOr = make([][]int, variables+1)

	copy(n.tablaOr,n.tablaAnd)

	n.generarPesos(variables)

	//modificar la or en la salida


}

func (n *Neurona) generarPesos(variable int)  {
	n.pesosAnd = make([]float32,variable+1)
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	for index, _ := range n.pesosAnd{
		n.pesosAnd[index] = r.Float32()
	}
	n.pesosOr = make([]float32,variable+1)
	copy(n.pesosOr,n.pesosAnd)

}

func (n *Neurona) calcularSalidaReal() (salida float32)  {

	for x := 0; x < len(n.tablaAnd); x++  {
		for y := 0; y <= len(n.pesosAnd) ;y++  {

			if x != len(n.tablaAnd) && y != len(n.tablaAnd) {
				fmt.Println(n.tablaAnd[x][y])
				salida += float32(n.tablaAnd[x][y]) * n.pesosAnd[y]
				//salida += n.pesosAnd[len(n.pesosAnd)]
			}

		}
	}
	return salida
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