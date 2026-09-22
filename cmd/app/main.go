package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"random/internal/random"
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

	num, err := random.ScanLevelGameNum(scanner)
	if err != nil {
		fmt.Errorf("Ошибка сканирования: %w", err)
		return
	}

	countNum, randomNum, randomNumMax = random.GetLevelGame(num)

	fmt.Printf("Игра 'Угадай число' - от 1 до %v началась!\n", randomNumMax)
	fmt.Printf("Угадайте число за %v попыток!\n", countNum)

	for {
		sl := make([]int, 0, countNum)
		win := false

		for i := 0; i < countNum; i++ {

			color.Yellow("Попытка #%v - Введите число: ", i+1)

			num, err := random.ScanNum(scanner)

			if err != nil {
				fmt.Println("Ошибка парсинга! Введите число!")
				continue
			}

			if len(sl) != 0 {
				str := fmt.Sprint(sl)
				str = strings.Trim(str, "[]")
				resStr := strings.ReplaceAll(str, " ", ", ")
				color.Yellow("До этого вы вводили")
				fmt.Println(resStr)
			}

			err = random.ValidateNum(num, randomNumMax)
			if err != nil {
				color.Red("Я не засчитал попытку - вот причина:\n%v\n", err)
				i--
				color.Yellow("Напиши другое число")
				continue
			}

			ok := random.CheckNum(randomNum, num)

			if ok {
				win = true
				err = random.SaveGameResult("победа", i+1)
				if err != nil {
					fmt.Printf("Ошибка сохранения результата! %v", err)
				}

				break
			}

			random.DistanceNum(randomNum, num)

			sl = append(sl, num)

		}

		switch win {
		case true:
			fmt.Println("Поздравляю!!!")
		default:
			color.Red("Вы проиграли! Число было %v", randomNum)
			err := random.SaveGameResult("проигрыш", countNum)
			if err != nil {
				fmt.Printf("Ошибка сохранения результата! %v", err)
			}
			fmt.Println("Не расстрайивайтесь!")
		}

		fmt.Println("Хотите попробовать ещё раз?!")
		fmt.Println("Напишите 'да' или 'нет'")

		input, err := random.ScanYesOrNo(scanner)
		if err != nil {
			fmt.Errorf("Ошибка при сканировании: %w", err)
			return
		}

		switch input {
		case "да":
			randomNum = rand.IntN(randomNumMax) + 1
			color.Green("НАЧИНАЕМ ИГРУ ЗАНОВО!")
			continue
		case "нет":
			return
		}
	}
}
