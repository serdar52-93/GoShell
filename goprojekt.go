package main

import (
	"fmt" // Importieren Package fmt für ausdrucken und im Terminal die Ergebnisse sehen
	"prototyp/goshell/befehle" // Importieren Package befehle , um die Funktionen Package main nutzen zu können
)

func main() {

				/* ********************** Intro von mein Shell ********************** */


	x := "... Willkommen im unsere neue GoShell ..."

	// Den String umkehren.
	str := x
	// Gibt den umgekehrten String "str" zurück.
	strRev := reverse(str)
	fmt.Println(str)
	fmt.Println(strRev)



				/* ********************** Mein Shell ********************** */

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


				/*************************************************************/

	// Autor
	fmt.Println("(c) Serdar Sadikoglu")
}


/* ***** Eine kleine Funktion, um das Projekt etwas lustiger zu machen ***** */

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
