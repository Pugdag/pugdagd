package protowire

import (
	"github.com/Pugdag/pugdagd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PugdagdMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PugdagdMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *PugdagdMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
