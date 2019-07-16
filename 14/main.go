package main
 
import "fmt"
 
func main() {
	sum := 0
	
	for i := 1; i <= 10; i++ {
		sum += i
	}
	fmt.Println("1부터 10까지 정수 합계:", sum)
	
	n := 2
	
	for n < 100 {
		fmt.Printf("count %d\n", n)
		
		n *= 2
	}

	/*
	for {
		fmt.Printf("무한루프입니다.\n")
	}
	*/

	var arr [6]int = [6]int{1, 2, 3, 4, 5, 6}

	for index, num := range arr {
		fmt.Printf("arr[%d]의 값은 %d입니다.\n", index, num)
	}
//인덱스를 생략한 예시

	var actors [4]string = [4]string{"정우성", "류준열", "박보검", "이정재"}

	for _, actor := range actors {
		fmt.Printf("제가 좋아하는 배우는 %s입니다.\n", actor)
	}

//요소값을 생략한 예시

	var actors2 [4]string = [4]string{"정우성", "류준열", "박보검", "이정재"}

	for index := range actors2 {
		fmt.Printf("배우가 %d명 입장했습니다.\n", index+1)
	}

//컬랙션의 맵활용
	
	var fruits map[string]string = map[string]string{
		"apple":  "red",
		"banana": "yellow",
		"grape":  "purple",
	}

	for fruit, color := range fruits {
		fmt.Printf("%s의 색깔은 %s입니다.\n", fruit, color)
	}

}
