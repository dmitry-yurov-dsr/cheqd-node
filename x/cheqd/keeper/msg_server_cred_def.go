package keeper

import (
	"context"
	"fmt"

	"github.com/cheqd/cheqd-node/x/cheqd/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) CreateCredDef(goCtx context.Context, msg *types.MsgWriteRequest) (*types.MsgCreateCredDefResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	credDefMsg, isMsgIdentity := msg.Data.GetCachedValue().(*types.MsgCreateCredDef)
	if !isMsgIdentity {
		errMsg := fmt.Sprintf("unrecognized %s message type: %T", types.ModuleName, msg)
		return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
	}
	k.AppendCredDef(
		ctx,
		credDefMsg.Id,
		credDefMsg.SchemaId,
		credDefMsg.Tag,
		credDefMsg.SignatureType,
		credDefMsg.Value,
	)

	return &types.MsgCreateCredDefResponse{
		Id: credDefMsg.Id,
	}, nil
}