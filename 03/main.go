package main

import "fmt"

const username="hyeonbell"

const(
	c1 = 10
	c2
	c3
	c4 = "bell"
	c5 
	c6 = iota
	c7
	c8
	c9 = "earth"
	c10
	c11 = "End"
)

func main(){
	
	fmt.Println(c1, c2, c3, c4, c5, c6, c7, c8, c9,c10, c11)
	fmt.Println(username)
}
