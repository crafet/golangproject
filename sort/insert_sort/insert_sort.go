package insert_sort

import (
	"fmt"
)

// insert sort
func Sort(arr [] int) {
	size := len(arr)
	for i:=1; i<size; i++ {
		for j:=i-1; j>=0; j-- {
			if arr[j+1] < arr[j] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}
}

func Display(arg []int) {
	for _, e := range arg {
		fmt.Print(e, " ")
	}
	fmt.Println()
}
