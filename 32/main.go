package main

import "fmt"

func exampleFunc1() {
	var ab int = 10 //지역변수 선언
	
	ab++
	
	fmt.Println("exampleFunc1의 ab는 ", ab)
}

func exampleFunc2() {
	var b int = 20 //지역변수 선언
	var c int = 30 //지역변수 선언

	b++
	c++

	fmt.Println("b와 c는 ", b, c)
}

var a int = 1 //전역변수 선언

func localVar() int { //지역변수로 연산
	var a int = 10 //지역변수 선언

	a += 3

	return a
}

func globalVar() int { //전역변수로 연산
	a += 3
		
	return a
}


func main() {
	var ab int = 28 //지역변수 선언

	exampleFunc1()
	exampleFunc2()

	fmt.Println("main의 ab는", ab)
	fmt.Println("지역변수 a의 연산: ", localVar())
	fmt.Println("전역변수 a의 연산: ", globalVar())
}

