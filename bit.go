package util

func BitSet(b interface{}, bit byte, set bool) {
  switch b.(type) {
  case *byte: // uint8
    if set {
      *b.(*byte) = *b.(*byte) | (byte(1) << bit)
    } else {
      *b.(*byte) = *b.(*byte) & (^(byte(1) << bit))
    }
  case *int8:
    if set {
      *b.(*int8) = *b.(*int8) | (int8(1) << bit)
    } else {
      *b.(*int8) = *b.(*int8) & (^(int8(1) << bit))
    }
  case *uint16:
    if set {
      *b.(*uint16) = *b.(*uint16) | (uint16(1) << bit)
    } else {
      *b.(*uint16) = *b.(*uint16) & (^(uint16(1) << bit))
    }
  case *int16:
    if set {
      *b.(*int16) = *b.(*int16) | (int16(1) << bit)
    } else {
      *b.(*int16) = *b.(*int16) & (^(int16(1) << bit))
    }
  case *uint32:
    if set {
      *b.(*uint32) = *b.(*uint32) | (uint32(1) << bit)
    } else {
      *b.(*uint32) = *b.(*uint32) & (^(uint32(1) << bit))
    }
  case *int32:
    if set {
      *b.(*int32) = *b.(*int32) | (int32(1) << bit)
    } else {
      *b.(*int32) = *b.(*int32) & (^(int32(1) << bit))
    }
  case *uint64:
    if set {
      *b.(*uint64) = *b.(*uint64) | (uint64(1) << bit)
    } else {
      *b.(*uint64) = *b.(*uint64) & (^(uint64(1) << bit))
    }
  case *int64:
    if set {
      *b.(*int64) = *b.(*int64) | (int64(1) << bit)
    } else {
      *b.(*int64) = *b.(*int64) & (^(int64(1) << bit))
    }
  case *uint:
    if set {
      *b.(*uint) = *b.(*uint) | (uint(1) << bit)
    } else {
      *b.(*uint) = *b.(*uint) & (^(uint(1) << bit))
    }
  case *int:
    if set {
      *b.(*int) = *b.(*int) | (int(1) << bit)
    } else {
      *b.(*int) = *b.(*int) & (^(int(1) << bit))
    }
  default:
    b = nil
  }
}
