package main

import "fmt"

/*
 물체 수직 낙하시 역학적 에너지 보존
	꼭대기 : 최고 위치에너지 최저 운동에너지
	바닥에 닿기 직전 : 최고 운동에너지 최저 위치에너지
	질량 m 속도 v 높이 h
	익명함수 일급 함수 특성 활용 위치 에너지, 운동에너지, 역학적 에너지를 구하는 실습
	운동에너지 공식 : K = 1/2mv^2 
	위치에너지 공식 : P = mgh(g = 9.8m/s^2)
	역학에너지 공식 : X = K + P
*/

//TODO 중력 (9.8)상수로 초기화
const g float32 = 9.8

//TODO type문으로 함수 원형 정의
type defEnergy func(float32, float32) float32

func calMechEnergy(f defEnergy, a float32, b float32) float32 {
	result := f(a, b)
  return result
}

func main() {
	var m, v, h float32
	fmt.Scanln(&m, &v, &h)
	
	//익명함수
	kinEnergy := func(paramM float32,paramV float32) float32{
		//운동에너지 계산
		return (paramM*paramV*paramV)/2
	}
	//익명함수
	potEnergy := func(paramM float32, paramH float32) float32 {
		//위치에너지 계산
		return paramM*paramH*g
	}
	
	ke := calMechEnergy(kinEnergy, m, v)
	pe := calMechEnergy(potEnergy, m, h)
	fmt.Printf("%.3f %.2f %.3f\n",ke,pe,ke+pe)
}
