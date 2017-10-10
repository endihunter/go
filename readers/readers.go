package main

import (
	"io"
	"strings"
	"os"
	"fmt"
	"golang.org/x/tour/reader"
)

// MyReader
type MyReader struct{}

func (r MyReader) Read(buf []byte) (int, error) {
	for i, _ := range buf {
		buf[i] = 'A'
	}

	return len(buf), nil
}

// rot13Reader
type rot13Reader struct {
	reader io.Reader
}

func (rot *rot13Reader) Read(buf []byte) (n int, err error) {
	n, err = rot.reader.Read(buf)

	// Apply rot13 algorithm
	for i, v := range buf {
		switch {
		case v >= 'A' && v <= 'M' || v >= 'a' && v <= 'm':
			buf[i] += 13
		case v >= 'N' && v <= 'Z' || v >= 'n' && v <= 'z':
			buf[i] -= 13
		}
	}

	return
}

func main() {
	///////////////////////////
	r := strings.NewReader("Hello, World")
	b := make([]byte, 8)

	for {
		n, err := r.Read(b)

		fmt.Printf("n = %v, err = %v, b = %v\n", n, err, b)
		fmt.Printf("b[:n] = %q\n", b[:n])

		if err == io.EOF {
			break
		}
	}

	///////////////////////////
	fmt.Println(strings.Repeat("=", 40))
	reader.Validate(MyReader{})

	///////////////////////////
	fmt.Println(strings.Repeat("=", 40))
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	rdr := rot13Reader{s}
	io.Copy(os.Stdout, &rdr)
}
