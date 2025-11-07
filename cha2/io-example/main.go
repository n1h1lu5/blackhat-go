package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

// FooReader defines an io.Reader to read from stdin.

type FooReader struct{}

// Read reads data from stdin
func (fooReader *FooReader) Read(b []byte) (int, error) {
	fmt.Print("in> ")
	return os.Stdin.Read((b))
}

// FooWriter defined an io.Writer to wwrite  to stdout
type FooWriter struct{}

// Write witres data to stdout
func (FooWriter *FooWriter) Write(b []byte) (int, error) {
	fmt.Print("out> ")
	return os.Stdout.Write(b)
}

func main() {
	var (
		reader FooReader
		writer FooWriter
	)
	if _, err := io.Copy(&writer, &reader); err != nil {
		log.Fatalln("Unable to read/write data")
	}
}
