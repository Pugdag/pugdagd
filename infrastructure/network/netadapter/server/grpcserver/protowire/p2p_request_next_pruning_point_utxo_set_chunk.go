package protowire

import (
	"github.com/pugdag/pugdagd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PugdagdMessage_RequestNextPruningPointUtxoSetChunk) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PugdagdMessage_RequestNextPruningPointUtxoSetChunk is nil")
	}
	return &appmessage.MsgRequestNextPruningPointUTXOSetChunk{}, nil
}

func (x *PugdagdMessage_RequestNextPruningPointUtxoSetChunk) fromAppMessage(_ *appmessage.MsgRequestNextPruningPointUTXOSetChunk) error {
	x.RequestNextPruningPointUtxoSetChunk = &RequestNextPruningPointUtxoSetChunkMessage{}
	return nil
}
