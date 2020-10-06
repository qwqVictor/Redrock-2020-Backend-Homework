package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		n     int
		tmp   int
		array []int
	)
	fmt.Scanf("%d", &n)
	for i := 0; i < n; i++ {
		fmt.Scanf("%d", &tmp)
		array = append(array, tmp)
	}
	sort.Ints(array)
	for i := len(array) - 1; i >= 0; i-- {
		fmt.Printf("%d ", array[i])
	}
}
