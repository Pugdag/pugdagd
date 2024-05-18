package protowire

import (
	"github.com/Pugdag/pugdagd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *PugdagdMessage_RequestBlockLocator) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "PugdagdMessage_RequestBlockLocator is nil")
	}
	return x.RequestBlockLocator.toAppMessage()
}

func (x *RequestBlockLocatorMessage) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "RequestBlockLocatorMessage is nil")
	}

	highHash, err := x.HighHash.toDomain()
	if err != nil {
		return nil, err
	}

	return &appmessage.MsgRequestBlockLocator{
		HighHash: highHash,
		Limit:    x.Limit,
	}, nil

}

func (x *PugdagdMessage_RequestBlockLocator) fromAppMessage(msgGetBlockLocator *appmessage.MsgRequestBlockLocator) error {
	x.RequestBlockLocator = &RequestBlockLocatorMessage{
		HighHash: domainHashToProto(msgGetBlockLocator.HighHash),
		Limit:    msgGetBlockLocator.Limit,
	}

	return nil
}
