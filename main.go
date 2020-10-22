package main

import "fmt"

func main() {
	fmt.Println("Introduzca el numero de la opcion deseada")
	fmt.Println("1.- Fibonacci")
	fmt.Println("2.- Funcion Sums (retorna numero mayor)")
	fmt.Println("3.- Generar impares")
	fmt.Println("4.- Intercambio de numeros con punteros")
	fmt.Println("5.- Salir")

	var opt int
	fmt.Scanln(&opt)
	switch opt {
	case 1:
		fmt.Println("Ingrese el numero de fibonacci que desea")
		var fib uint64
		fmt.Scanln(&fib)
		fmt.Println("El resultado es: ", Fibonacci(fib))
		main()
	case 2:
		var a, b int64
		fmt.Println("Ingrese el primer numero")
		fmt.Scanln(&a)
		fmt.Println("Ingrese el segundo numero")
		fmt.Scanln(&b)
		fmt.Println("El numero mas alto ingresado es: ", Sums(a,b))
		main()
	case 3:
		var n uint64
		fmt.Println("Cuantos numeros impares desea?")
		fmt.Scanln(&n)
		nextImpar:= generadorImpares()
		for i:= 1; i<= int(n); i++{
			fmt.Println(nextImpar())
		}
		main()
	case 4:
		var a, b int64
		fmt.Println("Ingrese el primer numero")
		fmt.Scanln(&a)
		fmt.Println("Ingrese el segundo numero")
		fmt.Scanln(&b)
		Intercambiar(&a,&b)
		fmt.Println("Despues del cambio el primer numero es:", a)
		fmt.Println("Despues del cambio el segundo numero es:", b)
		main()
	case 5:
		fmt.Println("Gracias.")
	default:
		fmt.Println("Opcion no disponible")
	}
}

func Fibonacci(x uint64) uint64 {

	if x < 2 {
		return x
	} else {
		return (Fibonacci(x-1) + Fibonacci(x-2))
	}
}

func Sums(num ...int64) int64 {
	if num[0] > num[1] {
		return num[0]
	} else {
		return num[1]
	}
}

func generadorImpares() func() uint64 {
	i := uint64(1) // i permanecerá en el clousure de la función anónima a retornar
	return func() uint64 {
		var par = i
		i += 2
		return par
	}
}

func Intercambiar(a, b *int64) {
	var temp = *a
	*a = *b
	*b = temp
}
