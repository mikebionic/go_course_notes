package main

type Person struct {
	Name     string
	Age      int
	EyeCount int
}

func NewPerson(name string, age int) *Person {
	if age < 16 {
		// ...
	}
	return &Person{name, age, 2}
}

func main() {
	// p := Person{"John", 22}
	p := NewPerson("John", 33)
	p.EyeCount = 1
}
