package ibc

import (
	. "github.com/kaifei-bianjie/msg-parser/modules"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

type DocMsgRecvPacket struct {
	Packet          Packet `bson:"packet"`
	ProofCommitment string `bson:"proof_commitment"`
	ProofHeight     Height `bson:"proof_height"`
	Signer          string `bson:"signer"`
}

func (m *DocMsgRecvPacket) GetType() string {
	return MsgTypeRecvPacket
}

func (m *DocMsgRecvPacket) BuildMsg(v interface{}) {
	msg := v.(*MsgRecvPacket)
	m.Signer = msg.Signer
	m.ProofHeight = loadHeight(msg.ProofHeight)
	m.ProofCommitment = utils.MarshalJsonIgnoreErr(msg.ProofCommitment)
	m.Packet = loadPacket(msg.Packet)

}

func (m *DocMsgRecvPacket) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgRecvPacket
	)

	utils.UnMarshalJsonIgnoreErr(utils.MarshalJsonIgnoreErr(v), &msg)
	packetData := UnmarshalPacketData(msg.Packet.GetData())
	addrs = append(addrs, msg.Signer, packetData.Receiver, packetData.Sender)
	handler := func() (Msg, []string) {
		return m, addrs
	}
	return CreateMsgDocInfo(v, handler)
}
