package types

import (
	"github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

var _ sdk.Msg = &MsgCreateCredDef{}

func NewMsgCreateCredDef(id string, schemaId string, tag string, signatureType string, value *types.Any) *MsgCreateCredDef {
	return &MsgCreateCredDef{
		Id:            id,
		SchemaId:      schemaId,
		Tag:           tag,
		SignatureType: signatureType,
		Value:         value,
	}
}

func (msg *MsgCreateCredDef) Route() string {
	return RouterKey
}

func (msg *MsgCreateCredDef) Type() string {
	return "CreateCredDef"
}

func (msg *MsgCreateCredDef) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{}
}

func (msg *MsgCreateCredDef) GetSignBytes() []byte {
	return []byte{}
}

func (msg *MsgCreateCredDef) ValidateBasic() error {
	return nil
}