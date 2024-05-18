package protowire

import (
	"github.com/Pugdag/pugdagd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PugdagdMessage_GetCurrentNetworkRequest) toAppMessage() (appmessage.Message, error) {
	return &appmessage.GetCurrentNetworkRequestMessage{}, nil
}

func (x *PugdagdMessage_GetCurrentNetworkRequest) fromAppMessage(_ *appmessage.GetCurrentNetworkRequestMessage) error {
	return nil
}

func (x *PugdagdMessage_GetCurrentNetworkResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PugdagdMessage_GetCurrentNetworkResponse is nil")
	}
	return x.toAppMessage()
}

func (x *PugdagdMessage_GetCurrentNetworkResponse) fromAppMessage(message *appmessage.GetCurrentNetworkResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.GetCurrentNetworkResponse = &GetCurrentNetworkResponseMessage{
		CurrentNetwork: message.CurrentNetwork,
		Error:          err,
	}
	return nil
}

func (x *GetCurrentNetworkResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetCurrentNetworkResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	if rpcErr != nil && len(x.CurrentNetwork) != 0 {
		return nil, errors.New("GetCurrentNetworkResponseMessage contains both an error and a response")
	}

	return &appmessage.GetCurrentNetworkResponseMessage{
		CurrentNetwork: x.CurrentNetwork,
		Error:          rpcErr,
	}, nil
}
