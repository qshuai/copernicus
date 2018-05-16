package block

import (
	"fmt"
	"io"

	"github.com/btcboost/copernicus/util"
)

type DiskBlockPos struct {
	File int
	Pos  int
}

type DiskTxPos struct {
	BlockIn    *DiskBlockPos
	TxOffsetIn int
}

// TxOffset is after header
var TxOffset int

func (diskBlockPos *DiskBlockPos) Serialize(writer io.Writer) error {
	err := util.WriteVarInt(writer, uint64(diskBlockPos.File))
	if err != nil {
		return err
	}
	return util.WriteVarInt(writer, uint64(diskBlockPos.Pos))
}

func (diskTxPos *DiskTxPos) SerializeDiskTxPos(writer io.Writer) error {
	err := diskTxPos.BlockIn.Serialize(writer)
	if err != nil {
		return err
	}
	return util.WriteVarInt(writer, uint64(diskTxPos.TxOffsetIn))
}

func DeserializeDiskBlock(reader io.Reader) (*DiskBlockPos, error) {
	file, err := util.ReadVarInt(reader)
	if err != nil {
		return nil, err
	}
	pos, err := util.ReadVarInt(reader)

	if err != nil {
		return nil, err
	}
	diskBlockPos := DiskBlockPos{File: int(file), Pos: int(pos)}
	return &diskBlockPos, nil
}

func (diskBlockPos *DiskBlockPos) SetNull() {
	diskBlockPos.File = -1
	diskBlockPos.Pos = 0
}

func (diskBlockPos *DiskBlockPos) Equal(other *DiskBlockPos) bool {
	return diskBlockPos.Pos == other.Pos && diskBlockPos.File == other.File

}

func (diskBlockPos *DiskBlockPos) IsNull() bool {
	return diskBlockPos.File == -1
}

func (diskBlockPos *DiskBlockPos) ToString() string {
	return fmt.Sprintf("BlcokDiskPos(File=%d, Pos=%d)", diskBlockPos.File, diskBlockPos.Pos)
}

func NewDiskBlockPos(file int, pos int) *DiskBlockPos {
	diskBlockPos := DiskBlockPos{File: file, Pos: pos}
	return &diskBlockPos
}

func NewDiskTxPos(blockIn *DiskBlockPos, offsetIn int) *DiskTxPos {
	diskTxPos := &DiskTxPos{
		BlockIn:    blockIn,
		TxOffsetIn: offsetIn,
	}
	return diskTxPos
}
func SetNull() {
	var blockPos *DiskBlockPos
	blockPos.SetNull()
	TxOffset = 0
}