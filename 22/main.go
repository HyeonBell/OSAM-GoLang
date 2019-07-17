package main

import "fmt"

func main() {
	i := 0

Hyeonbell:
	for {
		if i == 0 {
			break Hyeonbell
		}
	}
	
	fmt.Println("End")
	var sum = 0
	var j int
	
	j = 1
	
	for {
		sum += j
		if sum > 100 {
			break
		}
		j++
	}
	
	fmt.Println("1에서 ", j, " 까지 더하면 처음으로 100이 넘어요.")
	fmt.Println("합계:", sum)

	for y := 1; y < 16; y++ {
		if y%2 == 0 {
			continue //반복문 처음 부분으로 간다
		}
		
		fmt.Printf("%d ", y)
	}

	var num int

	fmt.Print("자연수를 입력하세요:")
	fmt.Scanln(&num)

	if num == 1 {
		goto ONE
	} else if num == 2 {
		goto TWO
	} else {
		goto OTHER
	}

ONE:
	fmt.Print("1을 입력했습니다.\n")
	goto END
TWO:
	fmt.Print("2를 입력했습니다.\n")
	goto END
OTHER:
	fmt.Print("1이나 2를 입력하지 않으셨습니다!\n")
END:


}
