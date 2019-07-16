package main

import "fmt"

func main() {
	var dan int
	fmt.Scanf("%d",&dan)
	
		
	for index:=1 ; index <= 9; index++ {
		fmt.Printf("%d X %d = %d\n", dan, index, dan*index)
	}
}
