package main

import (
	"bytes"
	"crypto/md5"
	"fmt"
)

func main() {
	//var hashChannel = make(chan []byte, 1)
	var buffer bytes.Buffer
	sum := md5.Sum(buffer.Bytes())
	//hashChannel <- sum[:]
	fmt.Println(sum)
}
