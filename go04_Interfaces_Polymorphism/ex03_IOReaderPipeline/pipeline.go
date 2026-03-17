package main

import "io"

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
