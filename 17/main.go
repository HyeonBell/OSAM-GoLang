package main
 
import "fmt"
 
func main() {
	
	/* 잘못된 조건문
	fmt.Print("정수입력 :")
	fmt.Scan(&num)

	if num == 1
	{
		fmt.Print("hello\n")
	}
	else if num == 2
	{
		fmt.Print("world\n")
	} else
	{
		fmt.Print("worng number..\n")
	}
	*/
	// 올바른 블록 시작 브레이스를 둔것 
	var num int
	
	fmt.Print("정수입력 :")
	fmt.Scan(&num)

	if num == 1 {
		fmt.Print("hello\n")
	} else if num == 2 {
		fmt.Print("world\n")
	} else {
		fmt.Print("worng number..\n")
	}
	// 조건식에 간단한 문장(Optional Statement) 실행 가능

	var num2 int

	fmt.Print("정수입력 :")
	fmt.Scan(&num2)

	if val := num2 * 2; val == 2 {
		fmt.Print("hello\n")
	} else if val := num2 * 3; val == 6 {
		fmt.Print("world\n")
	} else {
		fmt.Print("worng number..\n")
	}
	
	// 계산

	var opt int
	var num3, num4, result float32

	fmt.Print("1.덧셈 2.뺄셈 3.곱셈 4.나눗셈 선택:")
	fmt.Scan(&opt)
	fmt.Print("두 개의 실수 입력:")
	fmt.Scan(&num3, &num4)

	if opt == 1 {
		result = num3 + num4
	} else if opt == 2 {
		result = num3 - num4
	} else if opt == 3 {
		result = num3 * num4
	} else if opt == 4 {
		result = num3 / num4
	}

	fmt.Printf("결과: %f\n", result)


}
