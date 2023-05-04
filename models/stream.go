package models

import (
	"strconv"

	"github.com/xrpscan/platform/xrpl"
)

type Stream int8

const (
	StreamLedger Stream = iota
	StreamTransaction
	StreamValidation
	StreamManifest
	StreamPeerStatus
	StreamConsensus
	StreamPathFind
	StreamServer
	StreamResponse
)

func (s Stream) String() string {
	switch s {
	case StreamLedger:
		return xrpl.StreamTypeLedger
	case StreamTransaction:
		return xrpl.StreamTypeTransaction
	case StreamValidation:
		return xrpl.StreamTypeValidations
	case StreamManifest:
		return xrpl.StreamTypeManifests
	case StreamPeerStatus:
		return xrpl.StreamTypePeerStatus
	case StreamConsensus:
		return xrpl.StreamTypeConsensus
	case StreamPathFind:
		return xrpl.StreamTypePathFind
	case StreamServer:
		return xrpl.StreamTypeServer
	case StreamResponse:
		return xrpl.StreamTypeResponse
	}
	return strconv.Itoa(int(s))
}
