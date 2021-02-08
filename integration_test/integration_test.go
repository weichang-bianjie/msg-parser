package integration

import (
	msg_parser "github.com/kaifei-bianjie/msg-parser"
	. "github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/stretchr/testify/suite"
	"testing"
)

type IntegrationTestSuite struct {
	msg_parser.MsgClient
	suite.Suite
}

type SubTest struct {
	testName string
	testCase func(s IntegrationTestSuite)
}

func TestSuite(t *testing.T) {
	suite.Run(t, new(IntegrationTestSuite))
}

func (s *IntegrationTestSuite) SetupSuite() {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	s.MsgClient = msg_parser.NewMsgClient()
}
