package main

import (
	"fmt"
	"prototyp/goshell/befehle"
)

func main() {

	x := "... Willkommen im unsere neue GoShell ..."

	// Den String umkehren.
	str := x
	// Gibt den umgekehrten String "str" zurück.
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)

	// Variable i
	i := " "
	fmt.Println("Bitte geben Sie die gewünchte Aktion ein : cd, mkdir, ls, ls-l")
	// Scanner funktion, damit der Client die Aktion schreiben kann, die er ausführen möchte
	fmt.Scan(&i)

	// Switch case, um die gewünschte Funktion auswählen zu können
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

	// Autor
	fmt.Println("(c) Serdar Sadikoglu")
}

// Argument und gibt die Umkehrung von string zurück.
func reverse(s string) string {
	// Konvertieren zu rune
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {

		// die Buchstaben des Strings tauschen,
		// wie der erste mit dem letzten und so weiter.
		rns[i], rns[j] = rns[j], rns[i]
	}

	// Gibt den umgekehrten String zurück.
	return string(rns)
}
