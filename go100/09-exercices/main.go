package main

import "fmt"

//TODO: écrire une fonction qui prend en paramètre un entier et un nombre max d'itérations à effectuer, et qui effectue l'algorithme de
//la conjecture de Syracuse (voir README.md).
//Elle devra retourner un booléen indiquant si le nombre 1 a été atteint, suivi du nombre d'itérations qui ont été effectuées.

func syracuse(number int, maxIterationCount int) (bool, int) {
	fmt.Printf("inside syracuse\n")

	if number < 1 {
		return false, 0
	}

	i := 0
	for n := number; n > 1 && i < maxIterationCount; i++ {
		fmt.Printf("iteration %d, n = %d\n", i, n)
		// if the number is even then ....
		if n%2 == 0 {
			n = n / 2
		} else {
			// if it's odd then ...
			n = 3*n + 1
		}
	}
	// we return a bool and an int indicating the number of iterations that were necessary
	return i < maxIterationCount, i
}

func main() {
	var maxIterationCount = 1000
	var integer = 98375
	reached, iterationCount := syracuse(integer, maxIterationCount)
	if reached {
		fmt.Printf("L'algorithme a permis d'atteindre le nombre 1 à partir du nombre %d en %d itérations\n", integer, iterationCount)
	} else {
		fmt.Printf("L'algorithme n'a pas permis d'atteindre le nombre 1 à partir du nombre %d en moins de %d itérations\n", integer, maxIterationCount)
	}

	fmt.Printf("\n\n second round \n\n")

	maxIterationCount = 96
	integer = 98375
	reached, iterationCount = syracuse(integer, maxIterationCount)
	if reached {
		fmt.Printf("L'algorithme a permis d'atteindre le nombre 1 à partir du nombre %d en %d itérations\n", integer, iterationCount)
	} else {
		fmt.Printf("L'algorithme n'a pas permis d'atteindre le nombre 1 à partir du nombre %d en moins de %d itérations\n", integer, maxIterationCount)
	}
}
