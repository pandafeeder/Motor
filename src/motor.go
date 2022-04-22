package main

import . "motor/node"
import "fmt"

func main() {
	n := Node{
		Sourcefile: "abc",
		Inputs:     []string{"a", "b"},
		Outputs:    []string{"x", "y"},
	}
	fmt.Println(n)
}
