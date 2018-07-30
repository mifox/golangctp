package simplezip

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

func EncodePacket(w io.Writer, payload []byte) error {
	var iteratorIn = 0
	var iteratorOut = 0
	var inputNumBytes = len(payload)
	var outputBuffer [2000]byte
	for {
		if iteratorIn >= inputNumBytes {
			break
		}
		if (payload[iteratorIn] > 0xDF) && (payload[iteratorIn] < 0xF0) {

			outputBuffer[iteratorOut] = 0xE0
			iteratorOut = iteratorOut + 1
			outputBuffer[iteratorOut] = payload[iteratorIn]
			iteratorOut = iteratorOut + 1
			iteratorIn++
			continue
		}
		if 0x00 == payload[iteratorIn] {
			var step = iteratorIn + 1
			for {

				if step < inputNumBytes && (0x00 == payload[step]) && ((step - iteratorIn) < 15) {
					step = step + 1
				} else {
					break
				}
			}
			var tempChar uint8
			tempChar = uint8(step - iteratorIn)
			tempChar |= 0xE0
			outputBuffer[iteratorOut] = tempChar
			iteratorOut = iteratorOut + 1
			iteratorIn = step
			continue
		}
		//剩下的字符拷贝即可

		outputBuffer[iteratorOut] = payload[iteratorIn]
		iteratorOut = iteratorOut + 1
		iteratorIn = iteratorIn + 1
	}
	w.Write(outputBuffer[0:iteratorOut])
	return errors.New("EOF")
}

func DecodePacket(w io.Writer, payload []byte) error {
	var iteratorIn = 0
	var iteratorOut = 0
	var inputNumBytes = len(payload)
	var outputBuffer [2000]byte
	for {
		if iteratorIn >= inputNumBytes {
			break
		}

		if payload[iteratorIn] == 0xE0 {
			iteratorOut = iteratorOut + 1
			outputBuffer[iteratorOut] = payload[iteratorIn]
			iteratorIn = iteratorIn + 2
			continue
		}
		if (payload[iteratorIn] > 0xE0) && (payload[iteratorIn] < 0xF0) {
			var tempChar uint8
			tempChar = (uint8)(payload[iteratorIn] & 0x0F)
			//fmt.Println(tempChar)
			for i := tempChar; i > 0; i-- {
				//fmt.Println(tempChar)
				outputBuffer[iteratorOut] = 0x00
				iteratorOut = iteratorOut + 1
			}
			iteratorIn++
			continue
		}
		//剩下的字符拷贝即可

		outputBuffer[iteratorOut] = payload[iteratorIn]
		iteratorOut = iteratorOut + 1
		iteratorIn = iteratorIn + 1
	}
	w.Write(outputBuffer[0:iteratorOut])
	return errors.New("EOF")
}
func main() {
	var out bytes.Buffer
	bLoginPkt := [0x126]byte{0, 0, 1, 0x22, 1, 0, 8, 0, 0, 0, 0, 0, 0x30, 0, 0, 0, 0, 0, 0, 5, 1, 0xc, 0, 0, 0, 0, 0x10, 2, 0, 0xe0}
	fmt.Println(bLoginPkt)
	EncodePacket(&out, bLoginPkt[:])
	fmt.Println(out)
	var out2 bytes.Buffer
	DecodePacket(&out2, out.Bytes())
	fmt.Println(out2)

}
