package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var uIgnored = flag.Bool("u", false, "unbuffered (ignored)")

func main() {
	flag.Parse()
	args := flag.Args()
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
