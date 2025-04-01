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

func main()  {

	//// strings
	// var str string

	//// booleans
	// var b bool

	//// ints 
	// // numeros enteros
	// var entero8 int8
	// var entero16 int16
	// var entero32 int32
	// var entero64 int64

	// // numeros enteros sin signos
	// var uentero8 uint8
	// var uentero16 uint16
	// var uentero32 uint32
	// var uentero64 uint64

	// // alias 
	// var byte byte // alias de uint8
	// var rune rune // alias de int32

	// // tipos dependientes de la plataforma
	// var enteroUint uint // 32 o 64
	// var enteroInt int   // 32 o 64
	// var enteroUintptr uintptr // sirve para almacenar el valor de un puntero

	// //// floats
	// // valor decimal
	// var flotante32 float32
	// var flotante64 float64

	//// arrays
	// var arrayint [5]int
	// var arrayFloat [5]float
	// var arrayBolean [5]boolean
	// var arraystring [5]string

	// solo existe el ciclo for

	// var show bool = false
	// if show {
	// 	var i int = 0
	// 	for i < 5 {
	// 		fmt.Println("Hola")
	// 		i++
	// 	}

	// 	fmt.Println("------------------------------")

	// 	for i := 0; i < 5; i++ {
	// 		fmt.Println(i)
	// 	}

	// 	fmt.Println("------------------------------")

	// 	i = 0
	// 	for {
	// 		fmt.Println(i)
	// 		i++
	// 		if i > 5 { break }
	// 	}

	// }
	
	// var a int = 10
	// var b int = 10

	// if a > b {
	// 	fmt.Println("a es mayor que b")
	// } else if a == b {
	// 	fmt.Println("a es igual a b")
	// } else {
	// 	fmt.Println("a es menor que b")
	// }

	esPar()

}

// link vide
// https://www.youtube.com/watch?v=P9oAMtf_t30&list=PLl_hIu4u7P64MEJpR3eVwQ1l_FtJq4a5g&index=21