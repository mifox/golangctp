package main

import (
	"bytes"
	"encoding/binary"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"reflect"

	"./simplezip"
)
import . "./ftdstruct"
import (
	"encoding/hex"
	"time"
	"unsafe"

	. "./ctpstruct"
)

func main2() {
	var x = rune('c')
	fmt.Print(x)
	var u CThostFtdcReqUserLoginField

	file, err := os.Open("gob.bin")
	if err != nil {
		fmt.Println(err)
	}
	dec := gob.NewDecoder(file)
	err2 := dec.Decode(&u)
	fmt.Printf("%+v", u)
	// var f := u[0].(type)
	fmt.Println(err2)
	var testStruct = &CThostFtdcReqUserLoginField{}
	//testStruct.A = 90
	copy(testStruct.BrokerID[:], []byte("9999"))
	copy(testStruct.UserID[:], []byte("089303"))
	copy(testStruct.Password[:], []byte("198759"))
	copy(testStruct.UserProductInfo[:], []byte(PROD_INFO))
	copy(testStruct.InterfaceProductInfo[:], []byte(API_INFO))
	copy(testStruct.ProtocolInfo[:], []byte(PROT_INFO))
	copy(testStruct.MacAddress[:], []byte("00-FF-9E-D9-18-EB"))
	fmt.Printf("%+v", *testStruct)
	var b byte = 'C'
	fmt.Printf("output: %v\n", reflect.TypeOf(b).Kind())
}
func main() {
	var loginPackage = &CThostFtdcReqUserLoginField{}
	copy(loginPackage.TradingDay[:], []byte("19990101"))
	copy(loginPackage.BrokerID[:], []byte("9999"))
	copy(loginPackage.UserID[:], []byte("089303"))
	copy(loginPackage.Password[:], []byte("198759"))
	copy(loginPackage.UserProductInfo[:], []byte(PROD_INFO))
	copy(loginPackage.InterfaceProductInfo[:], []byte(API_INFO))
	copy(loginPackage.ProtocolInfo[:], []byte(PROT_INFO))
	copy(loginPackage.MacAddress[:], []byte("00-FF-9E-D9-18-EB"))
	//bQuickPkt :=[]byte{0x10,1,0,6,0,1,0,0,0,0,0x10,1,0,6,0,2,0xff,0xff,0xff,0xff,0x10,1,
	//                    0,6,0,3,0xff,0xff,0xff,0xff,0x10,1,0,6,0,4,0,0,0,0}
	// 4         2         20         4              224
	//XMPHeader|CRPHeader|FTDCHeader|FTDCFieldInfo|CThostFtdcReqUserLoginField
	var i9 = unsafe.Sizeof(*loginPackage)
	var ftdcFieldInfo = &FTDCFieldInfo{}
	ftdcFieldInfo.Id = FTD_RID_LoginField
	ftdcFieldInfo.Size = uint16(i9)

	var ftdcHeader = &FTDCHeader{}
	ftdcHeader.Version = 8
	ftdcHeader.Chain = CHAIN_LAST
	ftdcHeader.Sid = 0
	ftdcHeader.Tid = FTD_TID_ReqUserLogin
	ftdcHeader.Sn = 0
	ftdcHeader.Field_count = 1
	ftdcHeader.Content_length = uint16(unsafe.Sizeof(*ftdcFieldInfo)) + ftdcFieldInfo.Size
	ftdcHeader.Reqid = 0

	var crpHeader = &CRPHeader{}
	crpHeader.Protocol = 1
	crpHeader.Method = CRPCM_NONE

	var xmpHeader = &XMPHeader{}
	xmpHeader.Protocol = FTDTypeFTDC
	xmpHeader.Extra_size = 0
	xmpHeader.Content_length = uint16(unsafe.Sizeof(*crpHeader)) + uint16(unsafe.Sizeof(*ftdcHeader)) + ftdcHeader.Content_length

	//ftdcFieldInfo.Id =
	//fmt.Printf("%d",i9)
	bLoginPkt := [0x128]byte{FTDTypeFTDC, 0, 0, 0xfa, 1, 0, 0x0c, CHAIN_LAST, 0, 1, 0, 0, 0x30, 0, 0, 0, 0, 0, 0, 1, 0, 0xe4, 0, 0, 0, 1, 0x10, 2, 0, 0xe0}
	//fmt.Printf("%x\n",bLoginPkt)
	//bLoginPkt[0x16] = 0
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, xmpHeader)
	copy(bLoginPkt[0:4], buf.Bytes())
	buf.Reset()
	binary.Write(buf, binary.BigEndian, crpHeader)
	copy(bLoginPkt[4:6], buf.Bytes())
	buf.Reset()
	binary.Write(buf, binary.BigEndian, ftdcHeader)
	copy(bLoginPkt[6:26], buf.Bytes())
	buf.Reset()

	binary.Write(buf, binary.BigEndian, ftdcFieldInfo)
	copy(bLoginPkt[26:30], buf.Bytes())
	buf.Reset()

	binary.Write(buf, binary.LittleEndian, loginPackage)
	copy(bLoginPkt[30:], buf.Bytes())
	buf.Reset()

	conn, err := net.Dial("tcp", "180.168.146.187:10001")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("获取180.168.146.187:10001的tcp连接成功...")
	go Handle(conn)
	defer conn.Close()
	conn.Write(bLoginPkt[:])
	hex.Dump(bLoginPkt[:])
	fmt.Println("begin")
	for {
		fmt.Printf("______")
		time.Sleep(time.Duration(1) * time.Second)
	}
	//fmt.Println("end")
}

