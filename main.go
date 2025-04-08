package main

import (
	"fmt"
	"strings"
)

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

	// // array normal
	// var x [3]int8
	// x[0] = 100
	// fmt.Println(x)

	// // matriz 
	// var m[2][3]float32
	// m[1][1] = 10.333333
	// fmt.Println(m)

	// // declarar he inicializar un array 1
	// y := [5]int{1,2,3,5,6}
	// fmt.Println(y)

	// // declarar he inicializar un array 2
	// j := [...]int{1,2,3,4,5}
	// fmt.Println(j)

	// // declarar he inicializar un array 3
	// z := [...]int{
	// 	1,
	// 	2,
	// 	3,
	// 	4,
	// 	5,
	// }
	// fmt.Println(z)

	// var str []byte
	// str = append(str, 'h')
	// fmt.Printf("str: %v\n", str)


	// // comparar array

	// a1 := [3]int{1,2,3}
	// b1 := [...]int{1,2,3}
	// c1 := [3]int{1,3,3}
	// // d1 := [4]int{1,3,3}

	// fmt.Println(a1 == b1)
	// fmt.Println(a1 == c1)
	// // fmt.Println(a1 != d1) // no se puede porque el len de ambos es distinto

	// // podemosalterar el valor en dad pos
	// a1[2] = 5
	// fmt.Println(a1)


	// // slice, para crear un slice se utiliza el tipo de dato []T sin definir su length
	// var s []int
	// fmt.Println(s)

	// // declarar he inicializar un array 2
	// j1 := []int{1,2,3,4,5}
	// fmt.Println(j1)

	// // declarando un slice usando make, definiendo su length
	// s1 := make([]int, 5)
	// fmt.Println(s1)
	// fmt.Println(len(s1))
	// fmt.Println(cap(s1)) // capacidad del slice

	// // declarando un slice usando make, definiendo su length y capacidad
	// s2 := make([]int, 5, 10)
	// fmt.Println(s2)
	// fmt.Println(len(s2))
	// fmt.Println(cap(s2))

	// // definimos un array de 10 elementos
	// ar := [10]int{1,2,3,4,5,6,7,8,9}

	// // definimos dos slices

	// var a2, b2 []int
	// a2 = ar[3:5]
	// b2 = ar[2:5]

	// b2[0] = 25
	// fmt.Println(b2)
	// fmt.Println(a2)
	// fmt.Println(ar)
	// como se puede ver estamos trabajando con una copia de la subarreglo de ar, por lo que si cambiam

	// fmt.Println(cap(a2))
	// fmt.Println(cap(b2))
	// la capaciada del slice abarca desde donde empieza hasta donde termina

	// // append en slices
	// v := make([]byte, 4, 10)
	// fmt.Println(v)

	// // inicializamos x
	// v = []byte{'H', 'O', 'L', 'A'}
	// fmt.Println(v)
	// //  imprimimos el valor con el formato utf-8
	// fmt.Printf("%q\n", v)

	// // imprimimos cada valor
	// for i := 0; i < len(v); i++ {
	// 	fmt.Printf("v[%d] = %c\n", i, v[i])
	// }

	// // v[5] = ' ' // error
	// // fmt.Println(v)

	// v = append(v, ' ')
	// fmt.Println(v)

	// // agregamos MUNDO
	// v = append(v, 'M', 'U', 'N', 'D', 'O',)
	// fmt.Println(v)
	// fmt.Println(len(v))
	// fmt.Println(cap(v))

	// fmt.Printf("%q\n", v)

	// //realizar una copia de un slice

	// origen :=  []int{1,2,3}
	// destino := []int{3,4,5}

	// copy(origen, destino)
	// destino[2] = 22

	// fmt.Println(origin, destino)

	// // cuando el origen es mayor que el sestino
	// origin1 := []int{1, 2, 3}
	// destino1 := make([]int, 2)

	// copy(destino1, origin1)
	// fmt.Println(origin1, destino1)

	// // cuando el destino es mayor que el origen

	// origin2 := []int{1,2}
	// destino2 := []int{3,4,5}

	// copy(destino2, origin2)
	// fmt.Println(origin2, destino2)


	// for con range

	// var nombres []string = []string{
	// 	"juan",
	// 	"pedro",
	// 	"maria",
	// }

	// for i, nombre := range nombres {
	// 	fmt.Printf("Indice: %d, Nombre: %s\n", i, nombre)
	// 	for _, char := range nombre {
	// 		fmt.Printf("Indice: %d, Nombre: %s, Caracter: %c\n", i, nombre, char)
	// 	}
	// }

	// // tambien podemos recorrer strings
	// var texto string = "hola mundo desde goland"
	// for _, char := range texto {
	// 	fmt.Printf("Caracter: %c\n", char)
	// }

	// // declarando map

	// x := make(map[string]string)

	// x["nombre"] = "luis"
	// x["edad"] = "25"
	// x["ciudad"] = "bogota"
	// fmt.Println(x)

	// // acceder a los valores
	// fmt.Println(x["nombre"])
	// fmt.Println(x["edad"])

	// // declarando un map 3

	// var edades map[string]int = map[string]int{
	// 	"luis": 25,
	// 	"juan": 30,
	// 	"maria": 20,
	// }
	// fmt.Println(edades)

	// // borrar un registro del map'
	// delete(edades, "juan")
	// fmt.Println(edades)


	// // cuando buscamos un valor que no existe
	// fmt.Println(edades["juan"]) // cuando no existe  retorna 0

	// // saber si existe un valor
	// edad, ok := edades["juan"]
	// if  ok {
	// 	if edad > 18 {
	// 		fmt.Println("es mayor de edad")
	// 	} else {
	// 		fmt.Println("menor de edad")
	// 	}
	// } else {
	// 	fmt.Println("no existe")
	// }

	// edades["roberto"] = 66

	// // recorrer un map
	// for nombre, edad := range edades {
	// 	if edad > 18 {
	// 		fmt.Printf("%s es mayor de edad○\n", nombre)
	// 	} else {
	// 		fmt.Printf("%s es menor de edad\n", nombre)
	// 	}
	// }

	print("hola mundo")
	print(fmt.Sprint(suma(5, 5)))
	print(fmt.Sprint(resta(8, 5)))
	fmt.Println(sumar(5,6,3,6,9,5))
	fmt.Println(sumar())

	var numeros []int32 = []int32{1,2,3,4,5}
	// no acepta como parametro un slice of number
	fmt.Println((sumar(numeros...)))


	var cadena string = "primo"
	var cadenas []string = []string{"jodase", "pero", "bien", "jodido"}
	print2(cadena, cadenas...)

	clausura_interna := clausura

	print(fmt.Sprint(clausura_interna()))

	// pero tanbien podemos hacerlo de la siguiente manera
	var saludo string = "hola"
	clausura2 := func(){
		print(saludo)
	}
	clausura2()

	// funciones anonimas
	var cadena3 string = "123456789"
	cadena3 = strings.Map(func(r rune) rune {
		return r + 1
	}, cadena3)

	fmt.Println((cadena3))
	

	// recibimos multiples valores de la func
	var n1, n2, n3 int = multiplicar_varios(5)
	fmt.Println(n1, n2, n3)
}

// retornar multiples valores  
func multiplicar_varios(valor int) (n1, n2, n3 int) {
	n1 = valor * 2
	n2 = valor * 3
	n3 = valor * 4
	return
}

// no se puede crear un func dentro de otra
func clausura() int {
	return 10 + 25
}

func print(text string) {
	fmt.Println(text)
}

func suma(a int, b int) int {
	return a + b
}

func resta(a int, b int) (r int) {
	r = a - b
	return
}


// ...int usando este valor determinamos que va a recibir n cantidad de valores
func sumar(numeros ...int32) (r int32) {
	r = 0
	for _, numero := range numeros {
		r += int32(numero)
	}
	return
}

func print2(cadena string, cadenas ...string) {

	for _, c := range cadenas {
		cadena += " " + c
	}

	fmt.Println(cadena)
}

// link vide
// https://www.youtube.com/watch?v=lc5Gr5CKH6A&list=PLl_hIu4u7P64MEJpR3eVwQ1l_FtJq4a5g&index=36