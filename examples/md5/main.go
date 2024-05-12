package main

import (
	"crypto/md5"
	"fmt"
)

func main() {

	MD5()

}

func MD5() {
	data := []byte("hello")
	fmt.Printf("%x\n", md5.Sum(data))
}
