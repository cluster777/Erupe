package network

import (
	"bytes"
	"encoding/binary"
)

const (
	// CryptPacketHeaderLength represents the byte-length of
	// an encrypted packet header.
	CryptPacketHeaderLength = 14
)

// CryptPacketHeader represents the parsed information of an encrypted packet header.
type CryptPacketHeader struct {
	Pf0                     byte
	KeyRotDelta             byte
	PacketNum               uint16
	DataSize                uint16
	PrevPacketCombinedCheck uint16
	Check0                  uint16
	Check1                  uint16
	Check2                  uint16
}

// NewCryptPacketHeader parses raw bytes into a CryptPacketHeader
func NewCryptPacketHeader(data []byte) (*CryptPacketHeader, error) {
	var c = CryptPacketHeader{}

	r := bytes.NewReader(data)

	var err error
	err = binary.Read(r, binary.BigEndian, &c.Pf0)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &c.KeyRotDelta)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &c.PacketNum)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &c.DataSize)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &c.PrevPacketCombinedCheck)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &c.Check0)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &c.Check1)
	if err != nil {
		return nil, err
	}
	err = binary.Read(r, binary.BigEndian, &c.Check2)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
