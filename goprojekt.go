package main

import (
	"fmt"
	"prototyp/goshell/befehle"
)

func main() {

	x := "... Willkommen im unsere neue GoShell ..."
	// Reversing the string.
	str := x
	// returns the reversed string.
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)

	i := " "

	fmt.Println("Bitte geben Sie die gewünchte Aktion ein : cd, mkdir, ls, ls-l")
	fmt.Scan(&i)

	switch i {
	case "cd":
		befehle.Cd() // Aufruf func Cd in der package befehle
	case "mkdir":
		befehle.Mkdir() // Aufruf func Mkdir in der package befehle
	case "ls":
		befehle.Ls() // Aufruf func Ls in der package befehle
	case "ls-l":
		befehle.Lsl() // Aufruf func Lsl in der package befehle
	}

}

// argument and return the reverse of string.
func reverse(s string) string {
	rns := []rune(s) // convert to rune
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// swap the letters of the string,
		// like first with last and so on.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// return the reversed string.
	return string(rns)
}
