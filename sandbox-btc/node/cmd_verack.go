package node

import (
	"io"

	"github.com/asciiu/appa/sandbox-btc/protocol"
)

func (n Node) handleVerack(header *protocol.MessageHeader, conn io.ReadWriter) error {
	return nil
}
