package handler

import (
	"account-domain-rpc/common"
	"crypto/aes"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	log "github.com/cihub/seelog"
	"github.com/jacobsa/crypto/cmac"
)

// MarshalBinary phy to []byte
func (phy PHYPayload) MarshalBinary() ([]byte, error) {
	if phy.Payload == nil {
		return []byte{}, errors.New("plink: Data should not be nil")
	}
	var out []byte
	var b []byte
	var err error

	out = append(out, byte(phy.Header))

	if b, err = phy.Payload.MarshalBinary(); err != nil {
		return []byte{}, err
	}
	out = append(out, b...)
	out = append(out, phy.MIC[0:len(phy.MIC)]...)
	return out, nil
}

// UnmarshalBinary []byte to Phy
func (phy *PHYPayload) UnmarshalBinary(data []byte) error {
	if len(data) < 3 {
		return errors.New("plink: at least 3 bytes needed to decode PHYPayload")
	}
	phy.Header = Header(data[0])
	// MACPayload
	switch byte(phy.Header) {
	case JoinRequest:
		phy.Payload = &JoinRequestPayload{}
	case JoinAccept:
		phy.Payload = &JoinAcceptPayload{}
	default:
		return errors.New("plink: no type to decode PHYPayload,only JoinRequest")
	}
	if err := phy.Payload.UnmarshalBinary(data[1 : len(data)-4]); err != nil {
		return err
	}
	// MIC
	for i := 0; i < 4; i++ {
		phy.MIC[i] = data[len(data)-4+i]
	}
	return nil
}

// calculateJoinRequestMIC calculates and returns the join-request MIC.
func (phy PHYPayload) CalculateJoinRequestMIC(key common.AES128Key) ([]byte, error) {
	joinRequestPayload, ok := phy.Payload.(*JoinRequestPayload)
	if !ok {
		return []byte{}, errors.New("pslink: data should be of type *JoinRequestPayload")
	}
	var b []byte
	var err error
	var micBytes []byte

	micBytes = append(micBytes, byte(phy.Header))

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
		return []byte{}, errors.New("pslink: the hash returned less than 4 bytes")
	}
	return hb[0:4], nil
}

// calculateJoinAcceptMIC
func (phy PHYPayload) CalculateJoinAcceptMIC(key common.AES128Key) ([]byte, error) {
	if phy.Payload == nil {
		return []byte{}, errors.New("pslink: Payload should not be empty")
	}
	jaPayload, ok := phy.Payload.(*JoinAcceptPayload)
	if !ok {
		return []byte{}, errors.New("pslink: Payload should be of type *JoinAcceptPayload")
	}

	micBytes := make([]byte, 0, 20)

	micBytes = append(micBytes, byte(phy.Header))

	b, err := jaPayload.MarshalBinary()
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
		return []byte{}, errors.New("pslink: the hash returned less than 4 bytes")
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
	case *JoinAcceptPayload:
		mic, err = phy.CalculateJoinAcceptMIC(key)
	default:
		return false, errors.New("plink: not type")
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
