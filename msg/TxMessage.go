package msg

import (
	"github.com/btccom/copernicus/model"
	"io"
)

type TxMessage struct {
	Tx model.Transaction
}

func (txMessage *TxMessage) BitcoinSerialize(w io.Writer, size uint32) error {
	return nil
}

func (txMessage *TxMessage) BitcoinParse(reader io.Reader, size uint32) error {
	return nil
}

func (txMessage *TxMessage) Command() string {
	return COMMAND_TX
}

func (txMessage *TxMessage) MaxPayloadLength(size uint32) uint32 {
	return 0
}
