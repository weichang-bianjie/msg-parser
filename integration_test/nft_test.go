package integration

import (
	"encoding/hex"
	"fmt"
	. "github.com/kaifei-bianjie/msg-parser/codec"
	"github.com/kaifei-bianjie/msg-parser/utils"
)

func (s IntegrationTestSuite) TestNft() {
	cases := []SubTest{
		{
			"IssueDenom",
			IssueDenom,
		},
	}

	for _, t := range cases {
		s.Run(t.testName, func() {
			t.testCase(s)
		})
	}
}

func IssueDenom(s IntegrationTestSuite) {
	SetBech32Prefix(Bech32PrefixAccAddr, Bech32PrefixAccPub, Bech32PrefixValAddr,
		Bech32PrefixValPub, Bech32PrefixConsAddr, Bech32PrefixConsPub)
	txBytes, err := hex.DecodeString("0a6d0a6b0a1a2f697269736d6f642e6e66742e4d7367497373756544656e6f6d124d0a117472616e73616374696f6e7265706f7274120ce4baa4e69893e68aa5e5918a222a69616131797768616d68326b63327a383037373632746e6c38706b687970647136663272387438717465125e0a4a0a400a192f636f736d6f732e63727970746f2e736d322e5075624b657912230a21020597f2afe375d17e73ffc41d791562c116c5092222412b369e3ba474728e721212040a020801181712100a0a0a05706f696e741201351080b5181a4086966f7ebe5501c4921eaf5f64fa61c7b2350c5347ca4d09bf0fa8e25c2e0fddffc847ed5b1dc04c9e83cba10cf9f2da0f8ce3707b01502bea507e82bbb7b4c3")
	if err != nil {
		fmt.Println(err.Error())
	}
	authTx, err := GetSigningTx(txBytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	for _, msg := range authTx.GetMsgs() {
		if bankDoc, ok := s.Nft.HandleTxMsg(msg); ok {
			fmt.Println(utils.MarshalJsonIgnoreErr(bankDoc))
		}
	}
}
