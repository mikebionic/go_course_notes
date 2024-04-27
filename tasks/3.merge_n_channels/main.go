package main

// Слить n каналов в один
// Даны n каналов, надо написать функцию,
// которая смерджит все данные из этих
// каналов в один и вернет его

import (
	"fmt"
	"sync"
)

func joinChannels(chs ...chan int) <-chan int {
	mergedCh := make(chan int)
	go func() {
		wg := &sync.WaitGroup{}
		wg.Add((len(chs)))

		for _, ch := range chs {
			go func(ch <-chan int, wg *sync.WaitGroup) {
				defer wg.Done()
				for id := range ch {
					mergedCh <- id
				}
			}(ch, wg)
		}
		wg.Wait()
		close(mergedCh)
	}()
	return mergedCh
}

func main() {
	a := make(chan int)
	b := make(chan int)
	c := make(chan int)
	go func() {
		for _, num := range []int{1, 2, 3} {
			a <- num
		}
		close(a)
	}()
	go func() {
		for _, num := range []int{20, 10, 30} {
			b <- num
		}
		close(b)
	}()
	go func() {
		for _, num := range []int{300, 400, 200} {
			c <- num
		}
		close(c)
	}()
	for num := range joinChannels(a, b, c) {
		fmt.Println(num)
	}
}
