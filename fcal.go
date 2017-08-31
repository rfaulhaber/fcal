package main

import (
	"flag"
	"github.com/rfaulhaber/fdate"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	todayFlag := flag.Bool("t", false, "prints today's date instead of the month")
	flag.Parse()

	output := ""

	if *todayFlag {
		today := fdate.Today()

		output = today.String() + "\n"

		if today.Month() == 13 {
			output += " - " + fdate.CompDay(today.Day()).String()
		}
	} else {
		output = calendarBuilder(fdate.Today())
	}

	io.WriteString(os.Stdout, output)
}

func calendarBuilder(date fdate.Date) string {
	output := ""

	month := date.Month().String()
	year := date.RomanYear().String()

	monthString := month + " " + year
	spacesCount := 20 - len(monthString)/2

	monthSpaces := ""

	for i := 0; i < spacesCount; i++ {
		monthSpaces += " "
	}

	output += monthSpaces + monthString + "\n"

	var weekCount int

	if date.Month() < 13 {
		weekCount = 10
	} else {
		if date.IsLeapYear() {
			weekCount = 6
		} else {
			weekCount = 5
		}
	}

	for i := 0; i < weekCount; i++ {
		weekday := fdate.Weekday(i).String()
		abbr := strings.ToUpper(string(weekday[0])) + string(weekday[1:3])

		// "é" is two bytes, so we need to grab the next one
		if i == 9 {
			abbr += string(weekday[3])
		}

		output += abbr + " "
	}

	output += "\n"

	if date.Month() == 13 {

	} else {
		monthDays := 31

		for i := 1; i < monthDays; i++ {
			spaces := ""

			// TODO: do this in a cleaner way
			if i == 1 {
				spaces = "  "
			} else if i < 10 {
				spaces = "   "
			} else if i%10 == 1 {
				spaces = " "
			} else {
				spaces = "  "
			}

			dateStr := strconv.Itoa(i)

			// this doesn't work in the command prompt but works in Powershell on Windows
			if i == date.Day() {
				dateStr = highlightDate(dateStr)

				if strings.Contains(dateStr, " ") {
					spaces = spaces[:len(spaces)-1]
				}
			}

			output += spaces + dateStr

			if i%10 == 0 {
				output += "\n"
			}
		}
	}

	return output
}

func highlightDate(dateStr string) string {
	var str string

	if len(dateStr) == 1 {
		str = "\033[7m " + dateStr + "\033[27m"
	} else {
		str = "\033[7m" + dateStr + "\033[27m"
	}

	return str
}
