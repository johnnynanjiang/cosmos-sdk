package types

import (
	"encoding/json"
	"fmt"
	"os"
)

// Transactions messages must fulfill the Msg
type Msg interface {

	// Return the message type.
	// Must be alphanumeric or empty.
	Route() string

	// Returns a human-readable string for the message, intended for utilization
	// within tags
	Type() string

	// ValidateBasic does a simple validation check that
	// doesn't require access to any other information.
	ValidateBasic() Error

	// Get the canonical byte representation of the Msg.
	GetSignBytes() []byte

	// Signers returns the addrs of signers that must sign.
	// CONTRACT: All signatures must be present to be valid.
	// CONTRACT: Returns addrs in some deterministic order.
	GetSigners() []AccAddress
}

//__________________________________________________________

// Transactions objects must fulfill the Tx
type Tx interface {
	// Gets the all the transaction's messages.
	GetMsgs() []Msg

	// ValidateBasic does a simple and lightweight validation check that doesn't
	// require access to any other information.
	ValidateBasic() Error
}

//__________________________________________________________

// TxDecoder unmarshals transaction bytes
type TxDecoder func(txBytes []byte) (Tx, Error)

// TxEncoder marshals transaction to bytes
type TxEncoder func(tx Tx) ([]byte, error)

//__________________________________________________________

var _ Msg = (*TestMsg)(nil)

// msg type for testing
type TestMsg struct {
	signers []AccAddress
}

func NewTestMsg(addrs ...AccAddress) *TestMsg {
	return &TestMsg{
		signers: addrs,
	}
}

// TW >>>

// MsgSend to send coins from Input to Output
type MsgSend struct {
	From   AccAddress `json:"from"`
	To     AccAddress `json:"to"`
	Amount Coins      `json:"amount"`
}

func NewMsgSend(from, to AccAddress, amt Coins) MsgSend {
	return MsgSend{from, to, amt}
}
// TW <<<

//nolint
func (msg *TestMsg) Route() string { return "TestMsg" }
func (msg *TestMsg) Type() string  { return "Test message" }
func (msg *TestMsg) GetSignBytes() []byte {
	fmt.Fprintln(os.Stdout, "TestMsg.GetSignBytes()")
	fmt.Fprintln(os.Stdout, "msg.signers")
	fmt.Fprintln(os.Stdout, msg.signers)

	bz, err := json.Marshal(msg.signers)
	if err != nil {
		panic(err)
	}

	fmt.Fprintln(os.Stdout, "json.Marshal(msg.signers)")
	fmt.Fprintln(os.Stdout, bz)

	fmt.Fprintln(os.Stdout, "MustSortJSON(bz)")
	fmt.Fprintln(os.Stdout, MustSortJSON(bz))
	fmt.Fprintln(os.Stdout, "")

	return MustSortJSON(bz)
}
func (msg *TestMsg) ValidateBasic() Error { return nil }
func (msg *TestMsg) GetSigners() []AccAddress {
	return msg.signers
}
