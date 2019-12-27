/* Дан периметр окружности d.package main

найти её длину L=pi*d
в качестве pi использовать 3.14
*/

package main

import "fmt"

func main() {
	const pi = 3.14
	var d float64
	var L float64

	fmt.Println("Введите диаметр окружности :")
	fmt.Scanf("%f", &d)
	L = pi * d
	fmt.Println("Длина окружности равна L=", L)

}
