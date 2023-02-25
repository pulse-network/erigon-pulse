// Package pulse implements the PulseChain fork
package pulse

import (
	"github.com/ledgerwatch/erigon-lib/chain"
	"github.com/ledgerwatch/erigon/core/state"
)

// PrimordialPulseFork Apply PrimordialPulse fork changes
func PrimordialPulseFork(state *state.IntraBlockState, pulseChainConfig *chain.PulseChain) {
	applySacrificeCredits(state, pulseChainConfig)
	replaceDepositContract(state)
}
