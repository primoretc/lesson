//дана сторона квадрата а.
//найдите периметр  P=4*a
package main

import "fmt"

func main() {
	var a float64
	fmt.Println("Введите сторону квадрата ")
	fmt.Scanf("%f", &a)

	fmt.Println("Периметр квадрата равен ", (4 * a))

}
