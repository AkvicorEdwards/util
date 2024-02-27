package util

import (
	"fmt"
	"time"
)

func ExampleTimeNowFormat() {
	t := TimeNowFormat()
	if len(t) != 19 {
		fmt.Println(t)
	}

	// Output:
	//
}

func ExampleTimeUnixToFormat() {
	t := TimeUnixToFormat(1705675065)
	fmt.Println(t)

	// Output:
	// 2024-01-19 22:37:45
}

func ExampleTimeNowToBase36() {
	s := TimeNowToBase36()
	if len(s) < 1 {
		fmt.Println(s)
	}
	s = TimeNowToBase36(8)
	if len(s) != 8 {
		fmt.Println(s)
	}

	// Output:
	//
}

func ExampleTimeUnixToBase36() {
	s := TimeUnixToBase36(1705675065)
	fmt.Println(s)
	s = TimeUnixToBase36(1705675065, 8)
	fmt.Println(s)

	// Output:
	// s7ijax
	// 00s7ijax
}

func ExampleTimeBase36ToUnix() {
	t := TimeBase36ToUnix("s7ijax")
	fmt.Println(t)
	t = TimeBase36ToUnix("00s7ijax")
	fmt.Println(t)

	// Output:
	// 1705675065
	// 1705675065
}

