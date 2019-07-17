package main

import "fmt"

func main() {
	var i,j int
	
	for i=2;i<10;i++{
		for j=1;j<10;j++{
			if i%2 == 1{
				fmt.Printf("%d x %d = %d\n",i,j,i*j)
				if i == j {
					fmt.Print("\n")
					break
				}
			}
		}
	}
}
