package node

import (
	"encoding/hex"

	"github.com/asciiu/appa/sandbox-btc/protocol"
	"github.com/sirupsen/logrus"
)

// Mempool represents mempool.
type Mempool struct {
	NewBlockCh chan protocol.MsgBlock
	NewTxCh    chan protocol.MsgTx

	txs map[string]*protocol.MsgTx
}

// NewMempool returns a new Mempool.
func NewMempool() *Mempool {
	return &Mempool{
		NewBlockCh: make(chan protocol.MsgBlock),
		NewTxCh:    make(chan protocol.MsgTx),
		txs:        make(map[string]*protocol.MsgTx),
	}
}

// Run starts mempool state handling.
func (m Mempool) Run() {
	for {
		select {
		case tx := <-m.NewTxCh:
			hash, err := tx.Hash()
			if err != nil {
				logrus.Errorf("failed to calculate tx hash: %+v", err)
				break
			}

			txid := hex.EncodeToString(hash)
			m.txs[txid] = &tx
		case block := <-m.NewBlockCh:
			for _, tx := range block.Txs {
				hash, err := tx.Hash()
				if err != nil {
					logrus.Errorf("failed to calculate tx hash: %+v", err)
					break
				}

				txid := hex.EncodeToString(hash)
				delete(m.txs, txid)
			}
		}
	}
}

// Mempool ...
func (n Node) Mempool() map[string]*protocol.MsgTx {
	m := make(map[string]*protocol.MsgTx)

	for k, v := range n.mempool.txs {
		m[string(k)] = v
	}

	return m
}
