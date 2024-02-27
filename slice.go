package util

type RemoveDuplicates struct{}

func NewRemoveDuplicates() *RemoveDuplicates {
	return &RemoveDuplicates{}
}

func (d *RemoveDuplicates) String(s []string) []string {
	result := make([]string, 0, len(s))
	temp := map[string]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) Byte(s []byte) []byte {
	result := make([]byte, 0, len(s))
	temp := map[byte]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) Int8(s []int8) []int8 {
	result := make([]int8, 0, len(s))
	temp := map[int8]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) Int16(s []int16) []int16 {
	result := make([]int16, 0, len(s))
	temp := map[int16]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) Int(s []int) []int {
	result := make([]int, 0, len(s))
	temp := map[int]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) Int32(s []int32) []int32 {
	result := make([]int32, 0, len(s))
	temp := map[int32]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) Int64(s []int64) []int64 {
	result := make([]int64, 0, len(s))
	temp := map[int64]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) UInt8(s []uint8) []uint8 {
	result := make([]uint8, 0, len(s))
	temp := map[uint8]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) UInt16(s []uint16) []uint16 {
	result := make([]uint16, 0, len(s))
	temp := map[uint16]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) UInt(s []uint) []uint {
	result := make([]uint, 0, len(s))
	temp := map[uint]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) UInt32(s []uint32) []uint32 {
	result := make([]uint32, 0, len(s))
	temp := map[uint32]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) UInt64(s []uint64) []uint64 {
	result := make([]uint64, 0, len(s))
	temp := map[uint64]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) Float32(s []float32) []float32 {
	result := make([]float32, 0, len(s))
	temp := map[float32]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}

func (d *RemoveDuplicates) Float64(s []float64) []float64 {
	result := make([]float64, 0, len(s))
	temp := map[float64]struct{}{}
	for _, item := range s {
		if _, ok := temp[item]; !ok {
			temp[item] = struct{}{}
			result = append(result, item)
		}
	}
	return result
}
