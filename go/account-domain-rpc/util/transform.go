package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/pkg/errors"
	"strconv"
)

func BytesToInt(b []byte) int {
	bytesBuffer := bytes.NewBuffer(b)
	var tmp int32
	binary.Read(bytesBuffer, binary.BigEndian, &tmp)
	return int(tmp)
}

//整形转换成字节
func IntToBytes(n int) []byte {
	tmp := int32(n)
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func Int64ToBytes(n int64) []byte {
	tmp := n
	bytesBuffer := bytes.NewBuffer([]byte{})
	binary.Write(bytesBuffer, binary.BigEndian, tmp)
	return bytesBuffer.Bytes()
}

func Md5StringtoByte(s string) ([]byte, error) {
	b := make([]byte, 0, 16)
	str := make([]string, 0)
	if len(s) != 32 {
		return nil, errors.New("this string is not Md5")
	}
	for i := 0; i < 32; i += 2 {
		str = append(str, s[i:i+1]+s[i+1:i+2])
	}
	for _, v := range str {
		if s, err := strconv.ParseUint(v, 16, 10); err == nil {
			b = append(b, byte(s))
		}
	}
	fmt.Println(b)
	if len(b) != 16 {
		return nil, errors.New("this string is not Md5")
	}
	return b, nil
}
