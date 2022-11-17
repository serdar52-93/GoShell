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

// Output of the current directory/Ausgabe des aktuellen Verzeichnisses (ls)
func ls() {

	// Directory we want to get all files from.
	directory := "/Users/serda/OneDrive/Bureau/Techstarter/aufgabe/Go projekt/projekt"

	// Open the directory.
	outputDirRead, _ := os.Open(directory)

	// Call ReadDir to get all files.
	outputDirFiles, _ := outputDirRead.ReadDir(0)

	// Loop over files.
	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]

		// Get name of file.
		outputNameHere := outputFileHere.Name()

		// Print name
		fmt.Println(outputNameHere)

	}
}

// Output of the current directory as a list/Ausgabe des aktuellen Verzeichnisses als Liste (ls -l)
func lsl() {

	// Directory we want to get all files from.
	directory := "/Users/serda/OneDrive/Bureau/Techstarter/aufgabe/Go projekt/projekt"

	// Open the directory.
	outputDirRead, _ := os.Open(directory)

	// Call ReadDir to get all files.
	outputDirFiles, _ := outputDirRead.ReadDir(0)

	// Loop over files.
	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]

		// Get name of file.
		outputNameHere := outputFileHere.Name()

		// Print name
		fmt.Println("Name:", outputNameHere)

		stats, err := os.Stat(directory)
		if err != nil {
			log.Fatal(err)

		}

		fmt.Printf("Permission: %s\n", stats.Mode())
		fmt.Printf("Size: %d\n", stats.Size())
		fmt.Printf("Modification Time: %s\n", stats.ModTime())

	}
}





