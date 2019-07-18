package main

import "fmt"

func main() {
	
	var name string
	
	names := make([]string,0,1)
	
	for {
		fmt.Scanf("%s",&name)
		if name == "1" {
			break
		} else {
			names = append(names, name)
		}
	}
	var long string = names[0]
	for i:=0 ; i<len(names); i++ {
		if len(long) < len(names[i]) {
			long = names[i]
		}
	}
	
	fmt.Printf("%s %d\n",long,len(long))
}
