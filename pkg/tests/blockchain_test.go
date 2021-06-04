import (
	"BrunoCoin/pkg/block"
	"BrunoCoin/pkg/block/tx"
	"BrunoCoin/pkg/blockchain"
	"BrunoCoin/pkg/id"
	"BrunoCoin/pkg/proto"
	"encoding/hex"
	"sync"
)

blockchain = New(DefaultConfig())


func TestGetUTXOForAmt(t *testing.T) {
	got := blockchain.GetUTXOForAmt(100, )
	want :=
}