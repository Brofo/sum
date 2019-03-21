package main

import (
	"fmt"
	"os"
	"strconv"

	sum "github.com/Brofo/ica01/sum/sumpack"
)

func main() {
	s1 := os.Args[1]
	s2 := os.Args[2]

	//For at n1 og n2 skal få verdien fra i, må disse være samme type som i (uint64).
	//Vi konverterer disse til uint8 senere (Linje 43)
	var n1 uint64
	var n2 uint64
	//Programmet skal stoppe hvis input er utenfor kapasitet. (Linje 37)
	var inputError bool = false

	//Jeg har lagt inn en error-melding for input.
	if i, err := strconv.ParseUint(s1, 0, 8); err != nil {
		fmt.Println("Error: Input verdi 1 er utenfor kapasiteten til uint8")
		inputError = true
	} else {
		n1 = i
	}

	if i, err := strconv.ParseUint(s2, 0, 8); err != nil {
		fmt.Println("Error: Input verdi 2 er utenfor kapasiteten til uint8")
		inputError = true
	} else {
		n2 = i
	}

	if inputError == false {
		sum := sum.SumUint64(n1, n2)
		if sum >= 255 || sum <= 0 {
			fmt.Printf("Error: summen er utenfor kapasiteten til uint8 (%v).", sum)
		} else {
			//Hvis summen er innenfor grensa til uint8, konverterer vi til uint8:
			sumconv := uint8(sum)
			//Printer også ut type for å bevise at summen er av type uint8.
			fmt.Printf("Sum: %v, Type: %T", sumconv, sumconv)
		}
	}
}
