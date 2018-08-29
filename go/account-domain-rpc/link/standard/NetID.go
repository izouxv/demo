package standard

import (
	"encoding/hex"
	"errors"
	"fmt"
)

type NetID [3]byte

// String implements fmt.Stringer.
func (n NetID) String() string {
	return hex.EncodeToString(n[:])
}

// MarshalText implements encoding.TextMarshaler.
func (n NetID) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

// UnmarshalText implements encoding.TextUnmarshaler.
func (n *NetID) UnmarshalText(text []byte) error {
	b, err := hex.DecodeString(string(text))
	if err != nil {
		return err
	}
	if len(b) != len(n) {
		return fmt.Errorf("lorawan: exactly %d bytes are expected", len(n))
	}
	copy(n[:], b)
	return nil
}

// NwkID returns the NwkID bits of the NetID.
func (n NetID) NwkID() byte {
	return n[2] & 127 // 7 lsb
}

// marshalBinaary
func (n *NetID) MarshalBinary() ([]byte, error) {
	var out [3]byte
	for i, v := range n {
		out[i] = v
	}
	return out[:], nil
}

// UnmarshalBinary
func (n *NetID) UnmarshalBinary(data []byte) error {
	if len(data) != len(n) {
		return errors.New("plink: 3 bytes of data are expected")
	}
	for i, v := range data {
		n[i] = v
	}
	return nil
}
