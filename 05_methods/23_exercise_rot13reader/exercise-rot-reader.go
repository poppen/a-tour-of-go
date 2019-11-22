package main

import (
	"io"
	"os"
	"strings"
)

type rot13Reader struct {
	r io.Reader
}

func (s rot13Reader) Read(b []byte) (int, error) {
	n, err := s.r.Read(b)
	if err != nil {
		return n, err
	}

	for i, c := range b[:n] {
		switch {
		case 65 <= c && c <= 77:
			b[i] = c + 13
		case 78 <= c && c <= 90:
			b[i] = c - 13
		case 97 <= c && c <= 109:
			b[i] = c + 13
		case 110 <= c && c <= 123:
			b[i] = c - 13
		}
	}
	return len(b), nil
}

func main() {
	s := strings.NewReader("Lbh penpxrq gur pbqr!")
	r := rot13Reader{s}
	io.Copy(os.Stdout, &r)
}
