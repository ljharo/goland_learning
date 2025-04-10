package main

import (
	"fmt"
)

// // type
// // declaramos nuestro propio tipo de dato
// type Dinero int32

// // declaramos el metodo string para el tipo dinero
// func (d Dinero) String() string {
// 	return fmt.Sprintf("$%d", d)
// }

// type User struct {
// 	id       int32
// 	nombre   string
// 	dinero   Dinero
// 	edad     int8
// }

// func imprimir(){
// 	fmt.Println("hola mundo")
// 	// panic("error")

// 	// cadena := recover()

// 	defer func (){
// 		cadena := recover()
// 		fmt.Println(cadena)
// 	}()
// 	panic("error")		
// }


// func incrementar(numero *int){
// 	*numero++
// 	fmt.Printf("se ha incrementado el valor a %d\n", *numero)
// 	return
// }

// func main () {
	// imprimir()

	// // punteros
	// var a int8 = 25
	// fmt.Println("valor de a: ", a)
	// fmt.Println("direccion de a: ", &a)

	// // asignamos el a to b
	// b := &a
	// fmt.Println("direccion de b: ", b)
	// fmt.Println("valor de b: ", *b)

	// // b es del tipo *int no int
	// // b = 25

	// // le asignamos un nuevo valor a (a) a traves de b
	// *b = 35

	// fmt.Println("nuevo valor de a: ", a)
	// a++ // incrementamos el valor de a
	// fmt.Println("nuevo valor de b: ", *b)

	// // valor 0 de los punteros de nil
	// if b != nil {
	// 	fmt.Println("el puntero b es nil")
	// }

	// // los punteros son comparables
	// c := &a
	// if b == c {
	// 	fmt.Println("los punteros b y c son iguales")
	// }

	// // utilizamos new para iniciar un puntero vacio
	// d := new(int8)
	// fmt.Println("direccion de d: ", d)
	// fmt.Println("valor de d: ", *d)

	// d = b
	// fmt.Println("direccion de d: ", d)
	// fmt.Println("valor de d: ", *d)

	// numero := 10
	// fmt.Println("el valor de numero es: ", numero)
	// incrementar(&numero)
	// fmt.Println("el valor de numero es: ", numero)

// 	var sueldo Dinero = 25000
// 	fmt.Println("Sueldo: ", sueldo)

// 	var aumento int32 = 10000
// 	// no se puede hacer la suma porque no son del mismo tipo
// 	sueldo += Dinero(aumento)
// 	fmt.Println(sueldo)


// 	var user User = User{
// 		id: 1,
// 		nombre: "Juan",
// 		edad: 25,
// 		dinero: 2500,
// 	}

// 	fmt.Println(user)

// }

// estructura 
type Rectangulo struct {
	alto, ancho int
}

// aplicamos un metodo

func (r Rectangulo) area() int {
	return r.alto * r.ancho
}

func main(){
	var rectangulo Rectangulo = Rectangulo{5,5}
	var area int = rectangulo.area()
	fmt.Println("El area del rectangulo es: ", area)
}

// // link vide
// // https://www.youtube.com/watch?v=wGMAA5dabYc&list=PLl_hIu4u7P64MEJpR3eVwQ1l_FtJq4a5g&index=45