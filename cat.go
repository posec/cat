package main

import (
"io"
"os"
)

func main() {
    io.Copy(os.Stdin, os.Stdout)
}
