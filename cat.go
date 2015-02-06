package main

import (
	"flag"
	"fmt"
	"io"
	"os"
)

var uIgnored = flag.Bool("u", false, "unbuffered (ignored)")

func main() {
	var exitStatus int

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		io.Copy(os.Stdout, os.Stdin)
	} else {
		for _, f := range args {
			func() {
				var err error
				var in *os.File

				if f == "-" {
					// XCU7: If a file is '-', the cat utility shall read from the standard input at that point in the sequence
					in = os.Stdin
				} else {
					in, err = os.Open(f)
					if err != nil {
						fmt.Fprintf(os.Stderr, "%s\n", err)
						exitStatus = 2
						return
					}
					defer in.Close()
				}
				_, err = io.Copy(os.Stdout, in)
				if err != nil {
					fmt.Fprintf(os.Stderr, "%s\n", err)
					exitStatus = 2
					return
				}
			}()
		}
	}
	os.Exit(exitStatus)
}
