package main

import (
	"io"
	"os"
	"fmt"
	"crypto/sha1"
)

func main() {
	args := os.Args;
	h := sha1.New()
	io.WriteString(h, args[1])
    fmt.Printf("%x\n", h.Sum(nil))
}
