package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

type A struct {
	// should be exported member when read back from buffer
	One int32
	Two int32
}

// CTP传输协议报头
type XMPHeader struct {
	// 上层协议的协议ID
	Protocol uint8
	//  扩展报文长度
	Extra_size uint8
	// 包体长度
	Content_length uint16
}

var aa A
var h XMPHeader

func main() {
	buf := new(bytes.Buffer)
	aa.One = 1
	binary.Write(buf, binary.LittleEndian, aa)
	binary.Read(buf, binary.LittleEndian, &aa)
	fmt.Println("after aa is ", aa)

	buf1 := new(bytes.Buffer)
	h.Protocol = 1
	binary.Write(buf1, binary.LittleEndian, h)
	binary.Read(buf1, binary.LittleEndian, &h)
	fmt.Println("after aa is ", h)
}
