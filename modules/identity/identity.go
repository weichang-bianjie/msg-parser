package identity

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type identityClient struct {
}

func NewClient() Client {
	return identityClient{}
}

func (nft identityClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch v.Type() {
	case new(MsgCreateIdentity).Type():
		docMsg := DocMsgCreateIdentity{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	case new(MsgUpdateIdentity).Type():
		docMsg := DocMsgUpdateIdentity{}
		msgDocInfo = docMsg.HandleTxMsg(v)
		break
	default:
		ok = false
	}
	return msgDocInfo, ok
}
