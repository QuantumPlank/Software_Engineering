package main

import (
	"fmt"
)

type IPerson interface {
	GetMaturity() string
	SetMaturity(string)
	GetName() string
	SetName(string)
}

type Person struct {
	maturity string
	name     string
}

func (p *Person) GetMaturity() string {
	return p.maturity
}

func (p *Person) SetMaturity(maturity string) {
	p.maturity = maturity
}

func (p *Person) GetName() string {
	return p.name
}

func (p *Person) SetName(name string) {
	p.name = name
}

type Adult struct {
	Person
}

func NewAdult() IPerson {
	return &Adult{
		Person: Person{
			maturity: "Adult",
			name:     "Chris",
		},
	}
}

type Kid struct {
	Person
}

func NewKid() IPerson {
	return &Kid{
		Person: Person{
			maturity: "Child",
			name:     "Orlando",
		},
	}
}

func main() {
	p1 := NewKid()
	p2 := NewAdult()

	fmt.Printf("P1 -> Maturity: %v, Name: %v\n", p1.GetMaturity(), p1.GetName())
	fmt.Printf("P2 -> Maturity: %v, Name: %v\n", p2.GetMaturity(), p2.GetName())
}
