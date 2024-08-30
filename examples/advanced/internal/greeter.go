package internal

import "fmt"

//go:generate parsley-cli generate proxy

type Greeter interface {
	SayHello(name string, polite bool)
}

type greeter struct {
}

func (g *greeter) SayHello(name string, polite bool) {
	if polite {
		fmt.Printf("Good day, %s!\n", name)
	} else {
		fmt.Printf("Hello %s\n", name)
	}
}

func NewGreeter() Greeter {
	return &greeter{}
}
