package main

import (
    "io"
    "os"
    "strings"
)

type rot13Reader struct {
    r io.Reader
}

func rot13(b byte) byte {
    switch {
    case 'A' <= b && b <='M':
        b = b + 13
    case 'M' < b && b <= 'Z':
        b = b - 13
    case 'a' <= b && b <= 'm':
        b = b + 13
    case 'm' < b && b <= 'z':
        b = b - 13
    }
    return b;
}

func (r rot13Reader) Read (buf []byte) (int, error) {
    n, e := r.r.Read(buf)
    for i := 0; i < n; i++ {
        buf[i] = rot13(buf[i])
    }
    return n, e
}

func main(){
    s := strings.NewReader("Lbh penpxrq gur pbqr!")
    r := rot13Reader{s}
    io.Copy(os.Stdout, &r)
}

