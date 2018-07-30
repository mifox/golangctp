package main

import (
	"fmt"
	//"unsafe"
	"reflect"
)

type c_s6 [6]byte
type TestStructTobytes struct {
	data int8
	a    c_s6
}
type SliceMock struct {
	addr uintptr
	len  int
	cap  int
}
type TThostFtdcDateType [9]byte
type TThostFtdcBrokerIDType [11]byte
type TThostFtdcUserIDType [16]byte
type TThostFtdcPasswordType [41]byte
type TThostFtdcProductInfoType [11]byte
type TThostFtdcProtocolInfoType [11]byte
type TThostFtdcMacAddressType [21]byte
type TThostFtdcIPAddressType [16]byte
type TThostFtdcLoginRemarkType [36]byte

type CThostFtdcReqUserLoginField struct {
	///交易日
	TradingDay TThostFtdcDateType
	///经纪公司代码
	BrokerID TThostFtdcBrokerIDType
	///用户代码
	UserID TThostFtdcUserIDType
	///密码
	Password TThostFtdcPasswordType
	///用户端产品信息
	UserProductInfo TThostFtdcProductInfoType
	///接口端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType
	///协议信息
	ProtocolInfo TThostFtdcProtocolInfoType
	///Mac地址
	MacAddress TThostFtdcMacAddressType
	///动态密码
	OneTimePassword TThostFtdcPasswordType
	///终端IP地址
	ClientIPAddress TThostFtdcIPAddressType
	///登录备注
	LoginRemark TThostFtdcLoginRemarkType
}

func DeepFields(ifaceType reflect.Type) []reflect.StructField {
	var fields []reflect.StructField

	for i := 0; i < ifaceType.NumField(); i++ {
		v := ifaceType.Field(i)
		if v.Anonymous && v.Type.Kind() == reflect.Struct {
			fields = append(fields, DeepFields(v.Type)...)
		} else {
			fields = append(fields, v)
		}
	}

	return fields
}

func StructCopy(DstStructPtr interface{}, SrcStructPtr interface{}) {
	srcv := reflect.ValueOf(SrcStructPtr)
	dstv := reflect.ValueOf(DstStructPtr)
	srct := reflect.TypeOf(SrcStructPtr)
	dstt := reflect.TypeOf(DstStructPtr)
	if srct.Kind() != reflect.Ptr || dstt.Kind() != reflect.Ptr ||
		srct.Elem().Kind() == reflect.Ptr || dstt.Elem().Kind() == reflect.Ptr {
		panic("Fatal error:type of parameters must be Ptr of value")
	}
	if srcv.IsNil() || dstv.IsNil() {
		panic("Fatal error:value of parameters should not be nil")
	}
	srcV := srcv.Elem()
	dstV := dstv.Elem()
	srcfields := DeepFields(reflect.ValueOf(SrcStructPtr).Elem().Type())
	for _, v := range srcfields {
		if v.Anonymous {
			continue
		}
		dst := dstV.FieldByName(v.Name)
		src := srcV.FieldByName(v.Name)
		if !dst.IsValid() {
			continue
		}
		if src.Type() == dst.Type() && dst.CanSet() {
			dst.Set(src)
			continue
		}
		if src.Kind() == reflect.Ptr && !src.IsNil() && src.Type().Elem() == dst.Type() {
			dst.Set(src.Elem())
			continue
		}
		if dst.Kind() == reflect.Ptr && dst.Type().Elem() == src.Type() {
			dst.Set(reflect.New(src.Type()))
			dst.Elem().Set(src)
			continue
		}
	}
	return
}

func main() {
	slice1 := []byte{1, 2, 3, 4, 5}
	slice2 := []byte{5, 4, 3}
	var testStruct = &CThostFtdcReqUserLoginField{}
	//str2 := "99999"
	//data2 := [11]byte{'a','a','a','a','a','a','a'}
	//cc:=[]byte(str2)
	//data2 := [11]byte{"cccc"}
	//testStruct.BrokerID
	//copy(testStruct.BrokerID, data2)
	str2 := "99999"
	//cc:=[]byte(str2)
	//fmt.Printf("%x\n",cc)
	//
	//var a1=TThostFtdcBrokerIDType{'a','a'}
	//testStruct.BrokerID = a1
	copy(testStruct.BrokerID[:], []byte(str2))
	fmt.Printf("%x\n", testStruct.BrokerID)

	// StructCopy(testStruct.BrokerID, data2)
	//testStruct.BrokerID = cc
	fmt.Println(testStruct.BrokerID)
	fmt.Printf("%x\n", testStruct.BrokerID)

	copy(slice2, slice1) // 只会复制slice1的前3个元素到slice2中
	fmt.Printf("%x\n", slice1)
	fmt.Printf("%x\n", slice2)
	//var testStruct = &CThostFtdcReqUserLoginField{}
	//var x= [11]byte{"999999999"}
	//copy(x , []byte("9999999999"))
	//Len := unsafe.Sizeof(*testStruct)
	//testBytes := &SliceMock{
	//    addr: uintptr(unsafe.Pointer(testStruct)),
	//    cap:  int(Len),
	//    len:  int(Len),
	//}
	//data := *(*[]byte)(unsafe.Pointer(testBytes))
	//fmt.Println("[]byte is : ", data)
	//var testStruct = &TestStructTobytes{}
	//Len := unsafe.Sizeof(*testStruct)
	//testBytes := &SliceMock{
	//    addr: uintptr(unsafe.Pointer(testStruct)),
	//    cap:  int(Len),
	//    len:  int(Len),
	//}
	//data := *(*[]byte)(unsafe.Pointer(testBytes))
	//fmt.Println("[]byte is : ", data)
}
