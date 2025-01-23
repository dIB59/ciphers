package main

import (
	"fmt"
)

func main() {
	letters := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}
	ceaser("SOME HOW", 3, letters)
}

func ceaser(arg string, k int, letters []rune) {
	total := len(letters)

	letterNumMap := make(map[rune]int)
	numLetterMap := make(map[int]rune)

	for i, letter := range letters {
		letterNumMap[letter] = i
		numLetterMap[i] = letter
	}

	normal := []rune{}
	ciphered := []rune{}

	for _, c := range arg {
		letterNum := letterNumMap[c]
		cipheredNum := (letterNum + k) % total
		cipheredLetter := numLetterMap[cipheredNum]
		ciphered = append(ciphered, cipheredLetter)
		normal = append(normal, c)
	}

	fmt.Println(string(normal))
	fmt.Println(string(ciphered))

}
