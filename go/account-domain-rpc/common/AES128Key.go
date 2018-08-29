package common

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type AES128Key [16]byte

const (
	AppEUINode = 1
	DevEUINode = 2
)

func (a AES128Key) String() string {
	var appkey string = ""
	for i, v := range a[:] {
		appkey += hex.EncodeToString([]byte{v})
		if i < len(a)-1 {
			appkey += "-"
		}
	}
	return appkey
}
func (a *AES128Key) UnmarshalString(appkey string) error {
	b := make([]byte, 0, 16)
	sp := strings.Split(appkey, "-")
	if len(sp) != 16 {
		return errors.New("appkey is not AES128Key")
	}
	for _, v := range sp {
		if s, err := strconv.ParseUint(v, 16, 10); err == nil {
			b = append(b, byte(s))
		}
	}
	if len(a) != len(b) {
		return errors.New("key is not 16 bit")
	}
	for i, v := range b {
		a[i] = v
	}
	return nil
}
func (a AES128Key) MarshalBinary() ([]byte, error) {
	out := make([]byte, len(a))
	for i, v := range a {
		out[i] = v
	}
	return out, nil
}
func (a *AES128Key) UnmarshalBinary(data []byte) error {
	if len(data) != len(a) {
		return fmt.Errorf("common: %d bytes of data are expected", len(a))
	}
	for i, v := range data {
		a[i] = v
	}
	return nil
}
func (a AES128Key) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}
func (a *AES128Key) UnmarshalText(text []byte) error {
	b, err := hex.DecodeString(string(text))
	if err != nil {
		return err
	}
	if len(b) != len(a) {
		return fmt.Errorf("plink: exactly %d bytes are expected", len(a))
	}
	copy(a[:], b)
	return nil
}
func (a *AES128Key) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("plink: []byte type expected")
	}
	if len(b) != len(a) {
		return fmt.Errorf("plink []byte must have length %d", len(a))
	}
	copy(a[:], b)
	return nil
}
