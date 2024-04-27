package main

// На вход подаются два неупорядоченных слайса
// любой длины. Надо написать функцию, которая
// возвращает их пересечение

import "fmt"

func intersection(a, b []int) []int {
	counter := make(map[int]int)
	var result []int

	for _, elem := range a {
		if _, ok := counter[elem]; !ok {
			counter[elem] = 1
		} else {
			counter[elem] += 1
		}
	}

	for _, elem := range b {
		if count, ok := counter[elem]; ok && count > 0 {
			counter[elem] -= 1
			result = append(result, elem)
		}
	}

	return result
}

func main() {
	fmt.Println(intersection([]int{1, 3, 4, 4}, []int{4, 6, 1, 4}))
	// [4 1 4]
	fmt.Printf("%v\n", intersection([]int{1, 1, 1}, []int{1, 1, 1, 1}))
	// [1 1 1]
}
