package main

import "fmt"

// Interface Segregation Principle

type Document struct {
}

type Machine interface {
	Print(d Document)
	Fax(d Document)
	Scan(d Document)
}

type MultiFunctionPrinter struct {
}

func (m MultiFunctionPrinter) Print(d Document) {

}
func (m MultiFunctionPrinter) Fax(d Document) {

}
func (m MultiFunctionPrinter) Scan(d Document) {

}

type OldFashionedPrinter struct {
}

func (o OldFashionedPrinter) Print(d Document) {
	//ok
}
func (o OldFashionedPrinter) Fax(d Document) {
	panic("operation not supported")
}

// Deprecated: ...
func (o OldFashionedPrinter) Scan(d Document) {
	panic("operation not supported")
}

// ISP
type Printer interface {
	Print(d Document)
}
type Scanner interface {
	Scan(d Document)
}

type MyPrinter struct{}

func (m MyPrinter) Print(d Document) {}

type Photocopier struct{}

func (p Photocopier) Print(d Document) {}
func (p Photocopier) Scan(d Document) {
	fmt.Println("scanned..")
}

type MultiFunctionDevice interface {
	Printer
	Scanner
	// Fax
}

// decorator
type MultiFuncitonMachine struct {
	printer Printer
	scanner Scanner
}

func (m MultiFuncitonMachine) Print(d Document) {
	m.printer.Print(d)
}
func (m MultiFuncitonMachine) Scan(d Document) {
	m.scanner.Scan(d)
}

func main() {
	// ofp := OldFashionedPrinter{}

	doc := Document{}
	pc := Photocopier{}
	pc.Scan(doc)
}
