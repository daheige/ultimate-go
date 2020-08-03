package main

import (
	"crypto/md5"
	"fmt"
)

func main() {
	h := md5.New()
	h.Write([]byte("abc"))
	fmt.Printf("%x \n", h.Sum(nil))
}
