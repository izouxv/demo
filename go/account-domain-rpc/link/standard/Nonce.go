package standard

import "encoding/hex"

// DevNonce represents a 2 byte dev-nonce.
type DevNonce [2]byte

// String implements fmt.Stringer.
func (n DevNonce) String() string {
	return hex.EncodeToString(n[:])
}

// MarshalText implements encoding.TextMarshaler.
func (n DevNonce) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}

// AppNonce represents a 3 byte app-nonce.
type AppNonce [3]byte

// String implements fmt.Stringer.
func (n AppNonce) String() string {
	return hex.EncodeToString(n[:])
}

// MarshalText implements encoding.TextMarshaler.
func (n AppNonce) MarshalText() ([]byte, error) {
	return []byte(n.String()), nil
}
