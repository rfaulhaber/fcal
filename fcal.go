package main

import (
	"github.com/rfaulhaber/fdate"
	"io"
	"os"
	"strconv"
	"strings"
    "flag"
)

func main() {
    todayFlag := flag.Bool("t", false, "prints today's date instead of the month")
    flag.Parse()

    output := ""

    if *todayFlag {
        output = fdate.Today().String() + "\n"
    } else {
        output = calendarBuilder(fdate.Today())
    }

	io.WriteString(os.Stdout, output)
}

func calendarBuilder(date fdate.Date) string {
	output := ""

	month := date.Month().String()
	year := date.RomanYear().String()

	monthString := month + " " + year;
	spacesCount := 20 - len(monthString) / 2

	monthSpaces := ""

	for i := 0; i < spacesCount; i ++ {
		monthSpaces += " "
	}

	output += monthSpaces + monthString + "\n"

	for i := 0; i < 10; i++ {
		weekday := fdate.Weekday(i).String()
		abbr := strings.ToUpper(string(weekday[0])) + string(weekday[1:3])

		// "é" is two bytes, so we need to grab the next one
		if i == 9 {
			abbr += string(weekday[4])
		}

		output += abbr + " "
	}

	output += "\n"

	var monthDays int

	if date.Month() < 13 {
		monthDays = 31
	} else {
		monthDays = 6
	}

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
			dateStr = "\033[7m" + dateStr + "\033[27m"
        }

		output += spaces + dateStr

		if i%10 == 0 {
			output += "\n"
		}
	}

    return output;
}
