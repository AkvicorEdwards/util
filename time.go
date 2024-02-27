package util

import (
	"fmt"
	"strconv"
	"time"
)

const TimeFormatLayout = "2006-01-02 15:04:05"

func TimeNowFormat() string {
	return time.Now().Format(TimeFormatLayout)
}

func TimeNowToBase36(length ...int) string {
	if len(length) > 0 {
		return fmt.Sprintf("%0*s", length[0], strconv.FormatInt(time.Now().Unix(), 36))
	}
	return strconv.FormatInt(time.Now().Unix(), 36)
}

func TimeUnixToFormat(unix int64) string {
	return time.Unix(unix, 0).Format(TimeFormatLayout)
}

func TimeUnixToBase36(unix int64, length ...int) string {
	if len(length) > 0 {
		return fmt.Sprintf("%0*s", length[0], strconv.FormatInt(unix, 36))
	}
	return strconv.FormatInt(unix, 36)
}

func TimeBase36ToUnix(t string) int64 {
	date, err := strconv.ParseInt(t, 36, 64)
	if err != nil {
		return 0
	}
	return date
}

func YearBetweenTwoDate(t1, t2 time.Time) int64 {
	return int64(t1.Year() - t2.Year())
}

func YearBetweenTwoTime(t1, t2 time.Time) int64 {
	f := int64(1)
	if t1.Before(t2) {
		f = -1
		t1, t2 = t2, t1
	}
	y := int64(t1.Year() - t2.Year())
	t1 = t1.AddDate(-t1.Year(), 0, 0)
	t2 = t2.AddDate(-t2.Year(), 0, 0)
	if t1.Before(t2) {
		return f * (y - 1)
	}
	return f * y
}

func MonthBetweenTwoDate(t1, t2 time.Time) int64 {
	return 12*int64(t1.Year()-t2.Year()) + int64(t1.Month()-t2.Month())
}

func MonthBetweenTwoTime(t1, t2 time.Time) int64 {
	f := int64(1)
	if t1.Before(t2) {
		f = -1
		t1, t2 = t2, t1
	}
	m := 12*int64(t1.Year()-t2.Year()) + int64(t1.Month()-t2.Month())
	t1 = t1.AddDate(-t1.Year(), -int(t1.Month()-1), 0)
	t2 = t2.AddDate(-t2.Year(), -int(t2.Month()-1), 0)
	if t1.Before(t2) {
		return f * (m - 1)
	}
	return f * m
}

func DayBetweenTwoDate(t1, t2 time.Time) int64 {
	t1 = time.Date(t1.Year(), t1.Month(), t1.Day(), 0, 0, 0, 0, time.Local)
	t2 = time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, time.Local)
	return int64(t1.Sub(t2).Hours() / 24)
}

func DayBetweenTwoTime(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Hours() / 24)
}

func HourBetweenTwoTime(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Hours())
}

func MinuteBetweenTwoTime(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Minutes())
}

func SecondBetweenTwoTime(t1, t2 time.Time) int64 {
	return int64(t1.Sub(t2).Seconds())
}
