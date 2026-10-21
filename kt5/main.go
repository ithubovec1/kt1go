package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	secret := rand.Intn(100) + 1
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Я загадал число от 1 до 100. Попробуй угадать!")

	attempts := 0

	for {
		fmt.Print("Твой вариант: ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка: введите целое число")
			continue
		}

		attempts++

		switch {
		case guess < secret:
			fmt.Println("Больше")
		case guess > secret:
			fmt.Println("Меньше")
		default:
			fmt.Printf("Поздравляю! Ты угадал число %d за %d попыток.\n", secret, attempts)
			return
		}
	}
}