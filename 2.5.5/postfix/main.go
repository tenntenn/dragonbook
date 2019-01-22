package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	p := NewParser(os.Stdout, strings.NewReader("9-5+2"))
	if err := p.Parse(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Println()
}
