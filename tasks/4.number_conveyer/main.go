package main

// Конвейер чисел
// Даны два канала, в первый пишутся числа.
// Нужно чтобы числа читались из первого по мере поступления,
// что-то с ними происходило (допустим возводились в квадрат)
// и результат записывался во второй канал

import "fmt"

func main() {
	naturals := make(chan int)
	squares := make(chan int)

	go func() {
		for x := 0; x <= 10; x++ {
			naturals <- x
		}
		close(naturals)
	}()

	go func() {
		for x := range naturals {
			squares <- x * x
		}
		close(squares)
	}()
	for x := range squares {
		fmt.Println(x)
	}

}
