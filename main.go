package main

import (
	"fmt"

	"go-template/pkg/calculator"
)

func main() {
	c := calculator.NewCalculator(5, 5)
	fmt.Println(c.GetAnswer())
}
