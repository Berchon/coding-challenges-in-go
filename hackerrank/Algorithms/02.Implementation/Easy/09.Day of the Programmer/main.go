package main

import "fmt"

/*
# Day of the Programmer

## Problem Description

Marie invented a Time Machine and wants to test it by time-traveling to visit Russia on the Day of the Programmer (the 256th day of the year) during a year in the inclusive range from 1700 to 2700.

From 1700 to 1917, Russia's official calendar was the Julian calendar; since 1919 they used the Gregorian calendar system. The transition from the Julian to Gregorian calendar system occurred in 1918, when the next day after January 31st was February 14th. This means that in 1918, February 14th was the 32nd day of the year in Russia.

In both calendar systems, February is the only month with a variable amount of days; it has 29 days during a leap year, and 28 days during all other years.

In the Julian calendar, leap years are divisible by 4; in the Gregorian calendar, leap years are either of the following:

- Divisible by 400.
- Divisible by 4 and not divisible by 100.

Given a year, `year`, find the date of the 256th day of that year according to the official Russian calendar during that year. Then print it in the format `dd.mm.yyyy`, where `dd` is the two-digit day, `mm` is the two-digit month, and `yyyy` is the year.

For example, the given `year = 1984`. 1984 is divisible by 4, so it is a leap year. The 256th day of a leap year after 1918 is September 12, so the answer is `12.09.1984`.

## Function Description

Complete the `dayOfProgrammer` function in the editor below. It should return a string representing the date of the 256th day of the given year.

### Parameters

- `year`: an integer representing the year.

### Returns

- `string`: the date of the 256th day in the format `dd.mm.yyyy`.

## Input Format

A single integer denoting the year, `year`.

## Constraints

- `1700 <= year <= 2700`

## Output Format

Print the full date of the Day of the Programmer during the given year in the format `dd.mm.yyyy`, where `dd` is the two-digit day, `mm` is the two-digit month, and `yyyy` is the year.

## Sample Input 0

2017

## Sample Output 0

13.09.2017

### Explanation 0

In the year `2017`, January has 31 days, February has 28 days, March has 31 days, April has 30 days, May has 31 days, June has 30 days, July has 31 days, and August has 31 days.

When we sum the total number of days in the first eight months, we get:

`31 + 28 + 31 + 30 + 31 + 30 + 31 + 31 = 243`

The Day of the Programmer is the 256th day, so we calculate:

`256 - 243 = 13`

Therefore, it falls on day 13 of the ninth month (September). The output is `13.09.2017`.

## Sample Input 1

2016

## Sample Output 1

12.09.2016

### Explanation 1

The year `2016` is a leap year, so February has 29 days, while all the other months have the same number of days as in 2017.

When we sum the total number of days in the first eight months, we get:

`31 + 29 + 31 + 30 + 31 + 30 + 31 + 31 = 244`

The Day of the Programmer is the 256th day, so we calculate:

`256 - 244 = 12`

Therefore, it falls on day 12 of the ninth month (September). The output is `12.09.2016`.

## Sample Input 2

1800

## Sample Output 2

12.09.1800

### Explanation 2

Since 1800 is a leap year according to the Julian calendar, February has 29 days. Therefore, the Day of the Programmer falls on September 12, and the output is `12.09.1800`.
*/

/*
 * Complete the 'dayOfProgrammer' function below.
 *
 * The function is expected to return a STRING.
 * The function accepts INTEGER year as parameter.
 */

func daysOf(month int, year int32) int {
	switch month {
	case 1, 3, 5, 7, 8, 10, 12:
		return 31
	case 4, 6, 9, 11:
		return 30
	case 2:
		if year == 1918 {
			return 15
		}
		if isLeapYear(year) {
			return 29
		}
		return 28
	default:
		return 0
		//[31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
	}
}

func isLeapYear(year int32) bool {
	if year < 1918 {
		return year % 4 == 0
	}

	return (year % 400 == 0) || (year % 4 == 0 && !(year % 100 == 0))
}

func dayOfProgrammer(year int32) string {
	targetDay := 256
	day, month := targetDay, 0
	
	for day > daysOf(month, year) {
		day -= daysOf(month, year)
		month++
	}

	return fmt.Sprintf("%02d.%02d.%04d", day, month, year)
}

func main() {
	// Anos bissextos e não bissextos antes de 1918 (calendário juliano)
	fmt.Println(dayOfProgrammer(1700), dayOfProgrammer(1700) == "12.09.1700")
	fmt.Println(dayOfProgrammer(1800), dayOfProgrammer(1800) == "12.09.1800")
	fmt.Println(dayOfProgrammer(1819), dayOfProgrammer(1819) == "13.09.1819")
	fmt.Println(dayOfProgrammer(1896), dayOfProgrammer(1896) == "12.09.1896")
	fmt.Println(dayOfProgrammer(1900), dayOfProgrammer(1900) == "12.09.1900")
	fmt.Println(dayOfProgrammer(1916), dayOfProgrammer(1916) == "12.09.1916")
	fmt.Println(dayOfProgrammer(1917), dayOfProgrammer(1917) == "13.09.1917")

	// Transição de calendário
	fmt.Println(dayOfProgrammer(1918), dayOfProgrammer(1918) == "26.09.1918")

	// Anos bissextos e não bissextos após 1918 (calendário gregoriano)
	fmt.Println(dayOfProgrammer(1919), dayOfProgrammer(1919) == "13.09.1919")
	fmt.Println(dayOfProgrammer(1920), dayOfProgrammer(1920) == "12.09.1920")
	fmt.Println(dayOfProgrammer(2000), dayOfProgrammer(2000) == "12.09.2000")
	fmt.Println(dayOfProgrammer(2001), dayOfProgrammer(2001) == "13.09.2001")
	fmt.Println(dayOfProgrammer(2016), dayOfProgrammer(2016) == "12.09.2016")
	fmt.Println(dayOfProgrammer(2017), dayOfProgrammer(2017) == "13.09.2017")
	fmt.Println(dayOfProgrammer(2020), dayOfProgrammer(2020) == "12.09.2020")
	fmt.Println(dayOfProgrammer(2024), dayOfProgrammer(2024) == "12.09.2024")
	fmt.Println(dayOfProgrammer(2025), dayOfProgrammer(2025) == "13.09.2025")

	// Casos de borda das regras gregorianas
	fmt.Println(dayOfProgrammer(1900), dayOfProgrammer(1900) == "12.09.1900") // Juliano: bissexto
	fmt.Println(dayOfProgrammer(2100), dayOfProgrammer(2100) == "13.09.2100") // Gregoriano: não bissexto
	fmt.Println(dayOfProgrammer(2400), dayOfProgrammer(2400) == "12.09.2400") // Gregoriano: bissexto por 400
	fmt.Println(dayOfProgrammer(2700), dayOfProgrammer(2700) == "13.09.2700") // Limite superior

	// Limites do desafio
	fmt.Println(dayOfProgrammer(1700), dayOfProgrammer(1700) == "12.09.1700") // Limite inferior
}
