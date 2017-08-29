package model

import (
	"bytes"
	"testing"
)

var rawByte = [80]byte{
	0x01, 0x00, 0x00, 0x00, // Version 1
	0x6f, 0xe2, 0x8c, 0x0a, 0xb6, 0xf1, 0xb3, 0x72,
	0xc1, 0xa6, 0xa2, 0x46, 0xae, 0x63, 0xf7, 0x4f,
	0x93, 0x1e, 0x83, 0x65, 0xe1, 0x5a, 0x08, 0x9c,
	0x68, 0xd6, 0x19, 0x00, 0x00, 0x00, 0x00, 0x00, // PrevBlock
	0x3b, 0xa3, 0xed, 0xfd, 0x7a, 0x7b, 0x12, 0xb2,
	0x7a, 0xc7, 0x2c, 0x3e, 0x67, 0x76, 0x8f, 0x61,
	0x7f, 0xc8, 0x1b, 0xc3, 0x88, 0x8a, 0x51, 0x32,
	0x3a, 0x9f, 0xb8, 0xaa, 0x4b, 0x1e, 0x5e, 0x4a, // MerkleRoot
	0x29, 0xab, 0x5f, 0x49, // Timestamp
	0xff, 0xff, 0x00, 0x1d, // Bits
	0xf3, 0xe0, 0x01, 0x00, // Nonce
}

func TestParseBlock(t *testing.T) {

	blockHeadFirst, err := ParseBlock(rawByte[:])
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(blockHeadFirst.Raw, rawByte[:]) {
		t.Error("The two slice should be equal ")
	}

	if blockHeadFirst.Version != 1 {
		t.Error("err : the version should be 1")
	}
	if blockHeadFirst.Size != uint32(len(rawByte)) {
		t.Error("The size should be equal")
	}
}
