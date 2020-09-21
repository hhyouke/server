package api

import (
	"fmt"
	"testing"
)

func TestPanicAndRecover(t *testing.T) {
	recoverF()
	panicF2()
}

// For recovery
func recoverF() {
	defer func() {
		if a := recover(); a != nil {
			fmt.Println("RECOVER", a)
		} else {
			fmt.Println("a is nil")
		}
	}()
	panicF1()
}

// panicF1
func panicF1() {
	fmt.Println("panic func")
	panic("Panicked!!")
}

func panicF2() {
	fmt.Println("panic func2")
	panic("Panicked2!!")
}
