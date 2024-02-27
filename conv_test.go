package util

import (
	"fmt"
)

func ExampleUInt16ToBytesSlice() {
	val := uint16(65535)
	bs := UInt16ToBytesSlice(val)
	fmt.Println(bs)

	val /= 2
	bs = UInt16ToBytesSlice(val)
	fmt.Println(bs)

	val -= 1
	bs = UInt16ToBytesSlice(val)
	fmt.Println(bs)
	// Output:
	// [255 255]
	// [127 255]
	// [127 254]
}

func ExampleBytesSliceToUInt16() {
	bs := []byte{255, 255}
	val := BytesSliceToUInt16(bs)
	fmt.Println(val)

	bs = []byte{127, 255}
	val = BytesSliceToUInt16(bs)
	fmt.Println(val)

	bs = []byte{127, 254}
	val = BytesSliceToUInt16(bs)
	fmt.Println(val)

	// Output:
	// 65535
	// 32767
	// 32766
}

func ExampleUInt16ToBytesArray() {
	val := uint16(65535)
	bs := UInt16ToBytesArray(val)
	fmt.Println(bs)

	val /= 2
	bs = UInt16ToBytesArray(val)
	fmt.Println(bs)

	val -= 1
	bs = UInt16ToBytesArray(val)
	fmt.Println(bs)
	// Output:
	// [255 255]
	// [127 255]
	// [127 254]
}

func ExampleBytesArrayToUInt16() {
	bs := [2]byte{255, 255}
	val := BytesArrayToUInt16(bs)
	fmt.Println(val)

	bs = [2]byte{127, 255}
	val = BytesArrayToUInt16(bs)
	fmt.Println(val)

	bs = [2]byte{127, 254}
	val = BytesArrayToUInt16(bs)
	fmt.Println(val)

	// Output:
	// 65535
	// 32767
	// 32766
}

func ExampleUInt32ToBytesSlice() {
	val := uint32(4294967295)
	bs := UInt32ToBytesSlice(val)
	fmt.Println(bs)

	val = uint32(2147483647)
	bs = UInt32ToBytesSlice(val)
	fmt.Println(bs)

	val = uint32(2147417854)
	bs = UInt32ToBytesSlice(val)
	fmt.Println(bs)
	// Output:
	// [255 255 255 255]
	// [127 255 255 255]
	// [127 254 254 254]
}

func ExampleBytesSliceToUInt32() {
	bs := []byte{255, 255, 255, 255}
	val := BytesSliceToUInt32(bs)
	fmt.Println(val)

	bs = []byte{127, 255, 255, 255}
	val = BytesSliceToUInt32(bs)
	fmt.Println(val)

	bs = []byte{127, 254, 254, 254}
	val = BytesSliceToUInt32(bs)
	fmt.Println(val)

	// Output:
	// 4294967295
	// 2147483647
	// 2147417854
}

func ExampleUInt32ToBytesArray() {
	val := uint32(4294967295)
	bs := UInt32ToBytesArray(val)
	fmt.Println(bs)

	val = uint32(2147483647)
	bs = UInt32ToBytesArray(val)
	fmt.Println(bs)

	val = uint32(2147417854)
	bs = UInt32ToBytesArray(val)
	fmt.Println(bs)
	// Output:
	// [255 255 255 255]
	// [127 255 255 255]
	// [127 254 254 254]
}

func ExampleBytesArrayToUInt32() {
	bs := [4]byte{255, 255, 255, 255}
	val := BytesArrayToUInt32(bs)
	fmt.Println(val)

	bs = [4]byte{127, 255, 255, 255}
	val = BytesArrayToUInt32(bs)
	fmt.Println(val)

	bs = [4]byte{127, 254, 254, 254}
	val = BytesArrayToUInt32(bs)
	fmt.Println(val)

	// Output:
	// 4294967295
	// 2147483647
	// 2147417854
}

func ExampleUInt64ToBytesSlice() {
	val := uint64(18446744073709551615)
	bs := UInt64ToBytesSlice(val)
	fmt.Println(bs)

	val = uint64(9223372036854775807)
	bs = UInt64ToBytesSlice(val)
	fmt.Println(bs)

	val = uint64(9223089458054627070)
	bs = UInt64ToBytesSlice(val)
	fmt.Println(bs)
	// Output:
	// [255 255 255 255 255 255 255 255]
	// [127 255 255 255 255 255 255 255]
	// [127 254 254 254 254 254 254 254]
}

func ExampleBytesSliceToUInt64() {
	bs := []byte{255, 255, 255, 255, 255, 255, 255, 255}
	val := BytesSliceToUInt64(bs)
	fmt.Println(val)

	bs = []byte{127, 255, 255, 255, 255, 255, 255, 255}
	val = BytesSliceToUInt64(bs)
	fmt.Println(val)

	bs = []byte{127, 254, 254, 254, 254, 254, 254, 254}
	val = BytesSliceToUInt64(bs)
	fmt.Println(val)

	// Output:
	// 18446744073709551615
	// 9223372036854775807
	// 9223089458054627070
}

func ExampleUInt64ToBytesArray() {
	val := uint64(18446744073709551615)
	bs := UInt64ToBytesArray(val)
	fmt.Println(bs)

	val = uint64(9223372036854775807)
	bs = UInt64ToBytesArray(val)
	fmt.Println(bs)

	val = uint64(9223089458054627070)
	bs = UInt64ToBytesArray(val)
	fmt.Println(bs)
	// Output:
	// [255 255 255 255 255 255 255 255]
	// [127 255 255 255 255 255 255 255]
	// [127 254 254 254 254 254 254 254]
}

func ExampleBytesArrayToUInt64() {
	bs := [8]byte{255, 255, 255, 255, 255, 255, 255, 255}
	val := BytesArrayToUInt64(bs)
	fmt.Println(val)

	bs = [8]byte{127, 255, 255, 255, 255, 255, 255, 255}
	val = BytesArrayToUInt64(bs)
	fmt.Println(val)

	bs = [8]byte{127, 254, 254, 254, 254, 254, 254, 254}
	val = BytesArrayToUInt64(bs)
	fmt.Println(val)

	// Output:
	// 18446744073709551615
	// 9223372036854775807
	// 9223089458054627070
}

func ExampleUIntToBytes() {
	val8 := uint8(254)
	bs := UIntToBytes(val8)
	fmt.Println(bs)

	val16 := uint16(65534)
	bs = UIntToBytes(val16)
	fmt.Println(bs)

	val32 := uint32(4294967294)
	bs = UIntToBytes(val32)
	fmt.Println(bs)

	val64 := uint64(18446744073709551614)
	bs = UIntToBytes(val64)
	fmt.Println(bs)

	valErr := ""
	bs = UIntToBytes(valErr)
	fmt.Println(bs)

	// Output:
	// [254]
	// [255 254]
	// [255 255 255 254]
	// [255 255 255 255 255 255 255 254]
	// []
}

func ExampleBytesToUInt() {
	bs := []byte{255}
	fmt.Println(BytesToUInt(bs))
	bs = []byte{255, 255}
	fmt.Println(BytesToUInt(bs))
	bs = []byte{255, 255, 255, 255}
	fmt.Println(BytesToUInt(bs))
	bs = []byte{255, 255, 255, 255, 255, 255, 255, 255}
	fmt.Println(BytesToUInt(bs))
	bs = []byte{255, 255, 255, 255, 255, 255, 255}
	fmt.Println(BytesToUInt(bs))
	// Output:
	// 255
	// 65535
	// 4294967295
	// 18446744073709551615
	// 72057594037927935
}
