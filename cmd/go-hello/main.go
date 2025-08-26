package main

import (
	"github.com/guilhermehbueno/go-hello/internal/greet"
)

func main() {
	SayHello()
}

func SayHello() {
	greet.Greet()
}
