package common

import (
	"encoding/hex"
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type EUI64 [8]byte

func (e EUI64) String() string {
	var eui string = ""
	for i, v := range e[:] {
		eui += hex.EncodeToString([]byte{v})
		if i < len(e)-1 {
			eui += "-"
		}
	}
	return eui
}
func (e *EUI64) UnmarshalString(eui string) error {
	b := make([]byte, 0, 8)
	sp := strings.Split(eui, "-")
	if len(sp) != 8 {
		return errors.New("eui is not EUI64")
	}
	for _, v := range sp {
		if s, err := strconv.ParseUint(v, 16, 10); err == nil {
			b = append(b, byte(s))
		}
	}
	if len(e) != len(b) {
		return errors.New("eui is not EUI64")
	}
	for i, v := range b {
		e[i] = v
	}
	return nil
}
func (e EUI64) MarshalBinary() ([]byte, error) {
	out := make([]byte, len(e))
	for i, v := range e {
		out[i] = v
	}
	return out, nil
}
func (e *EUI64) UnmarshalBinary(data []byte) error {
	if len(data) != len(e) {
		return fmt.Errorf("common: %d bytes of data are expected", len(e))
	}
	for i, v := range data {
		e[i] = v
	}
	return nil
}
func (e EUI64) MarshalText() ([]byte, error) {
	return []byte(e.String()), nil
}
func (e *EUI64) UnmarshalText(text []byte) error {
	b, err := hex.DecodeString(string(text))
	if err != nil {
		return err
	}
	if len(e) != len(b) {
		return fmt.Errorf("lorawan: exactly %d bytes are expected", len(e))
	}
	copy(e[:], b)
	return nil
}
func (e *EUI64) Scan(src interface{}) error {
	b, ok := src.([]byte)
	if !ok {
		return errors.New("lorawan: []byte type expected")
	}
	if len(b) != len(e) {
		return fmt.Errorf("lorawan []byte must have length %d", len(e))
	}
	copy(e[:], b)
	return nil
}
