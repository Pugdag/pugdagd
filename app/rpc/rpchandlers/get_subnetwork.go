package rpchandlers

import (
	"github.com/Pugdag/pugdagd/app/appmessage"
	"github.com/Pugdag/pugdagd/app/rpc/rpccontext"
	"github.com/Pugdag/pugdagd/infrastructure/network/netadapter/router"
)

// HandleGetSubnetwork handles the respectively named RPC command
func HandleGetSubnetwork(context *rpccontext.Context, _ *router.Router, request appmessage.Message) (appmessage.Message, error) {
	response := &appmessage.GetSubnetworkResponseMessage{}
	response.Error = appmessage.RPCErrorf("not implemented")
	return response, nil
}
