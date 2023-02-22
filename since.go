package main

import (
	"fmt"
	"os"
	"time"
)

// Since handles since START_DATE [END_DATE|TODAY]
func Since(args []string) {
	MaxArgs(2, args)
	if len(args) == 0 {
		fmt.Fprintf(os.Stderr, "Need at least one YYYY-MM-DD date")
		os.Exit(1)
	}
	startDate, err := time.Parse("2006-01-02", args[0])
	if err != nil {
		fmt.Fprintf(os.Stderr, "Argument %s does not look like YYYY-MM-DD or a valid date", args[0])
		os.Exit(1)
	}
	endDate := time.Now()
	if len(args) == 2 {
		endDate, err = time.Parse("2006-01-02", args[1])
		if err != nil {
			fmt.Fprintf(os.Stderr, "Argument %s does not look like YYYY-MM-DD or a valid date", args[1])
			os.Exit(1)
		}
	}
	if endDate.Before(startDate) {
		startDate, endDate = endDate, startDate
	}
	days := endDate.Sub(startDate).Hours() / 24
	dYears, dMonths, dDays := diffDates(startDate, endDate)
	fmt.Printf("%d years, %d months, %d days / %d days between %s and %s",
		dYears, dMonths, dDays,
		int(days), startDate.Format("2006-01-02"), endDate.Format("2006-01-02"))
	if days >= 7 {
		fmt.Printf("\n= %.2f weeks", days/7)
	}
	if days >= 30 {
		fmt.Printf("\n= %.2f 30-day months", days/30)
	}
	if days >= 365 {
		fmt.Printf("\n= %.2f 365-day years", days/365)
	}
	fmt.Printf("\n")
}

func diffDates(a, b time.Time) (years, months, days int) {
	a = a.In(time.UTC)
	b = b.In(time.UTC)
	if a.After(b) {
		a, b = b, a
	}
	y1, M1, d1 := a.Date()
	y2, M2, d2 := b.Date()

	years = int(y2 - y1)
	months = int(M2 - M1)
	days = int(d2 - d1)

	if days < 0 {
		// days in month:
		t := time.Date(y1, M1, 32, 0, 0, 0, 0, time.UTC)
		days += 32 - t.Day()
		months--
	}
	if months < 0 {
		months += 12
		years--
	}

	return
}
