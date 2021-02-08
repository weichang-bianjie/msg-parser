package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestBank() {
	cases := []SubTest{
		{
			"TestSend",
			send,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func send(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0a8f010a8c010a1c2f636f736d6f732e62616e6b2e763162657461312e4d736753656e64126c0a2a69616131797768616d68326b63327a383037373632746e6c38706b687970647136663272387438717465122a69616131636e6d657274616c797a653973756530357279326d30666e34376679637a33647468666a64791a120a05706f696e74120931303030303030303012520a4a0a400a192f636f736d6f732e63727970746f2e736d322e5075624b657912230a21020597f2afe375d17e73ffc41d791562c116c5092222412b369e3ba474728e721212040a020801180912041080b5181a40a6628350d0a840a028d164384ad0d83c75a1b1c486b4384d0dc4f1ae159f1c3b81a89c3b1f9b7c13f2dc22efe6c0dd89d1d40624256c61f93492497b0cc857fc")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Bank.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
