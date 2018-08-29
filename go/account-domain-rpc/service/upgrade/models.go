package upgrade

import (
	"encoding/base64"
	"errors"
)

type UpgradeModele struct {
	Type byte
	Len  byte
	Data UpgradeDataModel
	CRC  byte
}

type UpgradeDataModel struct {
	DevEUI      [8]byte
	MainVersion [3]byte
	ComVersion  [3]byte
}

func (um UpgradeModele) String() (string, error) {
	input, err := um.MarshalBinary()
	if err != nil {
		return "", err
	}
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString, nil
}

func (um UpgradeModele) MarshalBinary() ([]byte, error) {
	out := make([]byte, 0)
	out = append(out, um.Type)
	out = append(out, um.Len)
	out = append(out, um.Data.DevEUI[:]...)
	out = append(out, um.Data.MainVersion[:]...)
	out = append(out, um.Data.ComVersion[:]...)
	out = append(out, um.CRC)
	return out, nil
}

func (um *UpgradeModele) UnmarshalBinary(in []byte) error {
	if len(in) != 17 {
		return errors.New("input in not 17 bytes")
	}
	um.Type = in[0]
	um.Len = in[1]
	for i, v := range in[2:10] {
		um.Data.DevEUI[i] = v
	}
	for i, v := range in[10:13] {
		um.Data.MainVersion[i] = v
	}
	for i, v := range in[13:16] {
		um.Data.ComVersion[i] = v
	}
	um.CRC = in[16]
	return nil
}

func (um *UpgradeModele) SetCRC() error {
	var cic int32 = 0
	data, err := um.MarshalBinary()
	if err != nil {
		return err
	}
	for _, v := range data[:len(data)-1] {
		cic += int32(v)
	}
	um.CRC = byte(cic & 0x00000000000000ff)
	return nil
}

type UpgradeResultModel struct {
	Type byte
	Len  byte
	Data UpgradeResultDataModel
	CRC  byte
}

type UpgradeResultDataModel struct {
	UpgradeType byte
	URLLen      byte
	URL         []byte
	Version     [3]byte
	MD5         [16]byte
}

func (urm UpgradeResultModel) String() (string, error) {
	input, err := urm.MarshalBinary()
	if err != nil {
		return "", err
	}
	encodeString := base64.StdEncoding.EncodeToString(input)
	return encodeString, nil
}

func (urm UpgradeResultModel) MarshalBinary() ([]byte, error) {
	out := make([]byte, 0)
	out = append(out, urm.Type)
	out = append(out, urm.Len)
	out = append(out, urm.Data.UpgradeType)
	out = append(out, urm.Data.URLLen)
	out = append(out, urm.Data.URL...)
	out = append(out, urm.Data.Version[:]...)
	out = append(out, urm.Data.MD5[:]...)
	out = append(out, urm.CRC)
	return out, nil
}

func (urm *UpgradeResultModel) UnmarshalBinary(in []byte) error {
	urm.Type = in[0]
	urm.Len = in[1]
	urm.Data.UpgradeType = in[2]
	urm.Data.URLLen = in[3]
	for _, v := range in[4 : in[3]+4] {
		urm.Data.URL = append(urm.Data.URL, v)
	}
	for i, v := range in[in[3]+4 : in[3]+7] {
		urm.Data.Version[i] = v
	}
	for i, v := range in[in[3]+7 : in[3]+17] {
		urm.Data.MD5[i] = v
	}
	urm.CRC = in[in[3]+17]
	return nil
}

func (urm *UpgradeResultModel) SetLen() {
	var length byte
	length = 21
	length += urm.Data.URLLen
	urm.Len = length
}

func (urm *UpgradeResultModel) SetCRC() error {
	var cic int32 = 0
	data, err := urm.MarshalBinary()
	if err != nil {
		return err
	}
	for _, v := range data[:len(data)-2] {
		cic += int32(v)
	}
	urm.CRC = byte(cic & 0x00000000000000ff)
	return nil
}
