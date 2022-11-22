package befehle

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Changing directories/Wechseln von Verzeichnissen (cd)
func Cd() {
	
	// Aktuelle Verzeichniss
	currentDir, err := os.Getwd()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Aktuelle Verzeichnis: ", currentDir)
	}
	
	// Die gewünschte Richtung
	os.Chdir("/Users/serda/")

	currentDir, err = os.Getwd()
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Gewünschte Verzeichnis: ", currentDir)
	}
}


// Creating directories/Erstellen von Verzeichnissen (mkdir)
func Mkdir() {

	// Geben Sie den Ordnernamen ein
	neuV := ""
	fmt.Println("Bitte geben Sie den Namen des neu zu erstellenden Datei ein: ")
	fmt.Scan(&neuV)

	// Mkdir erstellt die neue Verzeichnis
	err := os.Mkdir(neuV, 0750)
	fmt.Println("Die neue Verzeichnis von der Name", neuV, "wird installiert!")

	// Es zeigt an, dass es installiert ist, aber es passiert nichts, da die Verzeichnis bereits vorhanden ist
	if err != nil && !os.IsExist(err) {		
		log.Fatal(err)		
	}

}

// Output of the current directory/Ausgabe des aktuellen Verzeichnisses (ls)
func Ls() {

	// Rufen Sie ReadDir auf, um alle Dateien abzurufen.
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Dieser Ordner enthält die folgenden Dateien: ")
	fmt.Println("--------------------------------------------")
	for _, f := range files {
		
		// Liste die Namen auf
		fmt.Println(f.Name())
	fmt.Println("--------------------------------------------")
	
	}
}

// Output of the current directory as a list/Ausgabe des aktuellen Verzeichnisses als Liste (ls -l)
func Lsl() {

	// Verzeichnis, aus dem wir alle Dateien abrufen möchten.
	directory := "/Users/serda/OneDrive/Bureau/Techstarter/aufgabe/Go projekt/projekt"
	    
	// Öffnen Sie das Verzeichnis "directory"
	outputDirRead, _ := os.Open(directory)

	// Rufen Sie ReadDir auf, um alle Dateien abzurufen.
	outputDirFiles, _ := outputDirRead.ReadDir(0)

	// For Schleife über Dateien.
	for outputIndex := range outputDirFiles {
		outputFileHere := outputDirFiles[outputIndex]

		// Dateinamen abrufen.
		outputNameHere := outputFileHere.Name()

		// Liste die Namen auf
		fmt.Println("--------------------------------------------")
		fmt.Println("Name:", outputNameHere)
		

		// Statistiken von Verzeichnis "directory"
		stats, err := os.Stat(directory)
		if err != nil {
			log.Fatal(err)

		}

		// Berechtigungen
		fmt.Printf("Permission: %s\n", stats.Mode())

		// Größe
		fmt.Printf("Size: %d\n", stats.Size())

		// Änderungszeit
		fmt.Printf("Modification Time: %s\n", stats.ModTime())

	}
}





