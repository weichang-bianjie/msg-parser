module github.com/kaifei-bianjie/msg-parser

go 1.15

require (
	github.com/cosmos/cosmos-sdk v0.40.0-rc3
	github.com/irisnet/irismod v1.1.1-0.20201215020504-ae6a23d4bec2
	github.com/stretchr/testify v1.7.0
	github.com/tendermint/tendermint v0.34.0-rc6
)

replace (
	github.com/cosmos/cosmos-sdk => github.com/bianjieai/cosmos-sdk v0.34.4-0.20201127022001-791921d241f8
	github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.2-alpha.regen.4
	github.com/tendermint/tendermint => github.com/bianjieai/tendermint v0.33.1-dev0.0.20201126055325-2217bc51b6c7
)
