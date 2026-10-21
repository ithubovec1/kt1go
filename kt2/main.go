package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Введите строку: ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	var cleaned []rune
	for _, r := range input {
		if unicode.IsLetter(r) || unicode.IsDigit(r) {
			cleaned = append(cleaned, unicode.ToLower(r))
		}
	}

	isPalindrome := true
	left, right := 0, len(cleaned)-1
	for left < right {
		if cleaned[left] != cleaned[right] {
			isPalindrome = false
			break
		}
		left++
		right--
	}

	if isPalindrome {
		fmt.Println("Строка является палиндромом.")
	} else {
		fmt.Println("Строка не является палиндромом.")
	}
}