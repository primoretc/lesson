package main

import (
	"fmt"
)

type vehicle interface {
	move()
}

func main() {
	fmt.Println("Hello Go")
}
