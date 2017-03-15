package netlink

import (
	"encoding/binary"

	"github.com/vishvananda/netlink/nl"
)

var (
	native       = nl.NativeEndian()
	networkOrder = binary.BigEndian
)

func htonl(val uint32) []byte {
	bytes := make([]byte, 4)
	binary.BigEndian.PutUint32(bytes, val)
	return bytes
}

func htons(val uint16) []byte {
	bytes := make([]byte, 2)
	binary.BigEndian.PutUint16(bytes, val)
	return bytes
}

func htohl(val uint32) []byte {
	bytes := make([]byte, 4)
	native.PutUint32(bytes, val)
	return bytes
}

func htohs(val uint16) []byte {
	bytes := make([]byte, 2)
	native.PutUint16(bytes, val)
	return bytes
}

func ntohl(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}

func ntohs(buf []byte) uint16 {
	return binary.BigEndian.Uint16(buf)
}

// Convert a 32 bit value to the network byte order
func ToNetworkOrder32(n uint32) uint32 {
	if native == networkOrder {
		return n
	} else {
		return ntohl(htohl(n))
	}
}

// Convert a 16 bit value to the network byte order
func ToNetworkOrder16(n uint16) uint16 {
	if native == networkOrder {
		return n
	} else {
		return ntohs(htohs(n))
	}
}
