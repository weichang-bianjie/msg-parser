package record

type Content struct {
	DigestAlgo string `bson:"digest_algo"`
	Digest     string `bson:"digest"`
	URI        string `bson:"uri"`
	Meta       string `bson:"meta"`
}

type Packet struct {
	ID        string     `bson:"id"`
	Timestamp uint64     `bson:"timestamp"`
	Height    uint64     `bson:"height"`
	TxHash    string     `bson:"tx_hash"`
	Contents  []*Content `bson:"contents"`
	Creator   string     `bson:"creator"`
}
