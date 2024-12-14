package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	fmt.Println("This is simple implementation of Chyfer text")

	var plainText string
	var amount int
	decryptText := false
	encryptText := false

	fmt.Println("For encrypt text choose [1] and for plain text choose [2]: ")

	var choose int

	fmt.Scanln(&choose)

	if choose == 1 {
		encryptText = true
	} else if choose == 2 {
		decryptText = true
	} else {
		log.Fatal("Pls enter valid value")
	}

	scanner := bufio.NewScanner(os.Stdin)

	if encryptText {
		fmt.Println("Enter your plain text:")
	}

	if decryptText {
		fmt.Println("Enter your encrypt text:")
	}

	if scanner.Scan() {
		plainText = scanner.Text()

		println("Entered Text:", plainText)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal("Error while taking input")
	}
	fmt.Println("Enter your chyfer amount: ")

	fmt.Scanln(&amount)

	allChar := []rune(plainText)

	lowerCaseAlphabet := []rune("abcdefghijklmnopqrstuvwxyz")
	upperCaseAlphabet := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ")

	encryptedRune := make([]rune, len(allChar))
	againPlainText := make([]rune, len(allChar))

	if encryptText {

		for k, ch := range allChar {
			if ch == ' ' {
				encryptedRune[k] = '@'
			} else {
				for i, char := range lowerCaseAlphabet {
					if ch == char {
						encryptedRune[k] = lowerCaseAlphabet[(i+amount)%26]
						break
					}
				}
			}
		}
		for k, ch := range allChar {
			if ch == ' ' {
				encryptedRune[k] = '@'
			} else {
				for i, char := range upperCaseAlphabet {
					if ch == char {
						encryptedRune[k] = upperCaseAlphabet[(i+amount)%26]
						break
					}
				}
			}
		}

		fmt.Println("Encrypted text:", string(encryptedRune))
		return
	}

	if decryptText {

		for e, ch := range allChar {
			if ch == '@' {
				againPlainText[e] = ' '
			} else {
				if amount < 26 {
					for i, char := range lowerCaseAlphabet {

						if ch == char {
							charIndex := i - amount
							if charIndex < 0 {
								charIndex = len(lowerCaseAlphabet) + charIndex
							}
							againPlainText[e] = lowerCaseAlphabet[charIndex]

						}
					}
					for i, char := range upperCaseAlphabet {

						if ch == char {
							charIndex := i - amount
							if charIndex < 0 {
								charIndex = len(upperCaseAlphabet) + charIndex
							}
							againPlainText[e] = upperCaseAlphabet[charIndex]

						}
					}

				} else {
					for i, char := range lowerCaseAlphabet {

						if ch == char {

							charIndex := 26 + ((i - amount) % 26)
							fmt.Println(charIndex)
							if charIndex == 26 {
								charIndex = 0
							}

							againPlainText[e] = lowerCaseAlphabet[charIndex]
							break
						}
					}
					for i, char := range upperCaseAlphabet {

						if ch == char {

							charIndex := 26 + ((i - amount) % 26)

							if charIndex == 26 {
								charIndex = 0
							}

							againPlainText[e] = upperCaseAlphabet[charIndex]
							break
						}
					}
				}
			}
		}

		fmt.Println("plain text: ", string(againPlainText))
	}
}