//Handle 接收线程
func Handle(conn net.Conn) {

	for {
		//buffer = 0
		buffer := make([]byte, 4096)
		r1, _ := conn.Read(buffer)
		//fmt.Printf(" conn.Read(buffer) %d", r1)
		//fmt.Println(hex.Dump(buffer[0:r1]))
		Pyaser(buffer, r1)

	}

	//defer conn.Close()
}

//Pyaser 解析包
func Pyaser(buffer []byte, r1 int) {
	fmt.Printf("conn.Read %d   \n", r1)
	//fmt.Printf("%x\n", buffer[0:4])
	//fmt.Printf("%x\n", buffer[4:6])
	//fmt.Printf("%x\n", buffer[6:])
	// 4         2         20         4              224
	//XMPHeader|CRPHeader|FTDCHeader|FTDCFieldInfo|CThostFtdcReqUserLoginField
	for pos := 0; pos < r1; {
		var h XMPHeader
		binary.Read(bytes.NewBuffer(buffer[pos:pos+4]), binary.BigEndian, &h)

		pos += 4
		fmt.Printf("XMPHeader.Extra_size%d   ", h.Extra_size)
		fmt.Printf("XMPHeader.Content_length__%d\n", h.Content_length)

		var r CRPHeader
		binary.Read(bytes.NewBuffer(buffer[pos:pos+2]), binary.BigEndian, &r)
		var out2 bytes.Buffer
		if r.Method == CRPCM_ZERO {
			simplezip.DecodePacket(&out2, buffer[pos+2:pos+int(h.Content_length)])
			//fmt.Println((out2.Bytes()))
			//fmt.Println(string(out2.Bytes()))
		}
		//fmt.Printf("%x\n", out2)
		//fmt.Println(hex.Dump(buffer[pos:pos+int(h.Content_length)]))
		var ftdch FTDCHeader
		//fmt.Printf("%x\n", out2.Bytes()[0:20])
		binary.Read(bytes.NewBuffer(out2.Bytes()[0:20]), binary.BigEndian, &ftdch)
		//fmt.Printf(hex.Dump(out2.Bytes()[0:20]))
		//fmt.Printf("tid:   %x\n", ftdch.Tid)
		var instrumentStatus1 CThostFtdcInstrumentStatusField
		if ftdch.Tid == FTD_TID_IntlRtnDissemination {
			fmt.Printf("FTD_TID_IntlRtnDissemination tid:   %x\n", ftdch)
		} else if ftdch.Tid == FTD_TID_RspUserLogin {
			fmt.Printf("FTD_TID_RspUserLogin tid:   %x\n", ftdch)
		} else if ftdch.Tid == FTD_TID_RtnInstrumentStatus {
			fmt.Printf("FTD_TID_RtnInstrumentStatus tid:   %x\n", ftdch)

			//var fdi FTDCFieldInfo
			////fmt.Printf("%x\n", out2.Bytes()[20:24])
			//binary.Read(bytes.NewBuffer(out2.Bytes()[20:24]), binary.BigEndian, &fdi)
			fmt.Printf(hex.Dump(out2.Bytes()[24:]))
			//fmt.Printf("11111111%d\n",fdi.Size)//instrumentStatus.ExchangeInstID,instru

			binary.Read(bytes.NewBuffer(out2.Bytes()[24:]), binary.BigEndian, &instrumentStatus1)
			fmt.Print(instrumentStatus1) //instrumentStatus.ExchangeInstID,instrumentStatus.SettlementGroupID,
			//instrumentStatus.InstrumentID,instrumentStatus.InstrumentStatus,instrumentStatus.TradingSegmentSN,instrumentStatus.EnterTime,instrumentStatus.EnterReason

		}
		pos += int(h.Content_length)

	}

	//fmt.Println(hex.Dump(buffer[0:r1]))

	//fmt.Printf("ooooooooo__%d\n", r1)
}

