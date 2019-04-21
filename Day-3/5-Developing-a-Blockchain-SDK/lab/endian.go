package main

import (
	"encoding/binary"
	"log"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

func BigEndian() { // 大端序
	// 二进制形式：0000 0000 0000 0000 0001 0002 0003 0004
	var testInt int32 = 0x01020304 // 十六进制表示
	log.Printf("%d use big endian: \n", testInt)

	var testBytes []byte = make([]byte, 4)
	binary.BigEndian.PutUint32(testBytes, uint32(testInt)) //大端序模式
	log.Println("int32 to bytes:", testBytes)
	log.Printf("int32 to bytes: 0x%X \n", testBytes)
	log.Printf("int32 to bytes: %04b \n", testBytes)

	convInt := binary.BigEndian.Uint32(testBytes) //大端序模式的字节转为int32
	log.Printf("bytes to int32: %d\n\n", convInt)
}

func LittleEndian() { // 小端序
	//二进制形式： 0000 0000 0000 0000 0001 0002 0003 0004
	var testInt int32 = 0x01020304 // 16进制
	log.Printf("%d use little endian: \n", testInt)

	var testBytes []byte = make([]byte, 4)
	binary.LittleEndian.PutUint32(testBytes, uint32(testInt)) //小端序模式
	log.Println("int32 to bytes:", testBytes)
	log.Printf("int32 to bytes: 0x%X \n", testBytes)

	convInt := binary.LittleEndian.Uint32(testBytes) //小端序模式的字节转换
	log.Printf("bytes to int32: %d\n\n", convInt)
	log.Printf("int32 to bytes: %04b \n", testBytes)
}

func main() {
	BigEndian()
	LittleEndian()
}
