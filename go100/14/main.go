package main

import "fmt"

func main() {
	//i, j := 42, 2701

	i := 22
	j := 1987

	fmt.Println("i: ", i)
	fmt.Println()

	p := &i         // pointe vers i
	fmt.Println(*p) // lit i via le pointeur
	// what's the difference between *p and p?
	fmt.Println(p)
	fmt.Println("i: ", i)


	*p = 21         // modifie i via le pointeur
	// what's the difference here?
	fmt.Println(p)  // affiche la nouvelle valeur de i
	fmt.Println(*p)  // affiche la nouvelle valeur de i
	fmt.Println("i: ", i)


	// is p = xx same as *p = xx ?
	// p = 22    // this gives an error in the IDE - no it's not possible the compiler too gives error

	// does q too now points to p?
	fmt.Println()
	q := p
	fmt.Println("q: ", q)
	fmt.Println("*q: ", *q)
	fmt.Println("i: ", i)

	fmt.Println("p: ", p)
	fmt.Println("*p: ", *p)

	fmt.Println("i: ", i)


	// i change the content of q (*q)

	*q = 1001
	fmt.Println("*q ", *q)
	fmt.Println("i: ", i)

	fmt.Println()
	fmt.Println("before: j: ", j)
	p = &j         // pointe vers j
	*p = *p / 37   // divise j via le pointeur
	fmt.Println("*p: ", *p)
	fmt.Println("after j: ", j)

}
