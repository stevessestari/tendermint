package core

import (
	"errors"
	"fmt"

	ctypes "github.com/tendermint/tendermint/rpc/core/types"
	rpctypes "github.com/tendermint/tendermint/rpc/jsonrpc/types"
	"github.com/tendermint/tendermint/types"
)

// BroadcastHeaderTrace broadcasts evidence of the misbehavior.
// More: https://docs.tendermint.com/master/rpc/#/Info/broadcast_evidence
func BroadcastHeaderTrace(ctx *rpctypes.Context, trace *types.ConflictingHeadersTrace) (*ctypes.ResultBroadcastEvidence, error) {
	if trace == nil {
		return nil, errors.New("no evidence was provided")
	}

	if err := trace.ValidateBasic(); err != nil {
		return nil, fmt.Errorf("evidence.ValidateBasic failed: %w", err)
	}

	// TODO: Pass trace on to evidence pool to handle 
	// if err := env.EvidencePool.AddEvidence(trace); err != nil {
	// 	return nil, fmt.Errorf("failed to add evidence: %w", err)
	// }
	return &ctypes.ResultBroadcastEvidence{Hash: trace.Hash()}, nil
}
