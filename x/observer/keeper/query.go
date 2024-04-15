package keeper

import (
	"github.com/inco-fhevm/take-home-starter/x/observer/types"
)

var _ types.QueryServer = Keeper{}
