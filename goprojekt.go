package main

import (
	"fmt"
)

func main() {

	i := " "

	fmt.Println("Willkommen im unsere neue GoShell ...")
	fmt.Println("Bitte geben Sie die gewünchte Aktion ein : ls, mkdir, cd, ls-l")
	fmt.Scan(&i)

	switch i {
	case "ls":
		ls()
	case "mkdir":
		Mkdir()
	case "cd":
		cd()

	case "ls-l":
		lsl()
	}

}
