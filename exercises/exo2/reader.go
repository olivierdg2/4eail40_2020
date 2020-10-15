package main

import (
	"io"
    "strings"
    "os"
    "bytes"
)

type spaceEraser struct {
	r io.Reader
}


func (s spaceEraser) Read(p []byte) (int, error) {
    n, err := s.r.Read(p)
    //Remove all spaces by replacing them by nothing
    t := bytes.ReplaceAll(p, []byte(" "), []byte(""))
    //Some withe spaces can still be remaining
    t = bytes.Trim(t, "\x00")
    //Need to update the length of the byte
    n, err = s.r.Read(t)
	n = copy(p, t)
	return n, err
}
func main() {
    s := strings.NewReader("H e l l o w o r l d!")
    r := spaceEraser{s}
    io.Copy(os.Stdout, &r)
}
    