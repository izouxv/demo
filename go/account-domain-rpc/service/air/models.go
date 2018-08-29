package air

import (
	"account-domain-rpc/common"
	"account-domain-rpc/link/standard"
	"crypto/aes"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/jacobsa/crypto/cmac"
)

const (
	Successful  = 0x00
	Fail        = 0x01
	FailCode    = 0x00
	SystemError = 0x01
	AlreadyBind = 0x02
)

type MType byte
type MLen byte
type MHDR struct {
	MType MType
	MLen  MLen
}

// String
func (m MType) String() string {
	return fmt.Sprintf("Major(%d)", m)
}

// MarshalText implements encoding.TextMarshaler.
func (m MType) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

// MarshalBinary marshals the object in binary form.
func (h MHDR) MarshalBinary() ([]byte, error) {
	return []byte{byte(h.MType), byte(h.MLen)}, nil
}

// UnmarshalBinary decodes the object from binary form.
func (h *MHDR) UnmarshalBinary(data []byte) error {
	if len(data) != 2 {
		return errors.New("plink: 2 byte of data is expected")
	}
	h.MType = MType(data[0])
	h.MLen = MLen(data[1])
	return nil
}

// MIC represents the message integrity code.
type MIC [4]byte

// String implements fmt.Stringer.
func (m MIC) String() string {
	return hex.EncodeToString(m[:])
}

// MarshalText implements encoding.TextMarshaler.
func (m MIC) MarshalText() ([]byte, error) {
	return []byte(m.String()), nil
}

type Payload interface {
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(data []byte) error
}

// JoinRequestPayload represents the join-request message payload.
type JoinRequestPayload struct {
	AppEUI   common.EUI64      `json:"appEUI"`
	DevEUI   common.EUI64      `json:"devEUI"`
	DevNonce standard.DevNonce `json:"devNonce"`
}

// MarshalBinary marshals the object in binary form.
func (p JoinRequestPayload) MarshalBinary() ([]byte, error) {
	out := make([]byte, 0, 18)
	b, err := p.AppEUI.MarshalBinary()
	if err != nil {
		return nil, err
	}
	out = append(out, b...)
	b, err = p.DevEUI.MarshalBinary()
	if err != nil {
		return nil, err
	}
	out = append(out, b...)
	out = append(out, p.DevNonce[:]...)
	return out, nil
}

// UnmarshalBinary decodes the object from binary form.
func (p *JoinRequestPayload) UnmarshalBinary(data []byte) error {
	if len(data) != 18 {
		return errors.New("plink: 18 bytes of data are expected")
	}
	if err := p.AppEUI.UnmarshalBinary(data[0:8]); err != nil {
		return err
	}
	if err := p.DevEUI.UnmarshalBinary(data[8:16]); err != nil {
		return err
	}
	p.DevNonce[0] = data[16]
	p.DevNonce[1] = data[17]
	return nil
}

type JoinAcceptPayload struct {
	NetID    standard.NetID    `json:"netId"`
	DevEUI   common.EUI64      `json:"devEUI"`
	AppNonce standard.AppNonce `json:"appNonce"`
	Pading   []byte            `json:"pading"`
}

// MarshalBinary marshals the object in binary form.
func (p JoinAcceptPayload) MarshalBinary() ([]byte, error) {
	out := make([]byte, 0, 16)
	b, err := p.NetID.MarshalBinary()
	if err != nil {
		return nil, err
	}
	out = append(out, b...)
	b, err = p.DevEUI.MarshalBinary()
	if err != nil {
		return nil, err
	}
	out = append(out, b...)
	out = append(out, p.AppNonce[:]...)
	out = append(out, p.Pading...)
	return out, nil
}

// UnmarshalBinary decodes the object from binary form.
func (p *JoinAcceptPayload) UnmarshalBinary(data []byte) error {
	if len(data) != 16 {
		return errors.New("plink: 16 bytes of data are expected")
	}
	if err := p.NetID.UnmarshalBinary(data[0:3]); err != nil {
		return err
	}
	if err := p.DevEUI.UnmarshalBinary(data[3:11]); err != nil {
		return err
	}
	p.AppNonce[0] = data[11]
	p.AppNonce[1] = data[12]
	p.AppNonce[2] = data[13]
	p.Pading = data[14:]
	return nil
}

// calculateJoinRequestMIC calculates and returns the join-request MIC.
func (phy PHYPayload) CalculateJoinRequestMIC(key common.AES128Key) ([]byte, error) {
	joinRequestPayload, ok := phy.Payload.(*JoinRequestPayload)
	if !ok {
		return []byte{}, errors.New("plink: data should be of type *JoinRequestPayload")
	}
	var b []byte
	var err error
	var micBytes []byte

	b, err = phy.MHDR.MarshalBinary()
	micBytes = append(micBytes, b...)

	b, err = joinRequestPayload.MarshalBinary()
	micBytes = append(micBytes, b...)

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
	return hb[0:4], nil
}

