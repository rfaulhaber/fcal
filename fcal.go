package main

import (
	"github.com/rfaulhaber/fdate"
	"strings"
	"strconv"
	"io"
	"os"
)

func main() {
	today := fdate.Today()

	output := ""

	month := today.Month().String()
	year := today.RomanYear().String()

	monthSpaces := "            "

	output += monthSpaces + month + " " + year + "\n"

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

	for i := 1; i < 31; i++ {
		spaces := ""

		// TODO: do this in a cleaner way
		if i == 1 {
			spaces = "  "
		} else if i < 10 {
			spaces = "   "
		} else if i % 10 == 1 {
			spaces = " "
		} else {
			spaces = "  "
		}

		dateStr := strconv.Itoa(i)

		// this isn't working for some reason
		if i == today.Day() {
			dateStr = "\033[7m" + dateStr + "\033[27m"
		}

		output += spaces + strconv.Itoa(i)

		if i % 10 == 0 {
			output += "\n"
		}
	}

	io.WriteString(os.Stdout, output)
}

//func calbuilder(date fdate.Date) string {
//
//}
