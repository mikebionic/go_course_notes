package main

// Написать WorkerPool с заданной функцией.
// Нужно разбить процессы на несколько горутин -
// при этом не создавать новую горутину каждый раз,
// а просто переиспользовать уже имеющиеся.
// Для этого создадим канал с джобами и результирующий
// канал. Для каждого воркера создадим горутину, который
// будет ждать новую джобу, применять к ней заданную
// функцию и пулять ответ в результирующий канал.

import (
	"fmt"
)

func worker(id int, f func(int) int,
	jobs <-chan int, results chan<- int) {
	for j := range jobs {
		results <- f(j)
	}
}

func main() {
	const numJobs = 5
	jobs := make(chan int, numJobs)
	results := make(chan int, numJobs)
	multiplier := func(x int) int {
		return x * 10
	}
	for w := 1; w <= 3; w++ {
		go worker(w, multiplier, jobs, results)
	}
	for j := 1; j <= numJobs; j++ {
		jobs <- j
	}
	close(jobs)

	for i := 1; i <= numJobs; i++ {
		fmt.Println(<-results)
	}
}
