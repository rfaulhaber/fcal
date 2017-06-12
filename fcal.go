package main

import (
	"github.com/rfaulhaber/fdate"
	"io"
	"os"
	"strconv"
	"strings"
    "runtime"
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
		} else if i%10 == 1 {
			spaces = " "
		} else {
			spaces = "  "
		}

		dateStr := strconv.Itoa(i)

        // disabling highlighting on Windows because I don't know how to do it
        if runtime.GOOS != "windows" && i == today.Day() {
			dateStr = "\033[7m" + dateStr + "\033[27m"
        }

		output += spaces + dateStr

		if i%10 == 0 {
			output += "\n"
		}
	}

	io.WriteString(os.Stdout, output)
}