// calculateJoinAcceptMIC
func (phy PHYPayload) CalculateJoinAcceptMIC(key common.AES128Key) ([]byte, error) {
	if phy.Payload == nil {
		return []byte{}, errors.New("lorawan: MACPayload should not be empty")
	}
	jaPayload, ok := phy.Payload.(*JoinAcceptPayload)
	if !ok {
		return []byte{}, errors.New("lorawan: MACPayload should be of type *JoinAcceptPayload")
	}

	micBytes := make([]byte, 0, 20)

	b, err := phy.MHDR.MarshalBinary()
	if err != nil {
		return []byte{}, err
	}
	micBytes = append(micBytes, b...)

	b, err = jaPayload.MarshalBinary()
	if err != nil {
		return nil, err
	}
	micBytes = append(micBytes, b...)

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
	return hb[0:4], nil
}

// validataMIC
func (phy PHYPayload) ValidateMIC(key common.AES128Key) (bool, error) {
	var mic []byte
	var err error
	switch phy.Payload.(type) {
	case *JoinRequestPayload:
		mic, err = phy.CalculateJoinRequestMIC(key)
	default:
		return false, errors.New("node type")
	}

	if err != nil {
		return false, err
	}
	if len(mic) != 4 {
		return false, errors.New("plink: a MIC of 4 bytes is expected")
	}
	log.Info("calculateMic: ", mic)
	for i, v := range mic {
		if phy.MIC[i] != v {
			return false, nil
		}
	}
	return true, nil
}

// EncryptAesEcbJoinAcceptPayload
func (phy *PHYPayload) EncryptJoinAcceptPayload(appKey common.AES128Key) error {
	if _, ok := phy.Payload.(*JoinAcceptPayload); !ok {
		return errors.New("plink: Payload value must be of type *JoinAcceptPayload")
	}
	payload, err := phy.Payload.MarshalBinary()
	if err != nil {
		return err
	}
	dest := make([]byte, 16)
	block, err := aes.NewCipher(appKey[:])
	if err != nil {
		return err
	}
	if block.BlockSize() != len(payload) {
		return fmt.Errorf("block-size of %d bytes is expected", len(payload))
	}
	log.Info("Payload: ", payload)
	block.Encrypt(dest, payload)
	jaPL := &JoinAcceptPayload{}
	jaPL.UnmarshalBinary(dest)
	phy.Payload = jaPL
	return nil
}

// DecryptJoinAcceptPayload
func (phy *PHYPayload) DecryptJoinAcceptPayload(appKey common.AES128Key) error {
	if _, ok := phy.Payload.(*JoinAcceptPayload); !ok {
		return errors.New("plink: Payload value must be of type *JoinAcceptPayload")
	}
	payload, err := phy.Payload.MarshalBinary()
	if err != nil {
		return err
	}
	dest := make([]byte, 16)
	block, err := aes.NewCipher(appKey[:])
	if err != nil {
		return err
	}
	if block.BlockSize() != len(payload) {
		return fmt.Errorf("block-size of %d bytes is expected", len(payload))
	}
	block.Decrypt(dest, payload)
	jaPL := &JoinAcceptPayload{}
	jaPL.UnmarshalBinary(dest)
	phy.Payload = jaPL
	return nil
}

// MarshalText encodes the PHYPayload into base64.
func (phy PHYPayload) MarshalText() ([]byte, error) {
	b, err := phy.MarshalBinary()
	if err != nil {
		return nil, err
	}
	return []byte(base64.StdEncoding.EncodeToString(b)), nil
}

// UnmarshalText decodes the PHYPayload from base64.
func (phy *PHYPayload) UnmarshalText(text []byte) error {
	b, err := base64.StdEncoding.DecodeString(string(text))
	if err != nil {
		return err
	}
	return phy.UnmarshalBinary(b)
}

// MarshalJSON encodes the PHYPayload into JSON.
func (phy PHYPayload) MarshalJSON() ([]byte, error) {
	type phyAlias PHYPayload
	return json.Marshal(phyAlias(phy))
}

// set MIC
func (phy *PHYPayload) SetMIC(key common.AES128Key) error {
	var mic []byte
	var err error

	switch phy.Payload.(type) {
	case *JoinRequestPayload:
		mic, err = phy.CalculateJoinRequestMIC(key)
	case *JoinAcceptPayload:
		mic, err = phy.CalculateJoinAcceptMIC(key)
	default:
		return errors.New("not type")
	}

	if err != nil {
		return err
	}
	if len(mic) != 4 {
		return errors.New("lorawan: a MIC of 4 bytes is expected")
	}
	fmt.Println("------mic--------", mic)
	for i, v := range mic {
		phy.MIC[i] = v
	}
	return nil
}

