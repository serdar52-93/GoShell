package main

import(
	"fmt"
	"os"	
)

// Changing directories/Wechseln von Verzeichnissen (cd)
func cd() {

	os.Chdir("/Users/serda")
	newDir, err := os.Getwd()
	if err != nil {
	}
	fmt.Printf("Current Working Directory: %s\n", newDir)

}


