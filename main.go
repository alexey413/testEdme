package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	secretNumber := rand.Intn(10) + 1

	fmt.Println("Я загадал число от 1 до 10. Попробуй угадать!")

	for {
		var guess int
		fmt.Print("Твой вариант: ")
		_, err := fmt.Scan(&guess)

		if err != nil {
			fmt.Println("Ой! Нужно ввести число. Попробуй ещё раз.")
			// Очищаем буфер ввода
			var discard string
			fmt.Scanln(&discard)
			continue
		}

		if guess < 1 || guess > 10 {
			fmt.Println("Число должно быть от 1 до 10! Попробуй ещё раз.")
			continue
		}

		if guess == secretNumber {
			fmt.Printf("Ура! Ты угадал! Это было число %d.\n", secretNumber)
			break
		} else if guess < secretNumber {
			fmt.Println("Моё число больше!")
		} else {
			fmt.Println("Моё число меньше!")
		}
	}
}
