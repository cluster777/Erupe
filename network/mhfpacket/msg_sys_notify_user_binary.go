package mhfpacket

import (
	"errors"

	"github.com/Andoryuuta/Erupe/network"
	"github.com/Andoryuuta/Erupe/network/clientctx"
	"github.com/Andoryuuta/byteframe"
)

// MsgSysNotifyUserBinary represents the MSG_SYS_NOTIFY_USER_BINARY
type MsgSysNotifyUserBinary struct {
	CharID     uint32
	BinaryType uint8
}

// Opcode returns the ID associated with this packet type.
func (m *MsgSysNotifyUserBinary) Opcode() network.PacketID {
	return network.MSG_SYS_NOTIFY_USER_BINARY
}

// Parse parses the packet from binary
func (m *MsgSysNotifyUserBinary) Parse(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	return errors.New("Not implemented")
}

// Build builds a binary packet from the current data.
func (m *MsgSysNotifyUserBinary) Build(bf *byteframe.ByteFrame, ctx *clientctx.ClientContext) error {
	bf.WriteUint32(m.CharID)
	bf.WriteUint8(m.BinaryType)
	return nil
}
