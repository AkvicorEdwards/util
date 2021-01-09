package util

import "encoding/binary"

func UInt16ToBytes(i uint16) [2]byte {
	var buf = [2]byte{}
	binary.BigEndian.PutUint16(buf[:], i)
	return buf
}

func BytesToUInt16(buf []byte) uint16 {
	return binary.BigEndian.Uint16(buf)
}

func UInt32ToBytes(i uint32) [4]byte {
	var buf = [4]byte{}
	binary.BigEndian.PutUint32(buf[:], i)
	return buf
}

func BytesToUInt32(buf []byte) uint32 {
	return binary.BigEndian.Uint32(buf)
}

func UInt64ToBytes(i uint64) [8]byte {
	var buf = [8]byte{}
	binary.BigEndian.PutUint64(buf[:], i)
	return buf
}

func BytesToUInt64(buf []byte) uint64 {
	return binary.BigEndian.Uint64(buf)
}

func UIntToBytes(i interface{}) []byte {
	switch i.(type) {
	case uint8:
		return []byte{i.(uint8)}
	case uint16:
		return func() []byte {
			d := UInt16ToBytes(i.(uint16))
			return d[:]
		}()
	case uint32:
		return func() []byte {
			d := UInt32ToBytes(i.(uint32))
			return d[:]
		}()
	case uint64:
		return func() []byte {
			d := UInt64ToBytes(i.(uint64))
			return d[:]
		}()
	}
	return nil
}

func BytesToUInt(buf []byte) uint64 {
	switch len(buf) {
	case 1:
		return uint64(buf[0])
	case 2:
		return uint64(BytesToUInt16(buf))
	case 4:
		return uint64(BytesToUInt32(buf))
	case 8:
		return BytesToUInt64(buf)
	default:
		val := uint64(0)
		for _, v := range buf {
			val <<= 8
			val = val + uint64(v)
		}
		return val
	}
}
