package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pugdag/pugdagd/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.PugdagdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.PugdagdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.PugdagdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.PugdagdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.PugdagdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.PugdagdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.PugdagdMessage_BanRequest{}),
	reflect.TypeOf(protowire.PugdagdMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
