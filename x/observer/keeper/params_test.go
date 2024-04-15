package keeper_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/inco-fhevm/take-home-starter/x/observer/types"
)

func TestGetParams(t *testing.T) {
	k, ctx := newTestKeeper(t)
	params := types.DefaultParams()

	require.NoError(t, k.SetParams(ctx, params))
	require.EqualValues(t, params, k.GetParams(ctx))
}
