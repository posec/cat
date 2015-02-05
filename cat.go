package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	for _, f := range os.Args[1:] {
		in, err := os.Open(f)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s: %s\n", in, err)
			os.Exit(2)
		}
		defer in.Close()
		io.Copy(os.Stdout, in)
	}
}
