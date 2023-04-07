package xrpl

// XRPL stream types as defined in rippled:
//  1. https://github.com/XRPLF/xrpl.js/blob/main/packages/xrpl/src/models/common/index.ts#L36
//  2. https://github.com/XRPLF/rippled/blob/master/src/ripple/rpc/handlers/Subscribe.cpp#L127
const (
	StreamTypeLedger               = "ledger"
	StreamTypeTransaction          = "transactions"
	StreamTypeTransactionsProposed = "transactions_proposed"
	StreamTypeValidations          = "validations"
	StreamTypeManifests            = "manifests"
	StreamTypePeerStatus           = "peer_status"
	StreamTypeConsensus            = "consensus"
	StreamTypePathFind             = "path_find"
	StreamTypeServer               = "server"
)

// StreamResponseType returns a string denoting 'type' property present in the
// requested StreamType's response. It returns the empty string if there's no
// match for the requested StreamType.
func StreamResponseType(streamType string) string {
	switch streamType {
	case StreamTypeLedger:
		return "ledgerClosed"
	case StreamTypeTransaction:
		return "transaction"
	case StreamTypeTransactionsProposed:
		return "transaction"
	case StreamTypeValidations:
		return "validationReceived"
	case StreamTypeManifests:
		return "manifestReceived"
	case StreamTypePeerStatus:
		return "peerStatusChange"
	case StreamTypeConsensus:
		return "consensusPhase"
	case StreamTypePathFind:
		return "path_find"
	case StreamTypeServer:
		return "serverStatus"
	default:
		return ""
	}
}
