package random

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"strconv"
	"strings"
)

func GetLevelGame(num int) (int, int, int) {
	countNum := 0
	randomNum := 0
	randomNumMax := 0

	switch num {
	case 1:
		fmt.Println("Вы выбрали Easy режим")
		countNum = 15
		randomNum = rand.IntN(50) + 1
		randomNumMax = 50
	case 2:
		fmt.Println("Вы выбрали Medium режим")
		countNum = 10
		randomNum = rand.IntN(100) + 1
		randomNumMax = 100
	case 3:
		fmt.Println("Вы выбрали Hard режим")
		countNum = 5
		randomNum = rand.IntN(200) + 1
		randomNumMax = 200
	}

	return countNum, randomNum, randomNumMax
}

func ScanLevelGameNum(scanner *bufio.Scanner) (int, error) {
	if scanner.Scan() {
		input := scanner.Text()

		num, err := strconv.Atoi(input)
		if err != nil {
			return 0, fmt.Errorf("Ошибка парсинга! Введите число!")
		}
		if num > 3 || num <= 0 {
			fmt.Println("Число может быть только от 1 до 3")
			return 0, fmt.Errorf("Ошибка парсинга! Введите число!")
		}

		return num, nil
	}

	return 0, fmt.Errorf("Ошибка парсинга! неизвестная")
}

func ScanYesOrNo(scanner *bufio.Scanner) (string, error) {
	if scanner.Scan() {
		input := scanner.Text()

		input = strings.TrimSpace(input)

		if input != "да" && input != "нет" {
			return "", fmt.Errorf("Ответ только да или нет!")
		}
		return input, nil
	}

	return "", fmt.Errorf("неизвестная ошибка при сканировании")
}

func ScanNum(scanner *bufio.Scanner) (int, error) {
	if scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.Atoi(input)
		return num, err
	}

	return 0, fmt.Errorf("Ошибка сканирования!")
}
