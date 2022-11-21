package main

import (
	
	"prototyp/goshell/befehle"
	"fmt"
)

func main() {

	i := " "

	fmt.Println("Willkommen im unsere neue GoShell ...")
	fmt.Println("Bitte geben Sie die gewünchte Aktion ein : cd, mkdir, ls, ls-l")
	fmt.Scan(&i)

	switch i {
	case "cd":
		befehle.Cd()
	case "mkdir":
		befehle.Mkdir()
	case "ls":
		befehle.Ls()
	case "ls-l":
		befehle.Lsl()
	}

}