func ExampleYearBetweenTwoDate() {
	t1 := int64(1739929126)
	t2 := int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(YearBetweenTwoDate(time.Unix(t1, 0), time.Unix(t1, 0)))
	fmt.Println(YearBetweenTwoDate(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(YearBetweenTwoDate(time.Unix(t2, 0), time.Unix(t1, 0)))

	// Output:
	// 2025-02-19 09:38:46 -> 2024-01-18 22:37:45
	// 0
	// 1
	// -1
}

func ExampleYearBetweenTwoTime() {
	t1 := int64(1739929126)
	t2 := int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(YearBetweenTwoTime(time.Unix(t1, 0), time.Unix(t1, 0)))
	fmt.Println(YearBetweenTwoTime(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(YearBetweenTwoTime(time.Unix(t2, 0), time.Unix(t1, 0)))
	t1 = int64(1737211064)
	t2 = int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(YearBetweenTwoTime(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(YearBetweenTwoTime(time.Unix(t2, 0), time.Unix(t1, 0)))
	t1 = int64(1737211066)
	t2 = int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(YearBetweenTwoTime(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(YearBetweenTwoTime(time.Unix(t2, 0), time.Unix(t1, 0)))

	// Output:
	// 2025-02-19 09:38:46 -> 2024-01-18 22:37:45
	// 0
	// 1
	// -1
	// 2025-01-18 22:37:44 -> 2024-01-18 22:37:45
	// 0
	// 0
	// 2025-01-18 22:37:46 -> 2024-01-18 22:37:45
	// 1
	// -1
}

func ExampleMonthBetweenTwoDate() {
	t1 := int64(1739929126)
	t2 := int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(MonthBetweenTwoDate(time.Unix(t1, 0), time.Unix(t1, 0)))
	fmt.Println(MonthBetweenTwoDate(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(MonthBetweenTwoDate(time.Unix(t2, 0), time.Unix(t1, 0)))
	t1 = int64(1737211064)
	t2 = int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(MonthBetweenTwoDate(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(MonthBetweenTwoDate(time.Unix(t2, 0), time.Unix(t1, 0)))
	t1 = int64(1737211066)
	t2 = int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(MonthBetweenTwoDate(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(MonthBetweenTwoDate(time.Unix(t2, 0), time.Unix(t1, 0)))

	// Output:
	// 2025-02-19 09:38:46 -> 2024-01-18 22:37:45
	// 0
	// 13
	// -13
	// 2025-01-18 22:37:44 -> 2024-01-18 22:37:45
	// 12
	// -12
	// 2025-01-18 22:37:46 -> 2024-01-18 22:37:45
	// 12
	// -12
}

func ExampleMonthBetweenTwoTime() {
	t1 := int64(1739929126)
	t2 := int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(MonthBetweenTwoTime(time.Unix(t1, 0), time.Unix(t1, 0)))
	fmt.Println(MonthBetweenTwoTime(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(MonthBetweenTwoTime(time.Unix(t2, 0), time.Unix(t1, 0)))
	t1 = int64(1737211064)
	t2 = int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(MonthBetweenTwoTime(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(MonthBetweenTwoTime(time.Unix(t2, 0), time.Unix(t1, 0)))
	t1 = int64(1737211066)
	t2 = int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(MonthBetweenTwoTime(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(MonthBetweenTwoTime(time.Unix(t2, 0), time.Unix(t1, 0)))

	// Output:
	// 2025-02-19 09:38:46 -> 2024-01-18 22:37:45
	// 0
	// 13
	// -13
	// 2025-01-18 22:37:44 -> 2024-01-18 22:37:45
	// 11
	// -11
	// 2025-01-18 22:37:46 -> 2024-01-18 22:37:45
	// 12
	// -12
}

func ExampleDayBetweenTwoDate() {
	t1 := int64(1705628326)
	t2 := int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(DayBetweenTwoDate(time.Unix(t1, 0), time.Unix(t1, 0)))
	fmt.Println(DayBetweenTwoDate(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(DayBetweenTwoDate(time.Unix(t2, 0), time.Unix(t1, 0)))

	// Output:
	// 2024-01-19 09:38:46 -> 2024-01-18 22:37:45
	// 0
	// 1
	// -1
}

func ExampleDayBetweenTwoTime() {
	t1 := int64(1705628326)
	t2 := int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(DayBetweenTwoTime(time.Unix(t1, 0), time.Unix(t1, 0)))
	fmt.Println(DayBetweenTwoTime(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(DayBetweenTwoTime(time.Unix(t2, 0), time.Unix(t1, 0)))

	// Output:
	// 2024-01-19 09:38:46 -> 2024-01-18 22:37:45
	// 0
	// 0
	// 0
}

func ExampleHourBetweenTwoTime() {
	t1 := int64(1705628326)
	t2 := int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(HourBetweenTwoTime(time.Unix(t1, 0), time.Unix(t1, 0)))
	fmt.Println(HourBetweenTwoTime(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(HourBetweenTwoTime(time.Unix(t2, 0), time.Unix(t1, 0)))

	// Output:
	// 2024-01-19 09:38:46 -> 2024-01-18 22:37:45
	// 0
	// 11
	// -11
}

func ExampleMinuteBetweenTwoTime() {
	t1 := int64(1705628326)
	t2 := int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(MinuteBetweenTwoTime(time.Unix(t1, 0), time.Unix(t1, 0)))
	fmt.Println(MinuteBetweenTwoTime(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(MinuteBetweenTwoTime(time.Unix(t2, 0), time.Unix(t1, 0)))

	// Output:
	// 2024-01-19 09:38:46 -> 2024-01-18 22:37:45
	// 0
	// 661
	// -661
}

func ExampleSecondBetweenTwoTime() {
	t1 := int64(1705628326)
	t2 := int64(1705588665)
	fmt.Printf("%s -> %s\n", TimeUnixToFormat(t1), TimeUnixToFormat(t2))
	fmt.Println(SecondBetweenTwoTime(time.Unix(t1, 0), time.Unix(t1, 0)))
	fmt.Println(SecondBetweenTwoTime(time.Unix(t1, 0), time.Unix(t2, 0)))
	fmt.Println(SecondBetweenTwoTime(time.Unix(t2, 0), time.Unix(t1, 0)))

	// Output:
	// 2024-01-19 09:38:46 -> 2024-01-18 22:37:45
	// 0
	// 39661
	// -39661
}
