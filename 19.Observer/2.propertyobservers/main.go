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
	p.age = age
	p.Fire(PropertyChange{"Age", p.age})
}

type TrafficManagement struct {
	o Observable
}

func (t *TrafficManagement) Notify(data interface{}) {
	if pc, ok := data.(PropertyChange); ok {
		if pc.Value.(int) >= 18 {
			fmt.Println("Congrats, you can drive now!")
			t.o.Unsubscribe(t)
		}
	}
}

func main() {
	p := NewPerson(15)
	t := &TrafficManagement{p.Observable}
	p.Subscribe(t)
	for i := 16; i <= 20; i++ {
		fmt.Println("setting the age to", i)
		p.SetAge(i)
	}
}
