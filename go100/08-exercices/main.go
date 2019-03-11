package main

import "os"
import "fmt"
import "strconv"

//TODO: écrire une fonction prennant un nombre décimal 'x' en paramètre et retournant son inverse 1/x.

// func inverse () {

// }


func main() {
	//TODO: appeler la fonction ici et affichez le résultat avec x = 3 et x = 0
	fmt.Printf("len os.Args: %d\n", len(os.Args))

	numArgs := len(os.Args)
	if numArgs < 2 {
		fmt.Printf("Provide a numeric value")
		os.Exit(1)
	}

	fmt.Printf("arg: %s\n", os.Args[1])

	if number, err := strconv.ParseFloat(os.Args[1], 32); err != nil {
	    fmt.Println("Pls pass a number") // 3.1415927410125732
	    os.Exit()
	}

	// number, err := strconv.ParseFloat(os.Args[1], 64)
	// if err != "" {
	// 	fmt.Printf("%s is not a valid number\n")
	// 	os.Exit(1)
	// }
	fmt.Printf("err: %s\n", err)
	fmt.Printf("number: %f\n" , number)

	fmt.Printf("number: %f 1/number: %f\n", number, 1 / number)

	// fmt.Printf("1 / %f ", number, " = %f\n", (1 / number))
}
