package util

import (
	"bytes"
)

func BytesCombine(pBytes ...[]byte) []byte {
	return bytes.Join(pBytes, []byte(""))
}

func BitSet(b interface{}, bit byte, set bool) {
	switch b.(type) {
	case *byte:
		if set {
			*b.(*byte) = *b.(*byte) | (byte(1) << bit)
		} else {
			*b.(*byte) = *b.(*byte) & (^(byte(1) << bit))
		}
	case *uint16:
		if set {
			*b.(*uint16) = *b.(*uint16) | (uint16(1) << bit)
		} else {
			*b.(*uint16) = *b.(*uint16) & (^(uint16(1) << bit))
		}
	case *uint32:
		if set {
			*b.(*uint32) = *b.(*uint32) | (uint32(1) << bit)
		} else {
			*b.(*uint32) = *b.(*uint32) & (^(uint32(1) << bit))
		}
	case *uint64:
		if set {
			*b.(*uint64) = *b.(*uint64) | (uint64(1) << bit)
		} else {
			*b.(*uint64) = *b.(*uint64) & (^(uint64(1) << bit))
		}
	case *uint:
		if set {
			*b.(*uint) = *b.(*uint) | (uint(1) << bit)
		} else {
			*b.(*uint) = *b.(*uint) & (^(uint(1) << bit))
		}
	default:
		b = nil
	}
}
