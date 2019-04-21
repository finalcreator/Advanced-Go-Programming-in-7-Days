package main

import (
	"bufio"
	"bytes"
	"encoding/binary"
	"io"
	"log"
	"os"
	"time"
	"unsafe"
)

func init() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

type Package struct {
	Version        [2]byte // 协议版本
	Length         int16   // 数据部分长度, 数据部分占用多少个字节. (除了Version, Length, 其余都是数据部分)
	Timestamp      int64   // 时间戳
	HostnameLength int16   // 主机名长度
	Hostname       []byte  // 主机名
	TagLength      int16   // Tag长度
	Tag            []byte  // Tag
	Msg            []byte  // 数据部分长度
}

func (p *Package) Pack(writer io.Writer) error {
	var err error
	err = binary.Write(writer, binary.BigEndian, &p.Version)
	err = binary.Write(writer, binary.BigEndian, &p.Length)
	err = binary.Write(writer, binary.BigEndian, &p.Timestamp)
	err = binary.Write(writer, binary.BigEndian, &p.HostnameLength)
	err = binary.Write(writer, binary.BigEndian, &p.Hostname)
	err = binary.Write(writer, binary.BigEndian, &p.TagLength)
	err = binary.Write(writer, binary.BigEndian, &p.Tag)
	err = binary.Write(writer, binary.BigEndian, &p.Msg)
	return err
}
func (p *Package) Unpack(reader io.Reader) error {
	var err error
	err = binary.Read(reader, binary.BigEndian, &p.Version)
	err = binary.Read(reader, binary.BigEndian, &p.Length)
	err = binary.Read(reader, binary.BigEndian, &p.Timestamp)
	err = binary.Read(reader, binary.BigEndian, &p.HostnameLength)
	p.Hostname = make([]byte, p.HostnameLength)
	err = binary.Read(reader, binary.BigEndian, &p.Hostname)
	err = binary.Read(reader, binary.BigEndian, &p.TagLength)
	p.Tag = make([]byte, p.TagLength)
	err = binary.Read(reader, binary.BigEndian, &p.Tag)
	p.Msg = make([]byte, p.Length-8-2-p.HostnameLength-2-p.TagLength)
	err = binary.Read(reader, binary.BigEndian, &p.Msg)
	return err
}

func (p *Package) String() string {
	return log.Sprintf("version:%s length:%d timestamp:%d hostname:%s tag:%s msg:%s",
		p.Version,
		p.Length,
		p.Timestamp,
		p.Hostname,
		p.Tag,
		p.Msg,
	)
}

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		log.Fatal(err)
	}

	pack := &Package{
		Version:        [2]byte{'V', '1'},
		Timestamp:      time.Now().Unix(),
		HostnameLength: int16(len(hostname)),
		Hostname:       []byte(hostname),
		TagLength:      4,
		Tag:            []byte("demo"),
		//Msg:            []byte(("现在时间是:" + time.Now().Format("2006-01-02 15:04:05"))),
	}

	pack.Msg = []byte("Z")

	//sizeofNow := unsafe.Sizeof(time.Now().Unix())
	//log.Println(sizeofNow) // 8

	sizeOfHostnameLength := unsafe.Sizeof(pack.HostnameLength)
	log.Println(sizeOfHostnameLength) // 2

	sizeOfTagLength := unsafe.Sizeof(pack.TagLength)
	log.Println(sizeOfTagLength) // 2

	pack.Length = 8 + 2 + pack.HostnameLength + 2 + pack.TagLength + int16(len(pack.Msg))

	buf := new(bytes.Buffer)
	// 写入四次，模拟TCP粘包效果
	pack.Pack(buf)
	pack.Pack(buf)
	pack.Pack(buf)
	pack.Pack(buf)
	// scanner
	scanner := bufio.NewScanner(buf)
	scanner.Split(func(data []byte, atEOF bool) (advance int, token []byte, err error) {
		log.Printf("data -> %v", data)
		if !atEOF && data[0] == 'V' {
			if len(data) > 4 {
				length := int16(0)
				binary.Read(bytes.NewReader(data[2:4]), binary.BigEndian, &length)
				if int(length)+4 <= len(data) {
					return int(length) + 4, data[:int(length)+4], nil
				}
			}
		}
		return
	})

	scanCount := 0
	for scanner.Scan() {
		scanCount++
		scannedPack := new(Package)
		scannedPack.Unpack(bytes.NewReader(scanner.Bytes()))
		log.Printf("Count %d: %v \n", scanCount, scannedPack)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("无效数据包")
	}

	now1 := time.Now().Unix()
	log.Printf("%d", now1)
	bs := make([]byte, 8)
	binary.BigEndian.PutUint64(bs, uint64(now1))
	log.Println(bs)

	s := string([]byte{0, 0, 0, 0, 92, 188, 27, 241, 0, 9, 76, 69, 79, 78, 73, 68, 45, 80, 67, 0, 4, 100, 101, 109, 111, 90})
	log.Println(s) // ABC€
	s2 := string([]byte{86, 49, 0, 26})
	log.Println(s2) // ABC€
}
