package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateVerifiableRandomNumber{}

func NewMsgCreateVerifiableRandomNumber(creator string, shaseed []byte, publickey []byte, r []byte, s []byte, maxRange int64) *MsgCreateVerifiableRandomNumber {
	return &MsgCreateVerifiableRandomNumber{
		Creator:   creator,
		Shaseed:   shaseed,
		Publickey: publickey,
		R:         r,
		S:         s,
		MaxRange:  maxRange,
	}
}

func (msg *MsgCreateVerifiableRandomNumber) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
