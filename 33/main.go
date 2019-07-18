package main

import "fmt"

//Pass by value
func printSqureV(a int) {
	a *= a
		
	fmt.Println(a)
}
//Pass by reference
func printSqureR(b *int) {
	*b *= *b
	
	fmt.Println(*b)
}

// 가변 인자 매개 변수
func addOne(num ...int) int {
	var result int

	for i := 0; i < len(num); i++ { //for문을 이용한 num[i] 순차 접근
		result += num[i]
	}
	
	return result
}

// 가변 인자 매개 변수
func addTwo(num ...int) int {
	var result int

	for _, val := range num { //for range문을 이용한 num의 value 순차 접근
		result += val
	}
	
	return result
}

func main() {
	
	a := 4 //지역변수 선언
		
	printSqureV(a)
		
	fmt.Println(a)
	
	b := 4         //지역변수 선언
	
	printSqureR(&b) //참조를 위한 b의 주솟값을 매개변수로 전달
	
	fmt.Println(b)
	
	num1, num2, num3, num4, num5 := 1, 2, 3, 4, 5
	nums := []int{10, 20, 30, 40}

	fmt.Println(addOne(num1, num2))
	fmt.Println(addOne(num1, num2, num4))
	fmt.Println(addOne(nums...))
	fmt.Println(addTwo(num3, num4, num5))
	fmt.Println(addTwo(num1, num3, num4, num5))
	fmt.Println(addTwo(nums...))
}

