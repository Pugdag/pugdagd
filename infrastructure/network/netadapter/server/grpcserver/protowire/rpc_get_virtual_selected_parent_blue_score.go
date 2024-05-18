package protowire

import (
	"github.com/Pugdag/pugdagd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PugdagdMessage_GetVirtualSelectedParentBlueScoreRequest) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PugdagdMessage_GetVirtualSelectedParentBlueScoreRequest is nil")
	}
	return &appmessage.GetVirtualSelectedParentBlueScoreRequestMessage{}, nil
}

func (x *PugdagdMessage_GetVirtualSelectedParentBlueScoreRequest) fromAppMessage(message *appmessage.GetVirtualSelectedParentBlueScoreRequestMessage) error {
	x.GetVirtualSelectedParentBlueScoreRequest = &GetVirtualSelectedParentBlueScoreRequestMessage{}
	return nil
}

func (x *PugdagdMessage_GetVirtualSelectedParentBlueScoreResponse) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PugdagdMessage_GetVirtualSelectedParentBlueScoreResponse is nil")
	}
	return x.GetVirtualSelectedParentBlueScoreResponse.toAppMessage()
}

func (x *PugdagdMessage_GetVirtualSelectedParentBlueScoreResponse) fromAppMessage(message *appmessage.GetVirtualSelectedParentBlueScoreResponseMessage) error {
	var err *RPCError
	if message.Error != nil {
		err = &RPCError{Message: message.Error.Message}
	}
	x.GetVirtualSelectedParentBlueScoreResponse = &GetVirtualSelectedParentBlueScoreResponseMessage{
		BlueScore: message.BlueScore,
		Error:     err,
	}
	return nil
}

func (x *GetVirtualSelectedParentBlueScoreResponseMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "GetVirtualSelectedParentBlueScoreResponseMessage is nil")
	}
	rpcErr, err := x.Error.toAppMessage()
	// Error is an optional field
	if err != nil && !errors.Is(err, errorNil) {
		return nil, err
	}

	if rpcErr != nil && x.BlueScore != 0 {
		return nil, errors.New("GetVirtualSelectedParentBlueScoreResponseMessage contains both an error and a response")
	}

	return &appmessage.GetVirtualSelectedParentBlueScoreResponseMessage{
		BlueScore: x.BlueScore,
		Error:     rpcErr,
	}, nil
}
