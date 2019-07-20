package main

import "fmt"

var count int

func bubbleSort(nums ...int) {
	var temp int
	for j:=0 ; j<count ; j++ {
		for i:=0 ; i<(count-1) ; i++ {
			if nums[i] > nums[i+1] {
				temp = nums[i+1]
				nums[i+1] = nums[i]
				nums[i] = temp
			}
		}
	}
}

func inputNums() (ret []int) {
	var input int
	for i:=0; i<count ; i++ {
		fmt.Scanf("%d",&input)
		ret = append(ret, input)
	}
	
	return 
}

func outputNums(nums ...int) {
	for i:=0; i<count ; i++ {
		fmt.Printf("%d ",nums[i])
	}
}

func main() {
	fmt.Scanf("%d", &count)
	nums := inputNums()
	bubbleSort(nums...)
	outputNums(nums...)
}
