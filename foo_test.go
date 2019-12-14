// +build !integration

package goland_buildtags

import (
	"testing"
)

func TestFoo_Do(t *testing.T) {
	f := new(foo)
	f.Do()
}

func TestFoo_Greet(t *testing.T) {
	f := new(foo)
	f.Greet("Goland")
}
