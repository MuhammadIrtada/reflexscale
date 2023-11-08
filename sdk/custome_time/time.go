package custome_time

import (
	"fmt"
	"time"
)

func ConvertTimeFormat(date string) string {
	// Parsing tanggal ke dalam objek time.Time
	parsedDate, err := time.Parse("2006-01-02 15:04:05.999 -0700 -07", date)
	if err != nil {
		return ""
	}

	// Mengubah format tanggal ke "07-Nov-23" dengan bulan dalam bahasa Indonesia
	return fmt.Sprintf("%02d-%s-%02d", parsedDate.Day(), formatBulanIndonesia(parsedDate.Month()), parsedDate.Year()%100)
}

func formatBulanIndonesia(bulan time.Month) string {
	switch bulan {
	case time.January:
		return "Jan"
	case time.February:
		return "Feb"
	case time.March:
		return "Mar"
	case time.April:
		return "Apr"
	case time.May:
		return "Mei"
	case time.June:
		return "Jun"
	case time.July:
		return "Jul"
	case time.August:
		return "Ags"
	case time.September:
		return "Sep"
	case time.October:
		return "Okt"
	case time.November:
		return "Nov"
	case time.December:
		return "Des"
	default:
		return ""
	}
}