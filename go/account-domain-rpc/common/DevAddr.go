package common

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type DevAddr [3]byte

func (a DevAddr) String() string {
	var appkey string = ""
	for i, v := range a[:] {
		appkey += hex.EncodeToString([]byte{v})
		if i < len(a)-1 {
			appkey += "-"
		}
	}
	return appkey
}
func (a *DevAddr) UnmarshalString(appkey string) error {
	b := make([]byte, 0, 3)
	for _, v := range strings.Split(appkey, "-") {
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
func (a DevAddr) MarshalBinary() ([]byte, error) {
	out := make([]byte, len(a))
	for i, v := range a {
		out[i] = v
	}
	return out, nil
}
func (a *DevAddr) UnmarshalBinary(data []byte) error {
	if len(data) != len(a) {
		return fmt.Errorf("common: %d bytes of data are expected", len(a))
	}
	for i, v := range data {
		a[i] = v
	}
	return nil
}
func (a DevAddr) MarshalText() ([]byte, error) {
	return []byte(a.String()), nil
}
func (a *DevAddr) UnmarshalText(text []byte) error {
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
func (a *DevAddr) Scan(src interface{}) error {
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

func (a *DevAddr) GetDevAddr() (string, error) {
	if _, err := rand.Read(a[:]); err != nil {
		return "", err
	}
	return a.String(), nil
}
