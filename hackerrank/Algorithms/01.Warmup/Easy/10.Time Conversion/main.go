package main

import (
	"fmt"
	"strconv"
)

func timeConversion(s string) string {
	switch {
	case s[8:] == "AM" && s[:2] == "12":
		return "00" + s[2:8]

	case s[8:] == "PM" && s[:2] != "12":
		hour := (int(s[0]-'0')*10 + int(s[1]-'0')) + 12
		return strconv.Itoa(hour) + s[2:8]

	default:
		return s[:8]
	}
}

func main() {
	s := "07:05:45PM"

	result := timeConversion(s)

	fmt.Println(result)
}