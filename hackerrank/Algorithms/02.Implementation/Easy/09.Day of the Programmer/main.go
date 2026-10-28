package main

import "fmt"

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
