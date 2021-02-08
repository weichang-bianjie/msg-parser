module github.com/kaifei-bianjie/msg-parser

go 1.14

require (
	github.com/bianjieai/iritamod v0.0.0-20201202112849-cdce8a8df2d0
	github.com/cosmos/cosmos-sdk v0.40.0-rc3
	github.com/irisnet/irismod v1.1.1-0.20201126013702-4999558204d6
	github.com/jolestar/go-commons-pool v2.0.0+incompatible // indirect
	github.com/stretchr/testify v1.6.1
	github.com/tendermint/tendermint v0.34.0-rc6
	gitlab.bianjie.ai/cschain/cschain v1.1.1-0.20201203013652-8cce6b01a507
	go.uber.org/zap v1.15.0 // indirect
	gopkg.in/mgo.v2 v2.0.0-20190816093944-a6b53ec6cb22 // indirect
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.34.4-0.20201127022001-791921d241f8
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.1-dev0.0.20201126055325-2217bc51b6c7
)
