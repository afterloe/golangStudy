package main

import "golang.org/x/tour/reader"

type MyReader struct {}

func (r MyReader) Read(buf []byte) (int, error) {
    buf[0] = 'A'
    return 1, nil
}

func main() {
    reader.Validate(MyReader{})
}
