package util

import (
	"fmt"
	"strings"
)

type Size uint64

const SizeByte Size = 1
const SizeB Size = 1

// Decimal System
const (
	SizeKilobyte = 1000 * SizeByte
	SizeMegabyte = 1000 * SizeKilobyte
	SizeGigabyte = 1000 * SizeMegabyte
	SizeTerabyte = 1000 * SizeGigabyte
	SizePetabyte = 1000 * SizeTerabyte
	SizeExabyte  = 1000 * SizePetabyte
)

const (
	SizeKB = SizeKilobyte
	SizeMB = SizeMegabyte
	SizeGB = SizeGigabyte
	SizeTB = SizeTerabyte
	SizePB = SizePetabyte
	SizeEB = SizeExabyte
)

const (
	SizeKBSymbol = "KB"
	SizeMBSymbol = "MB"
	SizeGBSymbol = "GB"
	SizeTBSymbol = "TB"
	SizePBSymbol = "PB"
	SizeEBSymbol = "EB"
)

// Binary System
const (
	SizeKibibyte = 1024 * SizeByte
	SizeMebibyte = 1024 * SizeKibibyte
	SizeGigibyte = 1024 * SizeMebibyte
	SizeTebibyte = 1024 * SizeGigibyte
	SizePebibyte = 1024 * SizeTebibyte
	SizeExbibyte = 1024 * SizePebibyte
)

const (
	SizeKiB = SizeKibibyte
	SizeMiB = SizeMebibyte
	SizeGiB = SizeGigibyte
	SizeTiB = SizeTebibyte
	SizePiB = SizePebibyte
	SizeEiB = SizeExbibyte
)

const (
	SizeKiBSymbol = "KiB"
	SizeMiBSymbol = "MiB"
	SizeGiBSymbol = "GiB"
	SizeTiBSymbol = "TiB"
	SizePiBSymbol = "PiB"
	SizeEiBSymbol = "EiB"
)

const (
	SizeBSymbol = "B"
	SizeKSymbol = "K"
	SizeMSymbol = "M"
	SizeGSymbol = "G"
	SizeTSymbol = "T"
	SizePSymbol = "P"
	SizeESymbol = "E"
)

// SizeFormatModel save calculation result with uint64 and float64
type SizeFormatModel struct {
	// size value in uint64
	Value uint64
	// size value in float64
	ValueFloat float64
	// size unit symbol
	Symbol string
}

func (f *SizeFormatModel) String() string {
	return fmt.Sprintf("%d%s", f.Value, f.Symbol)
}

func (f *SizeFormatModel) StringFloat() string {
	return fmt.Sprintf("%.2f%s", f.ValueFloat, f.Symbol)
}

// Format return formatted string of specified length
//
//	use sep to separate number and symbol
//	use zeroPadding to determine whether to pad the number with 0
func (f *SizeFormatModel) Format(length int, sep string, zeroPadding bool) string {
	if length >= 0 {
		length -= len(f.Symbol)
	} else {
		length += len(f.Symbol)
	}
	if zeroPadding {
		return fmt.Sprintf("%0*d%s%s", length, f.Value, sep, f.Symbol)
	} else {
		return fmt.Sprintf("%*d%s%s", length, f.Value, sep, f.Symbol)
	}
}

// FormatFloat return formatted string of specified length
//
//	use precision to preserve number precision
//	use sep to separate number and symbol
//	use zeroPadding to determine whether to pad the number with 0
func (f *SizeFormatModel) FormatFloat(length int, precision int, sep string, zeroPadding bool) string {
	if length >= 0 {
		length -= len(f.Symbol)
	} else {
		length += len(f.Symbol)
	}
	if zeroPadding {
		return fmt.Sprintf("%0*.*f%s%s", length, precision, f.ValueFloat, sep, f.Symbol)
	} else {
		return fmt.Sprintf("%*.*f%s%s", length, precision, f.ValueFloat, sep, f.Symbol)
	}
}

// FormatValue return formatted string of specified length of value
//
//	use sep to separate number and symbol
//	use zeroPadding to determine whether to pad the number with 0
func (f *SizeFormatModel) FormatValue(valueLength int, sep string, zeroPadding bool) string {
	if zeroPadding {
		return fmt.Sprintf("%0*d%s%s", valueLength, f.Value, sep, f.Symbol)
	} else {
		return fmt.Sprintf("%*d%s%s", valueLength, f.Value, sep, f.Symbol)
	}
}

