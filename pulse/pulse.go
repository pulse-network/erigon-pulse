// Package pulse implements the PulseChain fork
package pulse

import (
	"github.com/ledgerwatch/erigon-lib/chain"
	"github.com/ledgerwatch/erigon/core/state"
)

// The first ethereum mainnet pos/beacon block
var firstBeaconBlock = uint64(15537394)

// IsBeaconBlock Returns true if the given block is after the ethereum mainnet merge
func IsBeaconBlock(number uint64) bool {
	return number >= firstBeaconBlock
}

// PrimordialPulseFork Apply PrimordialPulse fork changes
func PrimordialPulseFork(state *state.IntraBlockState, pulseChainConfig *chain.PulseChain) {
	applySacrificeCredits(state, pulseChainConfig)
	replaceDepositContract(state)
}
