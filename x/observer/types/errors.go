package types

// DONTCOVER

import (
	sdkerrors "cosmossdk.io/errors"
)

// x/observer module sentinel errors
var (
	ErrInvalidSigner        = sdkerrors.Register(ModuleName, 1100, "expected gov account as only signer for proposal message")
	ErrUnauthorizedObserver = sdkerrors.Register(ModuleName, 1101, "observer is not authorized to submit reports")
	ErrInvalidEthTxHash     = sdkerrors.Register(ModuleName, 1102, "invalid Ethereum transaction hash")
)
