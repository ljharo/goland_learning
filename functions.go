package main

import (
	"fmt"
)

func esPar(){

	fmt.Println("el programa valida si un valor es par o par")
	var valor int
	fmt.Print("ingresa el valor: ")
	fmt.Scanln(&valor)

	if (valor % 2) == 0 {
		fmt.Println("es par")
	} else {
		fmt.Println("es impar")
	}
}

func ejemploSwith (){
	
	var i int8 = 0

	for i < 1 {
		var diaSemana int8
		fmt.Printf("Ingresa el dia de la semana: ")
		fmt.Scanln(&diaSemana)

		switch diaSemana {
			case 1:
				fmt.Println("Lunes")
				i += 2
			case 2:
				fmt.Println("Martes")
				i += 2
			case 3:
				fmt.Println("Miercoles")
				i += 2
			case 4:
				fmt.Println("Jueves")
				i += 2
			case 5:
				fmt.Println("Viernes")
				i += 2
			case 6:
				fmt.Println("Sabado")
				i += 2
			case 7:
				fmt.Println("Domingo")
				i += 2
			default:
				fmt.Println("Dia no valido, ingresa otro.")
		}
	}
}

func mayorQue(value int8) {
	switch {
		case value > 10:
			fmt.Printf("mayor %d que 10\n", value)
			fallthrough // se usa para que se ejecute el siguiente case
		case value > 20:
			fmt.Printf("mayor %d que 20\n", value)
			fallthrough // se usa para que se ejecute el siguiente case
		case value > 30:
			fmt.Printf("mayor %d que 30\n", value)
	}
}