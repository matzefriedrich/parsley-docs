package services

import "fmt"

type Greeter interface {
	SayHello(name string, polite bool) string
}

type greeter struct {
}

func (g *greeter) SayHello(name string, polite bool) string {
	if polite {
		return fmt.Sprintf("Good day, %s!\n", name)
	} else {
		return fmt.Sprintf("Hi, %s\n", name)
	}
}

func NewGreeter() Greeter {
	return &greeter{}
}
