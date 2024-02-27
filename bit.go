package util

func BitSet(b any, bit byte, set bool) {
	switch data := b.(type) {
	case *byte: // uint8
		if set {
			*data = *data | (byte(1) << bit)
		} else {
			*data = *data & (^(byte(1) << bit))
		}
	case *int8:
		if set {
			*data = *data | (int8(1) << bit)
		} else {
			*data = *data & (^(int8(1) << bit))
		}
	case *uint16:
		if set {
			*data = *data | (uint16(1) << bit)
		} else {
			*data = *data & (^(uint16(1) << bit))
		}
	case *int16:
		if set {
			*data = *data | (int16(1) << bit)
		} else {
			*data = *data & (^(int16(1) << bit))
		}
	case *uint32:
		if set {
			*data = *data | (uint32(1) << bit)
		} else {
			*data = *data & (^(uint32(1) << bit))
		}
	case *int32:
		if set {
			*data = *data | (int32(1) << bit)
		} else {
			*data = *data & (^(int32(1) << bit))
		}
	case *uint64:
		if set {
			*data = *data | (uint64(1) << bit)
		} else {
			*data = *data & (^(uint64(1) << bit))
		}
	case *int64:
		if set {
			*data = *data | (int64(1) << bit)
		} else {
			*data = *data & (^(int64(1) << bit))
		}
	case *uint:
		if set {
			*data = *data | (uint(1) << bit)
		} else {
			*data = *data & (^(uint(1) << bit))
		}
	case *int:
		if set {
			*data = *data | (int(1) << bit)
		} else {
			*data = *data & (^(int(1) << bit))
		}
	}
}
