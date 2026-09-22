package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"random/internal/random"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Выберете режим сложности!")
	fmt.Println("(Easy: 1-50, 15 попыток) - 1")
	fmt.Println("(Medium: 1-100, 10 попыток) - 2")
	fmt.Println("(Hard: 1-200, 5 попыток) - 3")

	countNum := 0
	randomNum := 0
	randomNumMax := 0

	if scanner.Scan() {
		input := scanner.Text()

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Ошибка парсинга! Введите число!")
			return
		}
		if num > 3 || num < 0 {
			fmt.Println("Число может быть только от 1 до 3")
			return
		}

		countNum, randomNum, randomNumMax = random.GetLevelGame(num)
	}

	fmt.Printf("Игра 'Угадай число' - от 1 до %v началась!\n", randomNumMax)
	fmt.Printf("Угадайте число за %v попыток!\n", countNum)

	sl := make([]int, 0, countNum)

	for t := 0; t < 2; t++ {
		for i := 0; i < countNum; i++ {

			if scanner.Scan() {
				if len(sl) != 0 {
					str := fmt.Sprint(sl)
					str = strings.Trim(str, "[]")
					resStr := strings.ReplaceAll(str, " ", ", ")
					color.Yellow("До этого вы вводили")
					fmt.Println(resStr)
				}
				input := scanner.Text()
				num, err := strconv.Atoi(input)

				if err != nil {
					fmt.Println("Ошибка парсинга! Введите число!")
					i--
					continue
				}

				err = random.ValidateNum(num, randomNumMax)
				if err != nil {
					color.Red("Я не засчитал попытку - вот причина:\n%v\n", err)
					i--
					color.Yellow("Напиши другое число")
					continue
				}

				ok, err := random.CheckNum(randomNum, num)
				if ok {
					err = random.SaveGameResult("победа", i+1)
					if err != nil {
						fmt.Printf("Ошибка сохранения результата! %v", err)
					}
					return
				}
				if err != nil {
					fmt.Println("Ошибка чека числа: %w", err)
					return
				}

				sl = append(sl, num)
			}

		}
		color.Black("Вы проиграли! Число было %v", randomNum)
		err := random.SaveGameResult("проигрыш", randomNumMax)
		if err != nil {
			fmt.Printf("Ошибка сохранения результата! %v", err)
		}
		fmt.Println("Не расстрайивайтесь! Хотите попробовать ещё раз?")
		fmt.Println("Напишите 'да' или 'нет'")

		if scanner.Scan() {
			input := scanner.Text()

			input = strings.TrimSpace(input)

			if input != "да" && input != "нет" {
				fmt.Println("Ответ только да или нет!")
				return
			}

			switch input {
			case "да":
				randomNum = rand.IntN(randomNumMax)
				color.Green("НАЧИНАЕМ ИГРУ ЗАНОВО!")
				continue
			case "нет":
				return
			}
		}

		color.Black("Вы проиграли! Число было %v", randomNum)
		err = random.SaveGameResult("проигрыш", randomNumMax)
		if err != nil {
			fmt.Printf("Ошибка сохранения результата! %v", err)
		}
		return
	}
}
