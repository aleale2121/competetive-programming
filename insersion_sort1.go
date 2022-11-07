package main

import (
	"fmt"
	"strings"
)


func insertionSort1(n int32, arr []int32) {
	elem := arr[n-1]
	i := n - 2
	for i >= 0 && arr[i] > elem {
		arr[i+1] = arr[i]
		fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
		i--
	}
	arr[i+1] = elem
	fmt.Println(strings.Trim(fmt.Sprint(arr), "[]"))
}
