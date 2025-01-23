package main

import (
	"fmt"
)

func main() {
	letters := []rune{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M',
		'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
	}
	input := "SOMEHOW"
	x := ceaser(input, 3, letters)
	fmt.Println(x)
	y := decrypt(x, 3, letters)
	fmt.Println(y)
}

func ceaser(arg string, k int, letters []rune) string {
	total := len(letters)

	letterNumMap := make(map[rune]int)
	numLetterMap := make(map[int]rune)

	for i, letter := range letters {
		letterNumMap[letter] = i
		numLetterMap[i] = letter
	}

	ciphered := []rune{}

	for _, c := range arg {
		letterNum := letterNumMap[c]
		cipheredNum := (letterNum + k) % total
		cipheredLetter := numLetterMap[cipheredNum]
		ciphered = append(ciphered, cipheredLetter)
	}
	return string(ciphered)
}

func decrypt(text string, k int, letters []rune) string {
	res := []rune{}
	total := len(letters)

	letterNumMap := make(map[rune]int)
	numLetterMap := make(map[int]rune)

	for i, letter := range letters {
		letterNumMap[letter] = i
		numLetterMap[i] = letter
	}

	for _, c := range text {
		letterNum := letterNumMap[c]
		decipheredNum := (letterNum - k)
		if decipheredNum < 0 {
			decipheredNum = (decipheredNum * -1) % total
		}
		cipheredLetter := numLetterMap[decipheredNum]
		res = append(res, cipheredLetter)

	}
	return string(res)
}
