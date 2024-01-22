package util

import "encoding/binary"

func UInt16ToBytesSlice(i uint16) []byte {
  var buf = make([]byte, 2)
  binary.BigEndian.PutUint16(buf, i)
  return buf
}

func BytesSliceToUInt16(buf []byte) uint16 {
  return binary.BigEndian.Uint16(buf)
}

func UInt16ToBytesArray(i uint16) [2]byte {
  var buf = [2]byte{}
  binary.BigEndian.PutUint16(buf[:], i)
  return buf
}

func BytesArrayToUInt16(buf [2]byte) uint16 {
  return binary.BigEndian.Uint16(buf[:])
}

func UInt32ToBytesSlice(i uint32) []byte {
  var buf = make([]byte, 4)
  binary.BigEndian.PutUint32(buf, i)
  return buf
}

func BytesSliceToUInt32(buf []byte) uint32 {
  return binary.BigEndian.Uint32(buf)
}

func UInt32ToBytesArray(i uint32) [4]byte {
  var buf = [4]byte{}
  binary.BigEndian.PutUint32(buf[:], i)
  return buf
}

func BytesArrayToUInt32(buf [4]byte) uint32 {
  return binary.BigEndian.Uint32(buf[:])
}

func UInt64ToBytesSlice(i uint64) []byte {
  var buf = make([]byte, 8)
  binary.BigEndian.PutUint64(buf, i)
  return buf
}

func BytesSliceToUInt64(buf []byte) uint64 {
  return binary.BigEndian.Uint64(buf)
}

func UInt64ToBytesArray(i uint64) [8]byte {
  var buf = [8]byte{}
  binary.BigEndian.PutUint64(buf[:], i)
  return buf
}

func BytesArrayToUInt64(buf [8]byte) uint64 {
  return binary.BigEndian.Uint64(buf[:])
}

func UIntToBytes(i interface{}) []byte {
  switch i.(type) {
  case uint8:
    return []byte{i.(uint8)}
  case uint16:
    return func() []byte {
      d := UInt16ToBytesSlice(i.(uint16))
      return d[:]
    }()
  case uint32:
    return func() []byte {
      d := UInt32ToBytesSlice(i.(uint32))
      return d[:]
    }()
  case uint64:
    return func() []byte {
      d := UInt64ToBytesSlice(i.(uint64))
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
    return uint64(BytesSliceToUInt16(buf))
  case 4:
    return uint64(BytesSliceToUInt32(buf))
  case 8:
    return BytesSliceToUInt64(buf)
  default:
    val := uint64(0)
    for _, v := range buf {
      val <<= 8
      val = val + uint64(v)
    }
    return val
  }
}
