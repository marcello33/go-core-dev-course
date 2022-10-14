package main

import (
	"fmt"
	"io"
	"os"
	"strings"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

type rot13Reader struct {
	r io.Reader
}

// Read is a MyReader method to emits an infinite stream of ASCII character 'A'
func (r MyReader) Read(b []byte) (int, error) {
	i := 0
	for i < len(b) {
		b[i] = byte('A')
		i++
	}

	return i, nil
}

// Read is a rot13Reader that implements io.Reader and reads from it
// The method modifies the stream by applying the rot13 substitution cipher to all alphabetical characters
func (r *rot13Reader) Read(b []byte) (int, error) {
	n, err := r.r.Read(b)

	if err == nil {
		for i := 0; i < n; i++ {
			if b[i] > 'A' && b[i] < 'Z' {
				b[i] = 'A' + ((b[i] - 'A' + 13) % 26)
			} else if b[i] > 'a' && b[i] < 'z' {
				b[i] = 'a' + ((b[i] - 'a' + 13) % 26)
			}
		}
	}

	return n, err
}

func main() {
	// readers exercise 1
	fmt.Println("\nReaders exercise 1")
	reader.Validate(MyReader{})

	// readers exercise 2
	fmt.Println("\nReaders exercise 2")

	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}

	_, err := io.Copy(os.Stdout, &r)
	if err != nil {
		panic(err)
	}
}
