package main

import (
	"container/list"
	"fmt"
)

// Observable, Observer
// When person becomes Ill they send notification to Doctor
// Observable is patient Observer is doctor

type Observable struct {
	subs *list.List
}

func (o *Observable) Subscribe(x Observer) {
	o.subs.PushBack(x)
}

func (o *Observable) Unsubscribe(x Observer) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		if z.Value.(Observer) == x {
			o.subs.Remove(z)
		}
	}
}

func (o *Observable) Fire(data interface{}) {
	for z := o.subs.Front(); z != nil; z = z.Next() {
		z.Value.(Observer).Notify(data)
	}
}

type Observer interface {
	Notify(data interface{}) //that something happened
}

type PropertyChange struct {
	Name  string // ex "Age", "Height"
	Value interface{}
}

type Person struct {
	Observable
	age int
} // Age() - getter, SetAge() - setter

func NewPerson(age int) *Person {
	return &Person{
		Observable: Observable{new(list.List)},
		age:        age,
	}
}

func (p *Person) Age() int { return p.age }
func (p *Person) SetAge(age int) {
	if age == p.age {
		return
	}

	oldCanVote := p.CanVote()

	p.age = age
	p.Fire(PropertyChange{"Age", p.age})

	if oldCanVote != p.CanVote() {
		p.Fire(PropertyChange{"CanVote", p.CanVote()})
	}

}

func (p *Person) CanVote() bool {
	return p.age >= 18
}

type ElectoralRoll struct {
}

func (e *ElectoralRoll) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Name == "CanVote" && pc.Value.(bool) {
			fmt.Println("Congradulations, You can vote!")
		}
	}
}

// Property doesn't scale..

func main() {
	p := NewPerson(10)
	er := &ElectoralRoll{}
	p.Subscribe(er)
	for i := 10; i < 20; i++ {
		fmt.Println("setting age to ", i)
		p.SetAge(i)
	}
}
