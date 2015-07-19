package raindrops

import (
	"strconv"
)

func Convert(number int) string {
	primeNumbers := [...]int{3, 5, 7}
	rainDrops := [...]string{"Pling", "Plang", "Plong"}
	answer := ""

	for i := range primeNumbers {
		if number%primeNumbers[i] == 0 {
			answer += rainDrops[i]
		}
	}

	if len(answer) == 0 {
		return strconv.Itoa(number)
	} else {
		return answer
	}
}
