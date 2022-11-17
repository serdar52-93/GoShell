package main

import(
	"fmt"
	"os"	
	"log"
)

// Changing directories/Wechseln von Verzeichnissen (cd)
func cd() {

	os.Chdir("/Users/serda")
	newDir, err := os.Getwd()
	if err != nil {
	}
	fmt.Printf("Current Working Directory: %s\n", newDir)

}

// Creating directories/Erstellen von Verzeichnissen (mkdir)
func Mkdir() {

	neuD := ""
	fmt.Println("Bitte geben Sie den Namen des neu zu erstellenden Datei ein: ")
	fmt.Scan(&neuD)
	err := os.Mkdir(neuD, 0750)
	fmt.Println("Die neue Datei von der Name:", neuD, "wird installiert!")

	if err != nil && !os.IsExist(err) {
		log.Fatal(err)
	}

}



