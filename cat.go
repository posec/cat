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
		var loop func([]string) struct{}
		loop = func(args []string) struct{} {
			if len(args) == 0 {
				return struct{}{}
			}
			f := args[0]
			in, err := os.Open(f)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%s\n", err)
				os.Exit(2)
			}
			defer in.Close()
			io.Copy(os.Stdout, in)
			return loop(args[1:])
		}
		loop(args)
	}
}
