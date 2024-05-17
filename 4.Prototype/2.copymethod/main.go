package main

import "fmt"

type Address struct {
	StreetAddress, City, Country string
}

func (a *Address) DeepCopy() *Address {
	return &Address{
		a.StreetAddress,
		a.City,
		a.Country,
	}
}

type Person struct {
	Name    string
	Address *Address
	Friends []string
}

func (p *Person) DeepCopy() *Person {
	q := *p
	q.Address = p.Address.DeepCopy()
	copy(q.Friends, p.Friends)
	return &q
}

func main() {
	john := Person{"John",
		&Address{"123 London Rd", "London", "UK"},
		[]string{"Chris", "Chad"}}
	jane := john.DeepCopy()
	jane.Name = "Jane" //ok
	jane.Address.StreetAddress = "321 Baker St"
	jane.Friends = append(jane.Friends, "Matt")
	fmt.Println(john, john.Address)
	// {John 0xc000074090 [Chris Chad]} &{123 London Rd London UK}
	fmt.Println(jane, jane.Address)
	// &{Jane 0xc0000740f0 [Chris Chad Matt]} &{321 Baker St London UK}
}
