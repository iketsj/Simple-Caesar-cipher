package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	characterTablePtr := flag.String("c", "", "This will be the character lookup table")
	textPtr := flag.String("t", "", "The text to be ciphered or deciphered")
	shiftValuePtr := flag.Int("s", 0, "The shift value for the cipher")
	modePtr := flag.String("m", "", "e or d for mode(encode or decode)")

	flag.Parse()

	var exitStatus int

	exitStatus = checkIfArgumentsAreCorrect(characterTablePtr, textPtr, shiftValuePtr, modePtr)
	if exitStatus != 0 {
		flag.PrintDefaults()
		os.Exit(exitStatus)
	}

	exitStatus = checkCipherMode(shiftValuePtr, modePtr)
	if exitStatus != 0 {
		flag.PrintDefaults()
		os.Exit(exitStatus)
	}

	putTheWordIntoTheCipher(characterTablePtr, textPtr, shiftValuePtr, modePtr)

}

func checkCipherMode(shiftValuePtr *int, modePtr *string) int {
	var exitStatus int = 0

	if *modePtr == "e" {
		fmt.Printf("Mode is to encode with a shift value of %d.\n", *shiftValuePtr)
	} else if *modePtr == "d" {
		fmt.Printf("Mode is to decode with a shift value of %d.\n", *shiftValuePtr)
	} else {
		fmt.Println("e or d (encode or decode")
		exitStatus = 6
	}
	
	return exitStatus
} 

func putTheWordIntoTheCipher(characterTablePtr *string, textPtr *string, shiftValuePtr *int, modePtr *string) {
	if *modePtr == "d" {
		*shiftValuePtr = -(*shiftValuePtr);
	}

	lengthOfCharacterTable := len(*characterTablePtr)
	lengthOfText := len(*textPtr)

	var notFound bool
	for wordIterator := 0; wordIterator < lengthOfText; wordIterator++ {
		for alphabetTableIterator := 0; alphabetTableIterator < lengthOfCharacterTable; alphabetTableIterator++ {
			notFound = true
			if (*textPtr)[wordIterator] == (*characterTablePtr)[alphabetTableIterator] {
				fmt.Printf("%c", (*characterTablePtr)[(alphabetTableIterator + *shiftValuePtr) % lengthOfCharacterTable])
				notFound = false 
				break
			}
		}
		if notFound {
			fmt.Printf("%c", (*textPtr)[wordIterator])
		}

	}
	fmt.Println()
}

func checkIfArgumentsAreCorrect(characterTablePtr *string, textPtr *string, shiftValuePtr *int, modePtr *string) int {
	var exitStatus int = 0

	if *characterTablePtr == "" {
		fmt.Println("Character table is empty")
		exitStatus = 1
	}

	if *textPtr == "" {
		fmt.Println("Text to be ciphered or deciphered is empty")
		exitStatus = 2
	}

	if *shiftValuePtr == 0 {
		fmt.Println("Change value to a non-zero value")
		exitStatus = 3
	}

	if *modePtr == "" {
		fmt.Println("No mode selected")
		exitStatus = 4
	}

	if flag.NArg() != 0 {
		fmt.Println("Excessive number of arguments")
		exitStatus = 5
	}

	return exitStatus
}