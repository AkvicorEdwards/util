package util

import (
	"fmt"
	"testing"
)

func TestSize(t *testing.T) {
	var sz Size
	var szt Size
	var szu uint64
	var szf float64
	var szs string
	var szsa string
	var sf *SizeFormatModel
	var sfs []*SizeFormatModel

	// Size
	sz = 123*SizeB + 234*SizeKiB + 345*SizeMiB + 456*SizeGiB + 567*SizeTiB + 678*SizePiB + 1*SizeEiB

	szs = sz.Format(",", true)
	szsa = "1EiB,678PiB,567TiB,456GiB,345MiB,234KiB,123B"
	if szs != szsa {
		t.Fatalf("szu value is wrong, need [%s] got [%s]", szsa, szs)
	}

	sfs = sz.FormatSlice(true)
	szs = fmt.Sprint(sfs)
	szsa = "[1EiB 678PiB 567TiB 456GiB 345MiB 234KiB 123B]"
	if szs != szsa {
		t.Fatalf("szu value is wrong, need [%s] got [%s]", szsa, szs)
	}

	sf = sz.HighestUnit(true)
	szs = sf.String()
	szsa = "1EiB"
	if szs != szsa {
		t.Fatalf("szu value is wrong, need [%s] got [%s]", szsa, szs)
	}
	szs = sf.StringFloat()
	szsa = "1.66EiB"
	if szs != szsa {
		t.Fatalf("szu value is wrong, need [%s] got [%s]", szsa, szs)
	}
	szs = sf.FormatFloat(0, 2, " ", false)
	szsa = "1.66 EiB"
	if szs != szsa {
		t.Fatalf("szu value is wrong, need [%s] got [%s]", szsa, szs)
	}

	szu = 1916905554527365243
	if sz != Size(szu) {
		t.Fatalf("sz value is wrong, need [%d] got [%d]", szu, sz)
	}

	// SizeFormatModel
	sf = sz.B()
	szf = 1916905554527365120.0
	if sf.Value != szu {
		t.Fatalf("szu value is wrong, need [%d] got [%d]", szu, sf.Value)
	}
	if sf.ValueFloat != szf {
		t.Fatalf("szf value is wrong, need [%f] got [%f]", szf, sf.ValueFloat)
	}

	szs = "1916905554527365243B"
	szsa = sf.String()
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365120.00B"
	szsa = sf.StringFloat()
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365243B"
	szsa = sf.Format(0, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365243 B"
	szsa = sf.Format(0, " ", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "    1916905554527365243B"
	szsa = sf.Format(24, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365243    B"
	szsa = sf.Format(-24, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "00001916905554527365243B"
	szsa = sf.Format(24, "", true)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365120.00B"
	szsa = sf.FormatFloat(0, 2, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365120.00 B"
	szsa = sf.FormatFloat(0, 2, " ", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = " 1916905554527365120.00B"
	szsa = sf.FormatFloat(24, 2, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365120.00 B"
	szsa = sf.FormatFloat(-24, 2, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "01916905554527365120.00B"
	szsa = sf.FormatFloat(24, 2, "", true)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365243B"
	szsa = sf.FormatValue(0, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365243 B"
	szsa = sf.FormatValue(0, " ", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "     1916905554527365243B"
	szsa = sf.FormatValue(24, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365243     B"
	szsa = sf.FormatValue(-24, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "000001916905554527365243B"
	szsa = sf.FormatValue(24, "", true)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365120.00B"
	szsa = sf.FormatValueFloat(0, 2, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365120.00 B"
	szsa = sf.FormatValueFloat(0, 2, " ", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "  1916905554527365120.00B"
	szsa = sf.FormatValueFloat(24, 2, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "1916905554527365120.00  B"
	szsa = sf.FormatValueFloat(-24, 2, "", false)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szs = "001916905554527365120.00B"
	szsa = sf.FormatValueFloat(24, 2, "", true)
	if szs != szsa {
		t.Fatalf("szs value is wrong, need [%s] got [%s]", szs, szsa)
	}

	szt = sf.Size()
	if szt != sz {
		t.Fatalf("szt value is wrong, need [%d] got [%d]", sz, szt)
	}

	szt = sf.SizeFloat()
	if float64(szt) != szf {
		t.Fatalf("szt value is wrong, need [%f] got [%f]", szf, float64(szt))
	}

	sf = sz.MiB()
	szs = sf.SingleSymbol().Symbol
	if szs != SizeMSymbol {
		t.Fatalf("szt value is wrong, need [%s] got [%s]", SizeMSymbol, szs)
	}
	szs = sf.Symbol
	if szs != SizeMiBSymbol {
		t.Fatalf("szt value is wrong, need [%s] got [%s]", SizeMSymbol, szs)
	}

}
