package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите первое число: ")
	input1, _ := reader.ReadString('\n')
	num1, err := strconv.ParseFloat(strings.TrimSpace(input1), 64)
	if err != nil {
		fmt.Println("Ошибка: введено не число")
		return
	}

	fmt.Print("Введите второе число: ")
	input2, _ := reader.ReadString('\n')
	num2, err := strconv.ParseFloat(strings.TrimSpace(input2), 64)
	if err != nil {
		fmt.Println("Ошибка: введено не число")
		return
	}

	fmt.Print("Выберите операцию (+, -, *, /): ")
	op, _ := reader.ReadString('\n')
	op = strings.TrimSpace(op)

	switch op {
	case "+":
		fmt.Printf("Результат: %.2f\n", num1+num2)
	case "-":
		fmt.Printf("Результат: %.2f\n", num1-num2)
	case "*":
		fmt.Printf("Результат: %.2f\n", num1*num2)
	case "/":
		if num2 == 0 {
			fmt.Println("Ошибка: деление на ноль невозможно!")
		} else {
			fmt.Printf("Результат: %.2f\n", num1/num2)
		}
	default:
		fmt.Println("Ошибка: неизвестная операция")
	}
}