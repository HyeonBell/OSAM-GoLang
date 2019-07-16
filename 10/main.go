package main

import "fmt"

func main() {
	var num1, num2, num3 int
	
	fmt.Scanln(&num1, &num2, &num3)
	Nnum1 := float32(num1)
	Nnum3 := int64(num3)
	
	fmt.Printf("%T, %f\n",Nnum1, Nnum1)
	fmt.Printf("%T, %d\n",uint(num2), uint(num2))
	fmt.Printf("%T, %d\n",Nnum3, Nnum3)
	
	
}
