package main

import (
	"io"
	"unicode"
)

type UpperReader struct {
	inner io.Reader
}

func (u *UpperReader) Read(p []byte) (n int, err error) {
	n, err = u.inner.Read(p)

	for i := 0; i < n; i++ {
		b := p[i]
		if b >= 'a' && b <= 'z' {
			p[i] = b - 32
		}
	}
	return
}

type ROT13Reader struct {
	inner io.Reader
}

func rot13byte(b byte) byte {
	switch {
	case b >= 'A' && b <= 'Z':
		return 'A' + (b-'A'+13)%26
	case b >= 'a' && b <= 'z':
		return 'z' + (b-'z'+13)%26
	default:
		return b
	}
}

func (r ROT13Reader) Read(p []byte) (n int, err error) {
	n, err = r.inner.Read(p)
	for i := 0; i < n; i++ {
		p[i] = rot13byte(p[i])
	}
	return
}

type WordCountReader struct {
	inner  io.Reader
	Count  int
	inWord bool
}

func (w WordCountReader) Read(p []byte) (n int, err error) {
	n, err = w.inner.Read(p)
	for i := 0; i < n; i++ {
		isSpace := unicode.IsSpace(rune(p[i]))
		if !isSpace && !w.inWord {
			w.Count++
		}
		w.inWord = !isSpace
	}
	return
}
