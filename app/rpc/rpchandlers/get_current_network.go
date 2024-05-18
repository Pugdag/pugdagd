package rpchandlers

import (
	"github.com/Pugdag/pugdagd/app/appmessage"
	"github.com/Pugdag/pugdagd/app/rpc/rpccontext"
	"github.com/Pugdag/pugdagd/infrastructure/network/netadapter/router"
)

// HandleGetCurrentNetwork handles the respectively named RPC command
func HandleGetCurrentNetwork(context *rpccontext.Context, _ *router.Router, _ appmessage.Message) (appmessage.Message, error) {
	response := appmessage.NewGetCurrentNetworkResponseMessage(context.Config.ActiveNetParams.Net.String())
	return response, nil
}
