package main

import (
	"fmt"
)

func main() {
	var cod1, cod2, num1, num2 int
	var vl1, vl2 float64

	fmt.Scanf("%d %d %f\n", &cod1, &num1, &vl1)
	fmt.Scanf("%d %d %f", &cod2, &num2, &vl2)

	total := float64(num1)*vl1 + float64(num2)*vl2

	fmt.Printf("VALOR A PAGAR: R$ %.2f\n", total)
}
