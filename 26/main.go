package main

import "fmt"

func main() {
	var a []int        //슬라이스 변수 선언 아무것도 초기화 되지 않은 상태
	a = []int{1, 2, 3} //슬라이스에 리터럴값 지정

	a[1] = 10 //값이 할당되어 메모리가 생겼기 때문에 이렇게 접근 가능

	fmt.Println(a)

	var b []int //nil slice 선언

	if b == nil {
		fmt.Println("용량이", cap(b), "길이가", len(b), " Nil Slice입니다.")
	}

	s := make([]int, 0, 3) // len=0, cap=3 인 슬라이스 선언

	for i := 1; i <= 10; i++ { // 1부터 차례대로 한 요소씩 추가
		s = append(s, i)

		fmt.Println(len(s), cap(s)) // 슬라이스 길이와 용량 확인
	}

	fmt.Println(s) // 최종 슬라이스 출력

	//슬라이스에 슬라이스를 추가해서 붙이기.
	
    sliceA := []int{1, 2, 3}
    sliceB := []int{4, 5, 6}
 
    sliceA = append(sliceA, sliceB...)
    //sliceA = append(sliceA, 4, 5, 6)
 
    fmt.Println(sliceA) // [1 2 3 4 5 6] 출력


	//슬라이스를 잘라서 복사
	
	c := make([]int, 0, 3) //용량이 3이고 길이가0인 정수형 슬라이스 선언
	c = append(c, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println(len(c), cap(c))

	l := c[1:3] //인덱스 1요소부터 2요소까지 복사
	fmt.Println(l)

	l = c[2:] //인덱스 2요소부터 끝까지 복사
	fmt.Println(l)

	l[0] = 6

	fmt.Println(c) //슬라이스 l의 값을 바꿨는데 c의 값도 바뀜
	//값을 복사해온 것이 아니라 기존 슬라이스 주솟값을 참조

}
