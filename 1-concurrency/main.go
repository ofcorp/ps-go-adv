package main

import (
	"fmt"
	"math/rand"
	"time"
)

type result struct {
	num    int
	square int
}

func main() {
	const n = 10

	// Создаём свой генератор случайных чисел
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	numsCh := make(chan int)
	resultsCh := make(chan result)

	// 1-я горутина: генерирует 10 случайных чисел и отправляет во 2-ю
	go func() {
		defer close(numsCh)

		for i := 0; i < n; i++ {
			num := r.Intn(101) // от 0 до 100 включительно
			numsCh <- num
		}
	}()

	// 2-я горутина: читает числа, возводит в квадрат и шлёт в main
	go func() {
		defer close(resultsCh)

		for num := range numsCh {
			resultsCh <- result{
				num:    num,
				square: num * num,
			}
		}
	}()

	// main ждёт 10 чисел и выводит их
	for i := 0; i < n; i++ {
		res := <-resultsCh
		fmt.Printf("%d^2 = %d\n", res.num, res.square)
	}
}
