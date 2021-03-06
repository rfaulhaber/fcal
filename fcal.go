package main

import (
	"flag"
	"fmt"
	"github.com/rfaulhaber/fdate"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	stderr = log.New(os.Stderr, "fcal: ", 0)
)

func main() {
	todayFlag := flag.Bool("t", false, "prints today's date instead of the month")
	ruralFlag := flag.Bool("r", false, "prints the name of today's day on the Rural Calendar")
	flag.Parse()

	if *todayFlag && *ruralFlag {
		stderr.Fatalln("cannot use -t and -r flags together")
		os.Exit(1)
	}

	today := fdate.Today()


	if *todayFlag {
		output := ""
		if today.Month() == 13 {
			output += today.String() + " - " + fdate.CompDay(today.Day()-1).String() + "\n"
		} else {
			output = today.String() + "\n"
		}

		fmt.Println(output)
	} else if *ruralFlag {
		fmt.Println(today.RuralDay())
	} else {
		fmt.Println(CreateCalendar(today))
	}
}

func CreateCalendar(date fdate.Date) string {
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

	if date.Month() < 13 {
		for i := 0; i < 10; i++ {
			weekday := fdate.Weekday(i).String()
			abbr := strings.ToUpper(string(weekday[0])) + string(weekday[1:3])

			// "é" is two bytes, so we need to grab the next one
			if i == 9 {
				abbr += string(weekday[3])
			}

			output += abbr + " "
		}

		output += "\n"
	}

	if date.Month() == 13 {
		var days int

		if date.IsLeapYear() {
			days = 6
		} else {
			days = 5
		}

		for i := 0; i < days; i++ {
			dayIndex := i + 1
			if dayIndex == date.Day() {
				output += highlightDate(strconv.Itoa(dayIndex))
			} else {
				output += " "
				output += strconv.Itoa(dayIndex)
			}

			output += " - " + fdate.CompDay(i).String() + "\n"
		}
	} else {
		monthDays := 31

		for i := 1; i < monthDays; i++ {
			spaces := ""

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
