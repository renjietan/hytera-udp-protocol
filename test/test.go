package test

import (
	"fmt"
	"testing"
)

func TestMyFunction(t *testing.T) {
	a := map[string]int{}
	fmt.Println(a)
	if a["3"] > 0 {
		fmt.Println(a["1"])
	}
	a["3"] = 1
	if a["3"] > 0 {
		fmt.Println(a["2"])
	}
}
