package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	counts := make(map[string]int)
	for _, filename := range os.Args[1:] {
		data, error := ioutil.ReadFile(filename)
		if error != nil {
			fmt.Fprintf(os.Stderr, "duplication_3: %v\n", error)
			continue
		}
		// fmt.Print(data)
		fmt.Print(strings.Split(string(data), "\n"))
		fmt.Print("\n")
		for _, line := range strings.Split(string(data), "\n") {
			counts[line]++
		}
	}
	for line, n := range counts {
		if n > 1 {
			fmt.Printf("%d\t%s\n", n, line)
		}
	}
}
