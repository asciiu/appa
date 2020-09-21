package node

import (
	"io"

	"github.com/asciiu/appa/sandbox-btc/binary"
	"github.com/asciiu/appa/sandbox-btc/protocol"
)

func (n Node) handlePong(header *protocol.MessageHeader, conn io.ReadWriter) error {
	var pong protocol.MsgPing

	lr := io.LimitReader(conn, int64(header.Length))
	if err := binary.NewDecoder(lr).Decode(&pong); err != nil {
		return err
	}

	n.PongCh <- pong.Nonce

	return nil
}
