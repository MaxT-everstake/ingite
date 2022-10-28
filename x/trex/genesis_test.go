package trex_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	keepertest "trex/testutil/keeper"
	"trex/testutil/nullify"
	"trex/x/trex"
	"trex/x/trex/types"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.TrexKeeper(t)
	trex.InitGenesis(ctx, *k, genesisState)
	got := trex.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	// this line is used by starport scaffolding # genesis/test/assert
}
