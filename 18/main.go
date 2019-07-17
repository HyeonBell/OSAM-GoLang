package main

import "fmt"

func main() {
	for index:=1;index<101;index++ {
		
		if index%7 == 0 {
			fmt.Print(index)
			fmt.Print(" ")
		} else if index%9 == 0 {
			fmt.Print(index)
			fmt.Print(" ")
		}
		
	}
	
}
