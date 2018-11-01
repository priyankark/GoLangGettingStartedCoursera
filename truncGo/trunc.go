package main

import "fmt"

func main() {
	var x float32 = 0
	fmt.Println("Please enter a floating point number:")
	fmt.Scan(&x)
	y := int32(x)
	fmt.Printf("The converted number after truncation is %d", y)
}
