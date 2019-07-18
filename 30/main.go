package main

import "fmt"

func main() {
	
	var middleTest = make(map[string]int)
	var avg float32
	var subject string
	var point int
	
	for {
		fmt.Scanf("%s %d",&subject, &point)
		if subject == "0" {
			break
		} else {
			middleTest[subject] = point
		}
	}
	for sub, poi := range middleTest {
		fmt.Printf("%s %d\n",sub,poi)
		avg = avg + float32(poi)
	}
	fmt.Printf("%.2f\n",avg/float32(len(middleTest)))
	

	

	
	
}
