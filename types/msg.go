package types

import "github.com/cosmos/cosmos-sdk/types"

type (
	DocTxMsg struct {
		Type string `bson:"type"`
		Msg  Msg    `bson:"msg"`
	}

	Msg interface {
		GetType() string
		BuildMsg(msg interface{})
	}

	SdkMsg types.Msg
)
