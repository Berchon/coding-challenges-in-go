package main

import (
	"fmt"
	"strconv"
)

/*
Nome do problema:
Time Conversion

Link:
https://www.hackerrank.com/challenges/time-conversion/problem

Descrição do problema:

Dado um horário no formato de 12 horas, utilizando AM ou PM,
converta-o para o formato de 24 horas (military time).

O horário de entrada possui o formato:

hh:mm:ssAM
ou
hh:mm:ssPM

Regras importantes:

- 12:00:00AM corresponde a 00:00:00 no formato de 24 horas.
- 12:00:00PM corresponde a 12:00:00 no formato de 24 horas.
- Para horários PM diferentes de 12, deve-se adicionar 12 à hora.
- Para horários AM diferentes de 12, a hora permanece a mesma.

A entrada sempre representa um horário válido.

Exemplo:

Entrada:
07:05:45PM

Saída:
19:05:45
*/

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