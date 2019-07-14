package main

import "fmt"

func main() {
	var num1, num2, num3 int
	
	fmt.Print("정수 3개를 입력하세요 :")
	fmt.Scanln(&num1, &num2, &num3)
	fmt.Println(num1, num2, num3)
}
