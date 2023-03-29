// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 4.
//!+

// Echo1 prints its command-line arguments.
package main

import (
	"fmt"
	"os"
)

func main() {
	var s, sep string
	for i, arg := range os.Args[1:] {
		fmt.Println(i, arg)
		s += sep + arg
		sep = "*"
	}
	fmt.Println(s)
}

//!-
