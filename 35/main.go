package main

import "fmt"

//함수 원형 정의
type calculatorNum func(int, int) int 
type calculatorStr func(string, string) string

func calNum(f calculatorNum, a int, b int) int {
	result := f(a, b)
	return result
}

func calStr(f calculatorStr, a string, b string) string {
	sentence := f(a, b)
	return sentence
}

func addDeclared(nums ...int) (result int) {
	for i := 0; i < len(nums); i++ {
		result += nums[i]
	}
	return
}

func add() {
	fmt.Println("선언 함수를 호출했습니다.")
}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func main() {
	func() {
		fmt.Println("hello")
	}()

	func(a int, b int) {
		result := a + b
		fmt.Println(result)
	}(1, 3)

	result := func(a string, b string) string {
		return a + b
	}("hello", " world!")
	fmt.Println(result)

	i, j := 10.2, 20.4
	divide := func(a float64, b float64) float64 {
		return i / j
	}(i, j)
	fmt.Println(divide)
	
	var nums = []int{10, 12, 13, 14, 16}
	
	//변수에 초기화한 익명 함수는 변수 이름을 함수의 이름처럼 사용 가능
	addAnonymous := func(nums ...int) (result int) {
		for i := 0; i < len(nums); i++ {
			result += nums[i]
		}
		return
	}

	fmt.Println(addAnonymous(nums...))
	fmt.Println(addDeclared(nums...))
	
	// 선언함수와 익명함수가 읽히는 차이
	add := func() {
		fmt.Println("익명 함수를 호출했습니다.")
	}

	add()
	
	multi := func(i int, j int) int {
		return i * j
	}
	
	r1 := calc(multi, 10, 20)
	fmt.Println(r1)

	r2 := calc(func(x int, y int) int { return x + y }, 10, 20)
	fmt.Println(r2)

	multi2 := func(iq int, jq int) int {
		return iq * jq
	}
	duple := func(iq string, jq string) string {
		return iq + jq + iq + jq
	}

	r3 := calNum(multi2, 10, 20)
	fmt.Println(r3)

	r4 := calStr(duple, "Hello", " Golang ")
	fmt.Println(r4)
}

	
	
	

