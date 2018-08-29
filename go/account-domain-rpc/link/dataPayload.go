package link

import (
	"account-domain-rpc/storage"
)

type DataPayload interface {
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(data []byte) error
	HandlerDataPayload(node storage.Node) ([]byte, error)
}
