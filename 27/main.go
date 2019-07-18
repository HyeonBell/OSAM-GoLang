package main

import "fmt"

func main() {
	var a map[int]string

	if a == nil {
		fmt.Println("nil map")
	}

	var d = map[string]string{ //key:value, 형식으로 초기화한다
		"apple":  "red",
		"grape":  "purple",
		"banana": "yellow",
	}

	fmt.Println(d, "\nm의 길이는", len(d))

	//지역번호와 지역 저장
	var u = make(map[string]string)

	u["02"] = "서울특별시"
	u["031"] = "경기도"
	u["032"] = "충청남도"
	u["053"] = "대구광역시"

	fmt.Println(u)

	//동일한 key값으로 value값을 저장하면 갱신이 된다
	u["032"] = "인천"

	fmt.Println(u)

	//m에 있는 "031"key의 value와 함께 삭제
	delete(u, "031")

	fmt.Println(u)

	//지역번호와 지역 저장
	var m = make(map[string]string)

	m["02"] = "서울특별시"
	m["031"] = "경기도"
	m["032"] = "인천"
	m["053"] = "대구광역시"

	fmt.Println(m["032"])
	fmt.Println(m["042"], "빈 칸입니다.") //string형태로 존재하지 않는 key값은 ""가 출력된다

	val, exist := m["02"] //존재하는 key
	fmt.Println(val, exist)

	val, exist = m["042"] //존재하지 않는 key
	fmt.Println(val, exist)

	val = m["053"] //value 값만 반환
	fmt.Println(val)

	_, exist = m["053"] //true/false 값만 반환
	fmt.Println(exist)

	//맵도 똑같이 len() 함수를 사용할 수 있다 하지만 cap() 함수는 사용할 수 없다
	fmt.Println(len(m))


}
