package handler

import (
	"account-domain-rpc/common"
	"account-domain-rpc/link/standard"
	"errors"
)

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
