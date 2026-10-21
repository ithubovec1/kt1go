package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	const size = 10
	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = rand.Intn(100) + 1
	}

	fmt.Println("Массив:", arr)

	sum := 0
	minVal := arr[0]
	maxVal := arr[0]
	for _, v := range arr {
		sum += v
		if v < minVal {
			minVal = v
		}
		if v > maxVal {
			maxVal = v
		}
	}

	avg := float64(sum) / float64(size)

	fmt.Printf("Сумма элементов: %d\n", sum)
	fmt.Printf("Минимальное значение: %d\n", minVal)
	fmt.Printf("Максимальное значение: %d\n", maxVal)
	fmt.Printf("Среднее значение: %.2f\n", avg)
}