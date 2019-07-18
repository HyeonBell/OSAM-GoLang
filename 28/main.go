package main

import "fmt"

func main() {
	var A = [2][2]int{
		{7,3},
		{5,2},
	}
	d := A[0][0]*A[1][1] - A[0][1]*A[1][0]
	
	if d != 0 {
		fmt.Println("true")
		tmp := A[0][0]
		A[0][0] = (1/d)*A[1][1]
		A[1][1] = (1/d)*tmp
		tmp = A[1][0]
		A[1][0] = (1/d)*A[0][1]*(-1)
		A[0][1] = (1/d)*tmp*(-1)
		fmt.Printf("%d %d\n",A[0][0],A[1][0])
		fmt.Printf("%d %d\n",A[0][1],A[1][1])
	} else {
		fmt.Print("false")
		
	}
	
}
