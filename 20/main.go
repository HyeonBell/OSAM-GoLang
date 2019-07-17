package main
 
import "fmt"
 
func main() {
	var num int
	fmt.Print("정수 입력:")
	fmt.Scanln(&num)
	
	switch num {
	case 0:
		fmt.Println("영")
	case 1:
		fmt.Println("일")
	case 2:
		fmt.Println("이")
	case 3:
		fmt.Println("삼")
	case 4:
		fmt.Println("사")
	default:
		fmt.Println("모르겠어요.")
	}

	//switch에 전달되는 인자로 태그 사용

	var fruit string
	
	fmt.Print("apple, banana, grape중에 하나를 입력하시오:")
	fmt.Scanln(&fruit)
	
	if (fruit != "apple") && (fruit != "banana") && (fruit != "grape") {
		fmt.Println("잘못 입력했습니다.")
		return
	}

	switch fruit {
	case "apple":
		fmt.Println("RED")
	case "banana":
		fmt.Println("YELLOW")
	case "grape":
		fmt.Println("PURPLE")
	}

	//switch에 전달되는 인자로 표현식 사용

	var num2 int
	var result string
	
	fmt.Print("10, 20, 30중 하나를 입력하시오:")
	fmt.Scanln(&num2)

	switch num2 / 10 { //표현식
	case 1:
		result = "A"
	case 2:
		result = "B"
	case 3:
		result = "C"
	default:
		fmt.Println("모르겠어요.")
		return
	}
	
	fmt.Println(result)

	//switch에 전달되는 인자 없이 case에 표현식 사용(참/거짓 판별)

	var a, b int

	fmt.Print("정수 a와 b를 입력하시오:")
	fmt.Scanln(&a, &b)

	switch {
	case a > b:
		fmt.Println("a가 b보다 큽니다.")
	case a < b:
		fmt.Println("a가 b보다 작습니다.")
	case a == b:
		fmt.Println("a와 b가 같습니다.")
	default:
		fmt.Println("모르겠어요.")
	}

}