// FormatValueFloat return formatted string of specified length of value
//
//	use precision to preserve number precision
//	use sep to separate number and symbol
//	use zeroPadding to determine whether to pad the number with 0
func (f *SizeFormatModel) FormatValueFloat(valueLength int, precision int, sep string, zeroPadding bool) string {
	if zeroPadding {
		return fmt.Sprintf("%0*.*f%s%s", valueLength, precision, f.ValueFloat, sep, f.Symbol)
	} else {
		return fmt.Sprintf("%*.*f%s%s", valueLength, precision, f.ValueFloat, sep, f.Symbol)
	}
}

// Size calculate original Size
func (f *SizeFormatModel) Size() Size {
	switch f.Symbol {
	case SizeBSymbol:
		return Size(f.Value) * SizeB
	case SizeKiBSymbol, SizeKSymbol:
		return Size(f.Value) * SizeKiB
	case SizeKBSymbol:
		return Size(f.Value) * SizeKB
	case SizeMiBSymbol, SizeMSymbol:
		return Size(f.Value) * SizeMiB
	case SizeMBSymbol:
		return Size(f.Value) * SizeMB
	case SizeGiBSymbol, SizeGSymbol:
		return Size(f.Value) * SizeGiB
	case SizeGBSymbol:
		return Size(f.Value) * SizeGB
	case SizeTiBSymbol, SizeTSymbol:
		return Size(f.Value) * SizeTiB
	case SizeTBSymbol:
		return Size(f.Value) * SizeTB
	case SizePiBSymbol, SizePSymbol:
		return Size(f.Value) * SizePiB
	case SizePBSymbol:
		return Size(f.Value) * SizePB
	case SizeEiBSymbol, SizeESymbol:
		return Size(f.Value) * SizeEiB
	case SizeEBSymbol:
		return Size(f.Value) * SizeEB
	default:
		return Size(f.Value) * SizeByte
	}
}

// SizeFloat calculate original Size
//
//	due to precision, it is not guaranteed to restore the original Size
func (f *SizeFormatModel) SizeFloat() Size {
	switch f.Symbol {
	case SizeBSymbol:
		return Size(f.ValueFloat * float64(SizeB))
	case SizeKiBSymbol, SizeKSymbol:
		return Size(f.ValueFloat * float64(SizeKiB))
	case SizeKBSymbol:
		return Size(f.ValueFloat * float64(SizeKB))
	case SizeMiBSymbol, SizeMSymbol:
		return Size(f.ValueFloat * float64(SizeMiB))
	case SizeMBSymbol:
		return Size(f.ValueFloat * float64(SizeMB))
	case SizeGiBSymbol, SizeGSymbol:
		return Size(f.ValueFloat * float64(SizeGiB))
	case SizeGBSymbol:
		return Size(f.ValueFloat * float64(SizeGB))
	case SizeTiBSymbol, SizeTSymbol:
		return Size(f.ValueFloat * float64(SizeTiB))
	case SizeTBSymbol:
		return Size(f.ValueFloat * float64(SizeTB))
	case SizePiBSymbol, SizePSymbol:
		return Size(f.ValueFloat * float64(SizePiB))
	case SizePBSymbol:
		return Size(f.ValueFloat * float64(SizePB))
	case SizeEiBSymbol, SizeESymbol:
		return Size(f.ValueFloat * float64(SizeEiB))
	case SizeEBSymbol:
		return Size(f.ValueFloat * float64(SizeEB))
	default:
		return Size(f.ValueFloat * float64(SizeB))
	}
}

