package random

import (
	"fmt"
	"math/rand/v2"
)

func GetLevelGame(num int) (int, int, int) {
	countNum := 0
	randomNum := 0
	randomNumMax := 0

	switch num {
	case 1:
		fmt.Println("Вы выбрали Easy режим")
		countNum = 15
		randomNum = rand.IntN(51) + 1
		randomNumMax = 50
	case 2:
		fmt.Println("Вы выбрали Medium режим")
		countNum = 10
		randomNum = rand.IntN(101) + 1
		randomNumMax = 100
	case 3:
		fmt.Println("Вы выбрали Hard режим")
		countNum = 5
		randomNum = rand.IntN(201) + 1
		randomNumMax = 200
	}

	return countNum, randomNum, randomNumMax
}
