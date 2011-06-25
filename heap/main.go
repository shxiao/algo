package main

import (
	"fmt"
	"rand"
)


//使元素i的子树符合最大堆规则
func maxHeapify(data []int, i int){
	left := 2*i+1
	right := 2*i + 2
	largest := i
	if(left < len(data) && data[left] > data[largest]){
		largest = left;
	}
	if right < len(data) && data[right] > data[largest] {
		largest = right;
	}
	if largest != i {
		data[i],data[largest] = data[largest],data[i]
		maxHeapify(data,largest)
	}
}

func maxHeapifyIter(data []int, i int){
	largest := i
	pre := largest
	for {
		left := 2*largest + 1;
		right := 2*largest + 2;
		if left < len(data) && data[left] > data[largest] {
			largest = left
		}
		if right < len(data) && data[right] > data[largest] {
			largest = right;
		}
		if (largest != pre){
			data[largest],data[pre] = data[pre],data[largest]
			pre = largest
		}else{
			break
			fmt.Println("break")
		}
	}
}

func buildMaxHeap(data []int){
	for i := len(data)/2; i >=0; i-- {
		maxHeapifyIter(data,i)
	}
}

func heapSort(data []int){
	if len(data) <= 1 {
		return 
	}
	buildMaxHeap(data)
	heapSort(data[1:])
}

func randInts() []int{
	r := rand.Int() % 30
	d := make([]int, r)
	for i := 0; i < r; i++ {
		d[i] = rand.Int()%30
	}
	return d
}


func main(){
	ints := randInts() 
	fmt.Println(ints)
	heapSort(ints)
	fmt.Println(ints)
}