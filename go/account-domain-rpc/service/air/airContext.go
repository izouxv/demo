package air

import (
	"account-domain-rpc/common"
	"account-domain-rpc/link"
	"account-domain-rpc/storage"
	"crypto/aes"
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/jacobsa/crypto/cmac"
)

const (
	JoinRequest        byte = 0x41
	JoinAccept         byte = 0x41
	BindRequest        byte = 0x42
	BindResponse       byte = 0x42
	GetFwareInfo       byte = 0xF7
	ReportFwarInfo     byte = 0xF7
	UpgradeOrder       byte = 0xF8
	UpgradeOrderAck    byte = 0xF8
	CModelUpgradeAck   byte = 0xFB
	BModelUpgradeAck   byte = 0xFE
	Type_GET           byte = 0x01
	Type_BOOT          byte = 0x02
	Type_Sleep         byte = 0x03
	Type_WindSpeed     byte = 0x04
	Type_TimedShutdown byte = 0x05
	Type_Anion         byte = 0x06
	Type_Automatic     byte = 0x07
	Type_Manual        byte = 0x08
	Type_UVLamp        byte = 0x09
	Type_SuperAir      byte = 0x0a
)

type AirContext struct {
	dataPayload link.DataPayload
}

func (a *AirContext) NewDataContext(MType byte, data []byte) error {
	var err error
	switch MType {
	case JoinRequest:
		a.dataPayload = new(PHYPayload)
	case BindRequest:
		a.dataPayload = new(BindPayload)
	case ReportFwarInfo:
		a.dataPayload = new(ReportPayload)
	case UpgradeOrderAck:
		a.dataPayload = new(UOAPayload)
	case CModelUpgradeAck:
		a.dataPayload = new(CMUAPayload)
	case BModelUpgradeAck:
		a.dataPayload = new(BMUAPayload)
	case Type_GET, Type_BOOT, Type_Sleep, Type_WindSpeed, Type_TimedShutdown, Type_Anion, Type_Automatic, Type_Manual, Type_UVLamp, Type_SuperAir:
		a.dataPayload = new(OrderRespPayload)
	default:
		return errors.New("no type")
	}
	if err := a.dataPayload.UnmarshalBinary(data); err != nil {
		return err
	}
	return err
}

func (a *AirContext) HandlerDataPayload(node storage.Node) ([]byte, error) {
	dataPayload := a.dataPayload
	fmt.Println("dataPayload", dataPayload)
	out, err := dataPayload.HandlerDataPayload(node)
	if err != nil {
		return out, err
	}
	return out, nil
}

func ValidateCRC(data []byte, CRC byte) bool {
	crc := CalculateCRC(data)
	if crc != CRC {
		log.Infof("calculate crc %x  src CRC %x", crc, CRC)
		return false
	}
	return true
}

func CalculateCRC(data []byte) byte {
	var cic int32 = 0
	for _, v := range data {
		cic += int32(v)
	}
	return byte(cic & 0x00000000000000ff)
}

func ValidateMIC(key common.AES128Key, data []byte) (bool, error) {
	var mic []byte
	MIC := data[len(data)-4:]
	var err error
	mic, err = CalculateMIC(key, data[0:len(data)-4])
	if err != nil {
		return false, err
	}
	if len(mic) != 4 {
		return false, errors.New("plink: a MIC of 4 bytes is expected")
	}
	log.Info("calculateMic: ", mic)
	log.Info("calculate input Mic: ", MIC[:])
	for i, v := range mic {
		if MIC[i] != v {
			return false, nil
		}
	}
	return true, nil
}
func CalculateMIC(key common.AES128Key, micBytes []byte) ([]byte, error) {
	hash, err := cmac.New(key[:])
	if err != nil {
		return []byte{}, err
	}
	if _, err = hash.Write(micBytes); err != nil {
		return nil, err
	}
	hb := hash.Sum([]byte{})
	if len(hb) < 4 {
		return []byte{}, errors.New("plink: the hash returned less than 4 bytes")
	}
	fmt.Println("calculate Mic:", hb[:])
	return hb[0:4], nil
}

func EncryptDataPayload(appKey common.AES128Key, data []byte) ([]byte, error) {
	return encrypt(appKey, data)
}
func encrypt(appKey common.AES128Key, data []byte) ([]byte, error) {
	dest := make([]byte, len(data))
	block, err := aes.NewCipher(appKey[:])
	if err != nil {
		return dest, err
	}
	if block.BlockSize() != 16 {
		return dest, errors.New("block is not 16")
	}
	if len(data)%16 != 0 {
		return dest, errors.New("plaintext must be a multiple of 16 bytes")
	}
	for i := 0; i < len(dest)/16; i++ {
		offset := i * 16
		block.Encrypt(dest[offset:offset+16], data[offset:offset+16])
	}
	return dest, nil
}

func DecryptDataPayload(appKey common.AES128Key, data []byte) ([]byte, error) {
	return decrypt(appKey, data)
}
func decrypt(appKey common.AES128Key, data []byte) ([]byte, error) {
	dest := make([]byte, len(data))
	block, err := aes.NewCipher(appKey[:])
	if err != nil {
		return dest, err
	}
	if block.BlockSize() != 16 {
		return dest, errors.New("block is not 16")
	}
	if len(data)%16 != 0 {
		return dest, errors.New("plaintext must be a multiple of 16 bytes")
	}
	for i := 0; i < len(data)/16; i++ {
		offset := i * 16
		block.Encrypt(dest[offset:offset+16], data[offset:offset+16])
	}
	return dest, nil
}
