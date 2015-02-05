package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for _, f := range args {
			func() {
				in, err := os.Open(f)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					os.Exit(2)
				}
				defer in.Close()
				io.Copy(os.Stdout, in)
			}()
		}
	}
}
