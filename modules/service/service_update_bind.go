package service

import (
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/modules"
	models "github.com/kaifei-bianjie/msg-parser/types"
)

type (
	DocMsgUpdateServiceBinding struct {
		ServiceName string       `bson:"service_name" yaml:"service_name"`
		Provider    string       `bson:"provider" yaml:"provider"`
		Deposit     models.Coins `bson:"deposit" yaml:"deposit"`
		Pricing     string       `bson:"pricing" yaml:"pricing"`
		QoS         string       `bson:"qos" yaml:"qos"`
		Owner       string       `bson:"owner" yaml:"owner"`
	}
)

func (m *DocMsgUpdateServiceBinding) GetType() string {
	return MsgTypeUpdateServiceBinding
}

func (m *DocMsgUpdateServiceBinding) BuildMsg(v interface{}) {
	msg := v.(*MsgUpdateServiceBinding)

	m.ServiceName = msg.ServiceName
	m.Provider = msg.Provider
	m.Deposit = models.BuildDocCoins(msg.Deposit)
	m.Pricing = msg.Pricing
	m.QoS = fmt.Sprint(msg.QoS)
	m.Owner = msg.Owner
}

func (m *DocMsgUpdateServiceBinding) HandleTxMsg(v SdkMsg) MsgDocInfo {
	var (
		addrs []string
		msg   MsgUpdateServiceBinding
	)

	ConvertMsg(v, &msg)
	addrs = append(addrs, msg.Owner, msg.Provider)
	handler := func() (Msg, []string) {
		return m, addrs
	}

	return CreateMsgDocInfo(v, handler)
}
