package main

import (
	"fmt"
	"time"
	"math/rand"
)

func main()  {
	                   // A  B  C  S
	tablaAnd := [][]int{ {1, 1, 1, 1} ,
					     {1, 1, 0, 0} ,
					     {1, 0, 1, 0} ,
					     {1, 0, 0, 0} ,
					     {0, 1, 1, 0} ,
					     {0, 1, 0, 0} ,
					     {0, 0, 1, 0} ,
					     {0, 0, 0, 0} ,
					     }


	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	pesos := [4]float32{r.Float32(), r.Float32(), r.Float32(), r.Float32()}


	aprendiendo := true
	iteraciones := 0
	var tasaAprende float32 = 0.3


	for aprendiendo {
		iteraciones++
		aprendiendo = false
		fmt.Println("Peso Entrada A: ", pesos[0])
		fmt.Println("Peso Entrada B: ", pesos[1])
		fmt.Println("Peso Entrada C: ", pesos[2])
		fmt.Println("Peso Umbral: ", pesos[3])
		muestraTabla(tablaAnd, &pesos)

		for y := 0; y < 8 ; y++ {
			salidaReal := (float32(tablaAnd[y][0]) * pesos[0]) + ( float32(tablaAnd[y][1]) *pesos[1] ) + (float32(tablaAnd[y][2])*pesos[2]) + pesos[3]
			salidaEntera := salida(salidaReal) //retorna 1 si es mayor a 0, de lo contrario 0
			fmt.Print("Salida Real: ",salidaReal)
			fmt.Println("  Salida Entera: ",salidaEntera)
			error := calcularError(tablaAnd,salidaEntera,y)

			if  error != 0 {
				cambiarPesos(&pesos,tasaAprende,tablaAnd,y,error)
				aprendiendo = true
			}


		}
		fmt.Println("------------------------------------------------------------------")

	}
	fmt.Println("Cantidad de Iteraciones: ", iteraciones)
	fmt.Println("Peso 1: ", pesos[0])
	fmt.Println("Peso 2: ", pesos[1])
	fmt.Println("Peso 3: ", pesos[2])
	fmt.Println("Peso Umbral: ", pesos[3])

	fmt.Println("Aprendida satisfactoriamente")
	muestraTabla(tablaAnd, &pesos)
}


func salida(valor float32)  int  {
	if valor > 0 {
		return 1
	} else {
		return 0
	}
}

func muestraTabla(tablaAnd [][]int, pesos *[4]float32)  {

	//porcentaje := [4]int{0,0,0,0}
	for x := 0; x < 8 ; x++  {
		salidaReal := float32(tablaAnd[x][0]) * pesos[0] + float32(tablaAnd[x][1]) * pesos[1] + float32(tablaAnd[x][2]) * pesos[2] + pesos[3]
		salidaEntera := salida(salidaReal)
		//if salidaEntera == tablaAnd[x][2] {
		//	porcentaje[x] = 25
		//}
		fmt.Println("Entradas: ", tablaAnd[x][0], " y " , tablaAnd[x][1], " y ", tablaAnd[x][2] ," = ", tablaAnd[x][3] , "perceptron = ",salidaEntera)
	}
	//porcetanjeV := porcentaje[0]+porcentaje[1]+porcentaje[2]+porcentaje[3]
	//fmt.Println("Porcentaje de Aprendizaje: ",porcetanjeV , "%")
	fmt.Println("------------------------------------------------------------------")

}
//formula Frank Rosenblatt
func cambiarPesos(pesos *[4]float32, tasaAprende float32, tabla [][]int, posicion,error int)  {
	error1 := float32(error)
	pesos[0] += tasaAprende*error1*float32(tabla[posicion][0])
	pesos[1] += tasaAprende*error1*float32(tabla[posicion][1])
	pesos[2] += tasaAprende*error1*float32(tabla[posicion][2])
	pesos[3] += tasaAprende*error1*1
}

func calcularError(tabla [][]int, salidaEntera,posicion int) (error int) {
	error = tabla[posicion][3] - salidaEntera
	return error
}