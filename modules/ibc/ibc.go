package ibc

import (
	"github.com/cosmos/cosmos-sdk/types"
	. "github.com/kaifei-bianjie/msg-parser/modules"
)

type ibcClient struct {
}

func NewClient() Client {
	return ibcClient{}
}

func (ibc ibcClient) HandleTxMsg(v types.Msg) (MsgDocInfo, bool) {
	var (
		msgDocInfo MsgDocInfo
	)
	ok := true
	switch v.Type() {
	case new(MsgRecvPacket).Type():
		docMsg := DocMsgRecvPacket{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgCreateClient).Type():
		docMsg := DocMsgCreateClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	case new(MsgUpdateClient).Type():
		docMsg := DocMsgUpdateClient{}
		msgDocInfo = docMsg.HandleTxMsg(v)
	default:
		ok = false
	}
	return msgDocInfo, ok
}