// SingleSymbol convert symbol to single character
func (f *SizeFormatModel) SingleSymbol() *SizeFormatModel {
	res := &SizeFormatModel{Value: f.Value, ValueFloat: f.ValueFloat}
	switch f.Symbol {
	case SizeBSymbol:
		res.Symbol = SizeBSymbol
	case SizeKiBSymbol, SizeKBSymbol:
		res.Symbol = SizeKSymbol
	case SizeMiBSymbol, SizeMBSymbol:
		res.Symbol = SizeMSymbol
	case SizeGiBSymbol, SizeGBSymbol:
		res.Symbol = SizeGSymbol
	case SizeTiBSymbol, SizeTBSymbol:
		res.Symbol = SizeTSymbol
	case SizePiBSymbol, SizePBSymbol:
		res.Symbol = SizePSymbol
	case SizeEiBSymbol, SizeEBSymbol:
		res.Symbol = SizeESymbol
	default:
		res.Symbol = SizeBSymbol
	}
	return res
}

// B return size in B
func (s Size) B() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s),
		ValueFloat: float64(s),
		Symbol:     SizeBSymbol,
	}
}

// KB return size in KB
func (s Size) KB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeKilobyte),
		ValueFloat: float64(s) / float64(SizeKilobyte),
		Symbol:     SizeKBSymbol,
	}
}

// MB return size in MB
func (s Size) MB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeMegabyte),
		ValueFloat: float64(s) / float64(SizeMegabyte),
		Symbol:     SizeMBSymbol,
	}
}

// GB return size in GB
func (s Size) GB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeGigabyte),
		ValueFloat: float64(s) / float64(SizeGigabyte),
		Symbol:     SizeGBSymbol,
	}
}

// TB return size in TB
func (s Size) TB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeTerabyte),
		ValueFloat: float64(s) / float64(SizeTerabyte),
		Symbol:     SizeTBSymbol,
	}
}

// PB return size in PB
func (s Size) PB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizePetabyte),
		ValueFloat: float64(s) / float64(SizePetabyte),
		Symbol:     SizePBSymbol,
	}
}

// EB return size in EB
func (s Size) EB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeExabyte),
		ValueFloat: float64(s) / float64(SizeExabyte),
		Symbol:     SizeEBSymbol,
	}
}

// KiB return size in KiB
func (s Size) KiB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeKibibyte),
		ValueFloat: float64(s) / float64(SizeKibibyte),
		Symbol:     SizeKiBSymbol,
	}
}

// MiB return size in MiB
func (s Size) MiB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeMebibyte),
		ValueFloat: float64(s) / float64(SizeMebibyte),
		Symbol:     SizeMiBSymbol,
	}
}

// GiB return size in GiB
func (s Size) GiB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeGigibyte),
		ValueFloat: float64(s) / float64(SizeGigibyte),
		Symbol:     SizeGiBSymbol,
	}
}

// TiB return size in TiB
func (s Size) TiB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeTebibyte),
		ValueFloat: float64(s) / float64(SizeTebibyte),
		Symbol:     SizeTiBSymbol,
	}
}

// PiB return size in PiB
func (s Size) PiB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizePebibyte),
		ValueFloat: float64(s) / float64(SizePebibyte),
		Symbol:     SizePiBSymbol,
	}
}

// EiB return size in EiB
func (s Size) EiB() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeExbibyte),
		ValueFloat: float64(s) / float64(SizeExbibyte),
		Symbol:     SizeEiBSymbol,
	}
}

// K return size in K
func (s Size) K() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeKibibyte),
		ValueFloat: float64(s) / float64(SizeKibibyte),
		Symbol:     SizeKSymbol,
	}
}

// M return size in M
func (s Size) M() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeMebibyte),
		ValueFloat: float64(s) / float64(SizeMebibyte),
		Symbol:     SizeMSymbol,
	}
}

// G return size in G
func (s Size) G() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeGigibyte),
		ValueFloat: float64(s) / float64(SizeGigibyte),
		Symbol:     SizeGSymbol,
	}
}

// T return size in T
func (s Size) T() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeTebibyte),
		ValueFloat: float64(s) / float64(SizeTebibyte),
		Symbol:     SizeTSymbol,
	}
}

// P return size in P
func (s Size) P() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizePebibyte),
		ValueFloat: float64(s) / float64(SizePebibyte),
		Symbol:     SizePSymbol,
	}
}

// E return size in E
func (s Size) E() *SizeFormatModel {
	return &SizeFormatModel{
		Value:      uint64(s / SizeExbibyte),
		ValueFloat: float64(s) / float64(SizeExbibyte),
		Symbol:     SizeESymbol,
	}
}

