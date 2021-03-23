package random

import (
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type DocTxMsgRequestRand struct {
	Consumer      string        `bson:"consumer"`       // request address
	BlockInterval string        `bson:"block_interval"` // block interval after which the requested random number will be generated
	Oracle        bool          `bson:"oracle"`
	ServiceFeeCap []models.Coin `bson:"service_fee_cap"`
}

func (doctx *DocTxMsgRequestRand) GetType() string {
	return TxTypeRequestRand
}

func (doctx *DocTxMsgRequestRand) BuildMsg(txMsg interface{}) {
	msg := txMsg.(*MsgRequestRandom)
	doctx.Consumer = msg.Consumer
	doctx.BlockInterval = fmt.Sprint(msg.BlockInterval)
	doctx.Oracle = msg.Oracle
	doctx.ServiceFeeCap = models.BuildDocCoins(msg.ServiceFeeCap)
}

func (doctx *DocTxMsgRequestRand) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgRequestRandom
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Consumer)
	handler := func() (Msg, []string) {
		return doctx, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
