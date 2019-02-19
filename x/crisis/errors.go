package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type CodeType = sdk.CodeType

const (
	DefaultCodespace sdk.CodespaceType = ModuleName
	CodeInvalidInput CodeType          = 103
)

func ErrNilSender(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "sender address is nil")
}

func ErrUnknownInvariant(codespace sdk.CodespaceType) sdk.Error {
	return sdk.NewError(codespace, CodeInvalidInput, "unknown invariant")
}
