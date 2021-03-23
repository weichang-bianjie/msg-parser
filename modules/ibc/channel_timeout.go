package ibc

import (
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

type DocMsgTimeout struct {
	Packet           Packet `bson:"packet"`
	ProofUnreceived  string `bson:"proof_unreceived"`
	ProofHeight      Height `bson:"proof_height"`
	NextSequenceRecv string `bson:"next_sequence_recv"`
	Signer           string `bson:"signer"`
}

func (m *DocMsgTimeout) GetType() string {
	return MsgTypeTimeout
}

func (m *DocMsgTimeout) BuildMsg(v interface{}) {
	msg := v.(*MsgTimeout)
	m.Signer = msg.Signer
	m.NextSequenceRecv = fmt.Sprint(msg.NextSequenceRecv)
	m.ProofUnreceived = utils.MarshalJsonIgnoreErr(msg.ProofUnreceived)
	m.Packet = loadPacket(msg.Packet)
	m.ProofHeight = loadHeight(msg.ProofHeight)
}

func (m *DocMsgTimeout) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgTimeout
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	addrs = append(addrs, msg.Signer)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
