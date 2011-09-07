package main

import (
	"fmt"
)
// 2.2-2 
func selectSort(data []int){
	for i := 0; i < len(data)-1; i++{          // n-1
		smallest := i                        // n-1
		for j :=i+1 ; j < len(data); j ++{   // j = 2...n-1 ==> 2+...+n-1
			if data[smallest] > data[j] {   // 2 + ...+ n-1
				smallest = j               // 2+ ...+ n-1
			}
		}
		data[i],data[smallest] = data[smallest],data[i] // n-1
	}
}

func main(){
	t := []int{2,3,66,4,3,66,33,1,2,7}
	selectSort(t)
	fmt.Println(t)
}
