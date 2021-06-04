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
	got := blockchain.GetUTXOForAmt(100, GENPK)
	want := blockchain....

	if got != want {
		t.Fail()
	}
}

func TestAdd(t *testing.T) {
	got := ...
	want := ...
	if got != want {
		f.Fail()
	}
}