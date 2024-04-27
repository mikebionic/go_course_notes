package main

// Написать генератор случайных чисел
// Для решения использован небуферизированный канал.
// Асинхронно пишем туда случайные чила и закроем
// его, когда закончим писать

import (
	"fmt"
	"math/rand"
	"time"
)

func randNumsGenerator(n int) <-chan int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	out := make(chan int)
	fmt.Println(r.Intn(n))

	go func() {
		for i := 0; i < n; i++ {
			out <- r.Intn(n)
		}
		close(out)
	}()
	fmt.Println()
	return out
}

func main() {
	for num := range randNumsGenerator(4) {
		fmt.Println(num)
	}
}
