package strategy

import (
	"account-domain-rpc/storage"
)

type Context interface {
	NewDataContext(MType byte, data []byte) error
	HandlerDataPayload(node storage.Node) ([]byte, error)
}
