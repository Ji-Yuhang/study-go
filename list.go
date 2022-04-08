package main

import "fmt"

func list() {
	list := make([]int, 0)

	array := []int{4, 5, 9}

	list = append(list, array...)

	for _, val := range list {
		fmt.Println("val: ", val)
	}

	fmt.Println(list)
}