type AirInfo struct {
	Boot          int32 `json:"boot"`          //开关机
	Sleep         int32 `json:"sleep"`         //睡眠模式
	WindSpeed     int32 `json:"windSpeed"`     //风速
	TimedShutdown int32 `json:"timedShutdown"` //定时关机
	Anion         int32 `json:"anion"`         //负离子
	Automatic     int32 `json:"automatic"`     //自动模式
	Manual        int32 `json:"manual"`        //手动模式
	UVLamp        int32 `json:"uvLamp"`        //紫外灯
	SuperAir      int32 `json:"superAir"`      //超级风量
	/* BIT7＝0，数据为正；BIT7＝1，数据为负。
	    BIT6-BIT0：温度数据0－127；
	   （接收完成后将其转为signed char类型，就是温度值了）
	*/
	Temperature int32 `json:"temperature"` //温度 （+/-）
	Humidity    int32 `json:"humidity"`    //湿度	% 取值：0(0x00)~100(0x64)
	/*Byte12取高位：0x01
	  Byte13取低位：0x2C
	*/
	PM25 int32 `json:"pm25"` //PM2.5
	/*Byte14取高位：0x00
		  Byte15取低位：0x50
	      取值范围：0x00~0x64
	*/
	FilterCountdown int32 `json:"filterCountdown"` //当滤网剩余使用寿命不足10%时
}

func (air AirInfo) MarshalBinary() ([]byte, error) {
	if err := func(info AirInfo) error {
		if air.Boot>>1 != 0 ||
			air.Sleep>>1 != 0 ||
			air.WindSpeed < 1 || air.WindSpeed > 4 ||
			air.TimedShutdown < 0 || air.TimedShutdown > 12 ||
			air.Anion>>1 != 0 ||
			air.Automatic>>1 != 0 ||
			air.Manual>>1 != 0 ||
			air.UVLamp>>1 != 0 ||
			air.SuperAir>>1 != 0 ||
			air.Temperature > 127 || air.Temperature < -127 ||
			air.Humidity < 0 || air.Humidity > 100 ||
			air.PM25 < 0 ||
			air.FilterCountdown < 0 || air.FilterCountdown > 100 {
			return errors.New("input air info error")
		}
		return nil
	}(air); err != nil {
		return nil, err
	}
	out := [15]byte{}
	out[0] = byte(air.Boot)
	out[1] = byte(air.Sleep)
	out[2] = byte(air.WindSpeed)
	out[3] = byte(air.TimedShutdown)
	out[4] = byte(air.Anion)
	out[5] = byte(air.Automatic)
	out[6] = byte(air.Manual)
	out[7] = byte(air.UVLamp)
	out[8] = byte(air.SuperAir)
	if air.Temperature < 0 {
		temp := byte(air.Temperature * -1)
		out[9] = temp | 128
	} else {
		out[9] = byte(air.Temperature)
	}
	out[10] = byte(air.Humidity)
	out[11] = byte(air.PM25 & 0xFF00 >> 8)
	out[12] = byte(air.PM25 & 0x00FF)
	out[13] = 0x00
	out[14] = byte(air.FilterCountdown)
	return out[:], nil

}

func (air *AirInfo) UnMarshalBinary(value []byte) error {
	if len(value) != 15 {
		return errors.New("input air value error")
	}
	air.Boot = int32(value[0])
	air.Sleep = int32(value[1])
	air.WindSpeed = int32(value[2])
	air.TimedShutdown = int32(value[3])
	air.Anion = int32(value[4])
	air.Automatic = int32(value[5])
	air.Manual = int32(value[6])
	air.UVLamp = int32(value[7])
	air.SuperAir = int32(value[8])
	if value[9] > 127 {
		air.Temperature = -1 * int32(value[9]&0x7F)
	} else {
		air.Temperature = int32(value[9])
	}
	air.Humidity = int32(value[10])
	air.PM25 = int32(value[11])*16 + int32(value[12])
	air.FilterCountdown = int32(value[14])
	return nil
}

func (air AirInfo) MarshalJson() ([]byte, error) {
	value, err := json.Marshal(air)
	if err != nil {
		return value, err
	}
	return value, nil
}

func (air *AirInfo) UnMarshalJson(value []byte) error {
	return json.Unmarshal(value, air)
}

type OrderResponse struct {
	Type int32   `json:"type"`
	Air  AirInfo `json:"airInfo"`
}

func (or OrderResponse) MarshalBinary() ([]byte, error) {
	out := make([]byte, 0)
	out = append(out, byte(or.Type))
	air, err := or.Air.MarshalBinary()
	if err != nil {
		return out, err
	}
	out = append(out, air...)
	return out, nil
}

func (or *OrderResponse) UnMarshalBinary(value []byte) error {
	or.Type = int32(value[0])
	air := AirInfo{}
	if err := air.UnMarshalBinary(value[2:]); err != nil {
		return err
	}
	or.Air = air
	return nil
}

func (or OrderResponse) MarshalJson() ([]byte, error) {
	value, err := json.Marshal(or)
	if err != nil {
		return value, err
	}
	return value, err
}

func (or *OrderResponse) UnMarshalJson(value []byte) error {
	return json.Unmarshal(value, or)
}

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

func (um *UpgradeModele) UnMarshalBinary(in []byte) error {
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
	for _, v := range data[:len(data)-2] {
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

func (urm *UpgradeResultModel) UnMarshalBinary(in []byte) error {
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
