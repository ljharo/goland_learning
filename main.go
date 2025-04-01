package main

import "fmt"

// "fmt"



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

	// esPar()
	// ejemploSwith()
	// mayorQue(50)

	// array normal
	var x [3]int8
	x[0] = 100
	fmt.Println(x)

	// matriz 
	var m[2][3]float32
	m[1][1] = 10.333333
	fmt.Println(m)

	var str []byte
	str = append(str, 'h')
	fmt.Printf("str: %v\n", str)

}

// link vide
// https://www.youtube.com/watch?v=yHlUlBuba7Q&list=PLl_hIu4u7P64MEJpR3eVwQ1l_FtJq4a5g&index=22