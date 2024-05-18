package protowire

import (
	"github.com/Pugdag/pugdagd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PugdagdMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PugdagdMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *PugdagdMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
