package internal

import "fmt"

type Greeter interface {
	SayHello(name string, polite bool)
	Ouch() string
}

type greeter struct {
	salutation string
}

func (g *greeter) Ouch() string {
	return "Ouch!"
}

func (g *greeter) SayHello(name string, polite bool) {
	if polite {
		fmt.Printf("Good day, %s!\n", name)
	} else {
		fmt.Printf("%s, %s\n", g.salutation, name)
	}
}

func NewGreeterFactory(salutation string) func() Greeter {
	return func() Greeter {
		return &greeter{salutation: salutation}
	}
}
