package main

import "fmt"

// // replacement for inheriance - aggregation

// type Bird struct {
// 	Age int
// }

// func (b *Bird) Fly() {
// 	if b.Age >= 10 {
// 		fmt.Println("Flying!")
// 	}
// }

// type Lizard struct {
// 	Age int
// }

// func (l *Lizard) Crawl() {
// 	if l.Age < 10 {
// 		fmt.Println("Crawling")
// 	}
// }

// type Dragon struct {
// 	Bird
// 	Lizard
// 	// this can caught problems because both of the structs have an AGE field
// 	// this can be a problem to combine the operation
// 	// ERR: Ambigious selector
// }

// func (d *Dragon) Age() int {
// 	return d.Bird.Age
// }
// func (d *Dragon) SetAge(age int) {
// 	d.Bird.Age = age
// 	d.Lizard.Age = age
// }

/////////////////////

// Age Getter and Setter
type Aged interface {
	Age() int
	SetAge(age int)
}

type Bird struct {
	age int
}

func (b *Bird) Age() int {
	return b.age
}
func (b *Bird) SetAge(age int) {
	b.age = age
}

func (b *Bird) Fly() {
	if b.age >= 10 {
		fmt.Println("Flying!!")
	}
}

type Lizard struct {
	age int
}

func (l *Lizard) Age() int {
	return l.age
}
func (l *Lizard) SetAge(age int) {
	l.age = age
}

func (l *Lizard) Crawl() {
	if l.age < 10 {
		fmt.Println("Crawling!!")
	}
}

type Dragon struct {
	// operate on behaviours, cant access directly
	bird   Bird
	lizard Lizard
}

func (d *Dragon) Age() int {
	return d.bird.age
}

func (d *Dragon) SetAge(age int) {
	d.bird.SetAge(age)
	d.lizard.SetAge(age)
}

func (d *Dragon) Crawl() {
	d.lizard.Crawl()
}
func (d *Dragon) Fly() {
	d.bird.Fly()
}

func NewDragon() *Dragon {
	return &Dragon{Bird{}, Lizard{}}
}

func main() {
	d := Dragon{}
	d.SetAge(5)
	d.Fly()
	d.Crawl()
}
