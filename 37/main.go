package main

import "fmt"

/* 아이패드를 사주는 조건 [국어, 수학, 사회, 과학, 영어]
 1. 다섯 과목의 총점이 400점 이상
 2. 90점 이상 과목 수가 2개 이상
 3. 70점 미만 과목이 없어야함
*/
var sum int
var nums []int
type calE func(...int) (int, int)

func inputNums() (ret []int) {
	var input int
	for i := 0; i<5; i++{
		fmt.Scanf("%d",&input)
		if input > 100 {
			i--
		} else {
			ret = append(ret, input)
		}
	}
	
	return
}


func calExam(arr ...int) (up int, down int) {
	
	for i:=0;i<5;i++ {
		if arr[i] >= 90 {
			up++
		} else if arr[i] < 70 {
			down++
		}
		sum = sum + arr[i]
	}
	
	return
}



func printResult (f calE) {
	var result bool = true
	up, down := f(nums ...)
	if !(sum >= 400) {
		fmt.Printf("총점이 400점 미만입니다.\n")
		result = false
	} 
	if up < 2 {
		fmt.Println("90점 이상 과목 수가 2개 미만입니다.")
		result = false
	} 
	if down > 0 {
		fmt.Println("70점 미만 과목 수가 있습니다.")
		result = false
	}
	
	if result == false {
		fmt.Println("아이패드를 살 수 없습니다.")
	}else{
		fmt.Println("아이패드를 살 수 있습니다.")
	}
}

func main() {
	nums = inputNums()
	printResult(calExam)
}
