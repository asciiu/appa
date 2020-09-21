package node

import (
	"io"

	"github.com/asciiu/appa/sandbox-btc/binary"
	"github.com/asciiu/appa/sandbox-btc/protocol"
	"github.com/sirupsen/logrus"
)

func (no Node) handleTx(header *protocol.MessageHeader, conn io.ReadWriter) error {
	var tx protocol.MsgTx

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&tx); err != nil {
		return err
	}

	logrus.Debugf("transaction: %+v", tx)

	return nil
}
