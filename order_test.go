package netlink

import (
	"encoding/binary"
	"testing"
)

const bitsInByte = 8

func TestToNetworkOrder32(t *testing.T) {
	totalBytesInUint32 := 4
	n := uint32(0x0a0b0c0d)
	converted := ToNetworkOrder32(n)

	if native == binary.BigEndian {
		if converted != n {
			t.Errorf("ToNetworkOrder32(%x) != %x in the bigendian system.", n, converted)
		}
	} else {
		originalBytes := make([]byte, totalBytesInUint32)
		native.PutUint32(originalBytes, n)
		convertedBytes := make([]byte, totalBytesInUint32)
		native.PutUint32(convertedBytes, converted)
		for i := 0; i < totalBytesInUint32; i++ {
			if originalBytes[i] != convertedBytes[totalBytesInUint32-1-i] {
				t.Errorf("The byte of index %d of %x is not equal to the byte of the index %d of the converted value %x.",
					i, originalBytes, totalBytesInUint32-1-i, convertedBytes)
			}
		}

		reflected := ToNetworkOrder32(converted)
		if reflected != n {
			t.Errorf("ToNetworkOrder32(ToNetworkOrder32(%x)) != %x", n, n)
		}
	}
}

func TestToNetworkOrder16(t *testing.T) {
	totalBytesInUint16 := 2
	n := uint16(0x0a0b)
	converted := ToNetworkOrder16(n)

	if native == binary.BigEndian {
		if converted != n {
			t.Errorf("ToNetworkOrder16(%x) != %x in the bigendian system.", n, converted)
		}
	} else {
		originalBytes := make([]byte, totalBytesInUint16)
		native.PutUint16(originalBytes, n)
		convertedBytes := make([]byte, totalBytesInUint16)
		native.PutUint16(convertedBytes, converted)
		for i := 0; i < totalBytesInUint16; i++ {
			if originalBytes[i] != convertedBytes[totalBytesInUint16-1-i] {
				t.Errorf("The byte of index %d of %x is not equal to the byte of the index %d of the converted value %x.",
					i, originalBytes, totalBytesInUint16-1-i, convertedBytes)
			}
		}

		reflected := ToNetworkOrder16(converted)
		if reflected != n {
			t.Errorf("ToNetworkOrder16(ToNetworkOrder16(%x)) != %x", n, n)
		}
	}
}
