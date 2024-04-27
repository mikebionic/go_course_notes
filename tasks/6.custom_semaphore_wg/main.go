package main

import "fmt"

// Сделать кастомную WaitGroup на семафоре
// Семафор можно получить из канала.
// Чтобы не аллоцировать лишние данные,
// будем складывать туда пустые структуры.
// В нашем случае мы хотим сделать семафор,
// который бует ждать выполнения пяти горутин.
// Для этого просто добавим вместо обычного канала
// буфферизированный. И внутри каждой горутины
// положим в него значение. А в конце будем дожидаться,
// что все ок - мы вычитаем все значения из канала

type sema chan struct{}

func New(n int) sema {
	return make(sema, n)
}

func (s sema) Inc(k int) {
	for i := 0; i < k; i++ {
		s <- struct{}{}
	}
}

func (s sema) Dec(k int) {
	for i := 0; i < k; i++ {
		<-s
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	n := len(numbers)
	sem := New(n)
	for _, num := range numbers {
		go func(n int) {
			fmt.Println(n)
			sem.Inc(1)
		}(num)
	}
	sem.Dec(n)
}
