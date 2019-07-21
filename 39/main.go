package main

import "fmt"


//함수가 함수를 반환
func next() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	a, b := 10, 20
	str := "Hello goorm!"
	
	result := func () int{ // 익명함수 변수에 초기화
		return a + b // main 함수 변수 바로 접근
	}()
	
	func() {
		fmt.Println(str) // main 함수 변수 바로 접근
	}()
 
	fmt.Println(result)	
	
	//함수가 함수를 반환
	nextInt := next() // next안의 i가 0으로 초기화 후 nextInt는 i를 +1하는 함수를 가지고 있음

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInt := next() // 다시 새로운 next가 i를 0으로 초기화후 i를 +1함
	fmt.Println(newInt())
}

