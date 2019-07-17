package main

import "fmt"

func main() {
	
	for i:=0;i<10;i++ {
		for j:=0;j<10;j++ {
			if (i*10+j)+(j*10+i) == 99{
				fmt.Printf("%d%d + %d%d = 99\n",i,j,j,i)
			}
		}
	}
}
