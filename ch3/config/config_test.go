package config

import (
	"testing"
	"fmt"
)

type A struct {
	a int
}

func TestConfig(t *testing.T) {
	a1 := &A{1}
	a2 := &A{2}

	a3 := a1
	a1 = a2
	fmt.Println(a1.a, a2.a, a3.a)
}
