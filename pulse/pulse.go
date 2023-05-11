// Package pulse implements the PulseChain fork
package pulse

import (
	"math/big"

	"github.com/ledgerwatch/erigon-lib/chain"
	"github.com/ledgerwatch/erigon/core/state"
)

var MainnetChainID = big.NewInt(369)
var TestnetV4ChainID = big.NewInt(943)

// PrimordialPulseFork Apply PrimordialPulse fork changes
func PrimordialPulseFork(state *state.IntraBlockState, pulseChainConfig *chain.PulseChain, chainID *big.Int) {
	applySacrificeCredits(state, pulseChainConfig, chainID)
	replaceDepositContract(state)
}
