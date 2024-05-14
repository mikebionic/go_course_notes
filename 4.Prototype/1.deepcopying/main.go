package main

import "fmt"

// This idea doesn't scale

type Address struct {
	StreetAddress, City, Country string
}

type Person struct {
	Name    string
	Address *Address
}

func main() {
	john := Person{"John", &Address{"123 London Rd", "London", "UK"}}
	jane := john
	// jane.Name = "Jane" //ok
	// jane.Address.StreetAddress = "321 Baker St"
	// fmt.Println(john, john.Address)
	// // {John 0xc000074090} &{321 Baker St London UK}
	// fmt.Println(jane, jane.Address)
	// // {Jane 0xc000074090} &{321 Baker St London UK}

	// copying the data, not only a pointer
	jane.Address = &Address{
		john.Address.StreetAddress,
		john.Address.City,
		john.Address.Country,
	}

	jane.Name = "Jane" //ok
	jane.Address.StreetAddress = "321 Baker St"
	fmt.Println(john, john.Address)
	// {John 0xc000074090} &{123 London Rd London UK}
	fmt.Println(jane, jane.Address)
	// {Jane 0xc000074090} &{321 Baker St London UK}
}
