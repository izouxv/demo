package common

import (
	"encoding/hex"
	"fmt"
)

type Topic struct{}

func (t *Topic) GetEUI64InTopic(DevEUI string) ([]byte, error) {
	b := make([]byte, 0)
	v := ""
	fmt.Println("devEUI:", DevEUI)
	for i := 0; i < len(DevEUI); i = i + 2 {
		v = DevEUI[i : i+2]
		b = append(b, byte(parseHex(v)))
	}
	fmt.Println("devEUI:", b)
	return b, nil
}

func (t *Topic) GetTopicInEUI64(devEUI []byte) string {
	str := ""
	for _, b := range devEUI {
		str += toHex(int32(b))
	}
	return str
}

func toHex(ten int32) (str string) {
	var m int32 = 0
	hex := make([]int32, 0, 2)
	for {
		m = ten / 16
		ten = ten % 16

		if m == 0 {
			hex = append(hex, ten)
			break
		}

		hex = append(hex, m)
	}
	if len(hex) == 1 {
		str = "0"
	}
	for i := 0; i < len(hex); i++ {
		if hex[i] >= 10 {
			str += fmt.Sprintf("%c", 'a'+hex[i]-10)
		} else {
			str += fmt.Sprint(hex[i])
		}
	}
	return
}

func parseHex(str string) (ten int32) {
	ten = 0
	h := str[0]
	l := str[1]
	if h >= 48 && h <= 57 {
		ten = (int32(h) - 48) * 16
	} else {
		ten = (int32(h) - 87) * 16
	}
	if l >= 48 && l <= 57 {
		ten = ten + int32(l) - 48
	} else {
		ten = ten + int32(l) - 87
	}
	return
}

func ShowBytes(e []byte) string {
	str := ""
	for i, v := range e[:] {
		str += hex.EncodeToString([]byte{v})
		if i < len(e)-1 {
			str += "-"
		}
	}
	return str
}
func ShowAclUsernameByDevEUI(devEUI []byte) string {
	str := ""
	for _, v := range devEUI[:] {
		str += hex.EncodeToString([]byte{v})
	}
	return str
}
func ShowAclPasswordByDevEUI(devEUI []byte) string {
	str := ""
	for _, v := range devEUI[:] {
		str += hex.EncodeToString([]byte{v})
	}
	return str
}
func ShowAclTopicByDevEUI(devEUI []byte) string {
	str := ""
	for _, v := range devEUI[:] {
		str += hex.EncodeToString([]byte{v})
	}
	return str
}
