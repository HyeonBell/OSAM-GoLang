package main

import "fmt"

func main() {
	var i, j int
	var length int
	fmt.Scanf("%d", &length)
	
	j=0
	for i=0; i<length; i++ {
		for index2:=0; index2<j; index2++ {
			fmt.Print("o ")
		}
		fmt.Print("* \n")
		j++
	}
}
