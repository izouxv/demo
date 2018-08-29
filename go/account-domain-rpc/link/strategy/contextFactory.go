package strategy

import (
	"github.com/pkg/errors"

	"account-domain-rpc/service/air"

	"account-domain-rpc/common"
	"account-domain-rpc/storage"
)

type ContextFactory struct {
	Node storage.Node
}

func (dc ContextFactory) CreateContext(MType byte, data []byte) (Context, error) {
	var context Context
	if dc.Node.Category == 0 {
		return nil, errors.New("CreateContext error")
	}
	switch dc.Node.Category {
	case common.Air:
		context = &air.AirContext{}

	}
	if err := context.NewDataContext(MType, data); err != nil {
		return context, err
	}
	return context, nil
}