func (s Size) String() string {
	return s.Format(" ", true)
}

// Format size
//
//	use sep to separate each unit
func (s Size) Format(sep string, binaryPrefixes bool) string {
	res := strings.Builder{}
	isFirst := true
	wres := func(s string) {
		if isFirst {
			res.WriteString(s)
			isFirst = false
		} else {
			res.WriteString(sep)
			res.WriteString(s)
		}
	}
	if binaryPrefixes {
		if s >= SizeEiB {
			wres(s.EiB().String())
			s = s % SizeEiB
		}
		if s >= SizePiB {
			wres(s.PiB().String())
			s = s % SizePiB
		}
		if s >= SizeTiB {
			wres(s.TiB().String())
			s = s % SizeTiB
		}
		if s >= SizeGiB {
			wres(s.GiB().String())
			s = s % SizeGiB
		}
		if s >= SizeMiB {
			wres(s.MiB().String())
			s = s % SizeMiB
		}
		if s >= SizeKiB {
			wres(s.KiB().String())
			s = s % SizeKiB
		}
	} else {
		if s >= SizeEB {
			wres(s.EB().String())
			s = s % SizeEB
		}
		if s >= SizePB {
			wres(s.PB().String())
			s = s % SizePB
		}
		if s >= SizeTB {
			wres(s.TB().String())
			s = s % SizeTB
		}
		if s >= SizeGB {
			wres(s.GB().String())
			s = s % SizeGB
		}
		if s >= SizeMB {
			wres(s.MB().String())
			s = s % SizeMB
		}
		if s >= SizeKB {
			wres(s.KB().String())
			s = s % SizeKB
		}
	}
	if s >= SizeB {
		wres(s.B().String())
	}
	return res.String()
}

// FormatSlice format size
//
//	each unit is stored in a separate SizeFormatModel
func (s Size) FormatSlice(binaryPrefixes bool) []*SizeFormatModel {
	t := s
	res := make([]*SizeFormatModel, 0, 7)
	if binaryPrefixes {
		if t >= SizeEiB {
			res = append(res, t.EiB())
			t = t % SizeEiB
		}
		if t >= SizePiB {
			res = append(res, t.PiB())
			t = t % SizePiB
		}
		if t >= SizeTiB {
			res = append(res, t.TiB())
			t = t % SizeTiB
		}
		if t >= SizeGiB {
			res = append(res, t.GiB())
			t = t % SizeGiB
		}
		if t >= SizeMiB {
			res = append(res, t.MiB())
			t = t % SizeMiB
		}
		if t >= SizeKiB {
			res = append(res, t.KiB())
			t = t % SizeKiB
		}
	} else {
		if t >= SizeEB {
			res = append(res, t.EB())
			t = t % SizeEB
		}
		if t >= SizePB {
			res = append(res, t.PB())
			t = t % SizePB
		}
		if t >= SizeTB {
			res = append(res, t.TB())
			t = t % SizeTB
		}
		if t >= SizeGB {
			res = append(res, t.GB())
			t = t % SizeGB
		}
		if t >= SizeMB {
			res = append(res, t.MB())
			t = t % SizeMB
		}
		if t >= SizeKB {
			res = append(res, t.KB())
			t = t % SizeKB
		}
	}
	if t >= SizeB {
		res = append(res, t.B())
	}
	return res
}

// HighestUnit return the highest unit value
func (s Size) HighestUnit(binaryPrefixes bool) *SizeFormatModel {
	if binaryPrefixes {
		if s >= SizeEiB {
			return s.EiB()
		}
		if s >= SizePiB {
			return s.PiB()
		}
		if s >= SizeTiB {
			return s.TiB()
		}
		if s >= SizeGiB {
			return s.GiB()
		}
		if s >= SizeMiB {
			return s.MiB()
		}
		if s >= SizeKiB {
			return s.KiB()
		}
	} else {
		if s >= SizeEB {
			return s.EB()
		}
		if s >= SizePB {
			return s.PB()
		}
		if s >= SizeTB {
			return s.TB()
		}
		if s >= SizeGB {
			return s.GB()
		}
		if s >= SizeMB {
			return s.MB()
		}
		if s >= SizeKB {
			return s.KB()
		}
	}
	return s.B()
}
