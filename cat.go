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
				var in *os.File
				if f == "-" {
					// XCU7: If a file is '-', the cat utility shall read from the standard input at that point in the sequence
					in = os.Stdin
				} else {
					var err error
					in, err = os.Open(f)
					if err != nil {
						fmt.Fprintf(os.Stderr, "%s\n", err)
						os.Exit(2)
					}
					defer in.Close()
				}
				io.Copy(os.Stdout, in)
			}()
		}
	}
}
