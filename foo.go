package goland_buildtags

import (
	"fmt"
)

type foo struct {}

func (f foo) Do() {
	fmt.Println("foo->Do()")
}

func (f foo) Greet(s string) {
	fmt.Println("hello", s)
}