//Pyaser2 临时
func Pyaser2(buffer []byte, r1 int) {
	var r CRPHeader
	binary.Read(bytes.NewBuffer(buffer[4:6]), binary.BigEndian, &r)
	//fmt.Printf("%x\n",r)
	fmt.Printf("%d %d\n", r.Protocol, r.Method)
	var out2 bytes.Buffer
	if r.Method == CRPCM_ZERO {
		fmt.Printf("%x\n", buffer[6:])
		simplezip.DecodePacket(&out2, buffer[6:r1])
		//fmt.Println((out2.Bytes()))
		//fmt.Println(string(out2.Bytes()))
	}
	fmt.Printf("%x\n", out2)

	var ftdch FTDCHeader
	fmt.Printf("%x\n", out2.Bytes()[0:20])
	binary.Read(bytes.NewBuffer(out2.Bytes()[0:20]), binary.BigEndian, &ftdch)
	fmt.Println(ftdch)
	fmt.Printf("%d\n", ftdch.Field_count)

	//out2.Bytes()[0:20]
	fmt.Printf("_________________%d\n", len(out2.Bytes()))
	fmt.Printf("%x\n", out2.Bytes()[0:])
	// 4         2         20         4              224
	//XMPHeader|CRPHeader|FTDCHeader|FTDCFieldInfo|CThostFtdcReqUserLoginField
	var fdi FTDCFieldInfo
	fmt.Printf("%x\n", out2.Bytes()[20:24])
	binary.Read(bytes.NewBuffer(out2.Bytes()[20:24]), binary.BigEndian, &fdi)
	fmt.Println(fdi)
	fmt.Printf("%d\n", fdi.Size)

	fmt.Printf("%x\n", out2.Bytes()[20:265])
	binary.Read(bytes.NewBuffer(out2.Bytes()[20+85+4:20+85+4+4]), binary.BigEndian, &fdi)
	fmt.Println(fdi)
	fmt.Printf("%d\n", fdi.Size)
	fmt.Printf("%x\n", out2.Bytes()[20+85+4+4:265])
	var response CThostFtdcRspUserLoginField
	binary.Read(bytes.NewBuffer(out2.Bytes()[20+85+4+4:]), binary.BigEndian, &response)
	fmt.Printf("%+v\n", response)
	file, err := os.Create("gob.bin")
	if err != nil {
		fmt.Println(err)
	}
	enc := gob.NewEncoder(file)
	err2 := enc.Encode(response)
	fmt.Println(err2)

	// close connection before exit

}
