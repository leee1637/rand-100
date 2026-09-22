package random

import (
	"github.com/fatih/color"
)

func CheckNum(randomNum, outNum int) bool {

	if randomNum == outNum {
		color.Green("Вы угадали число!!!")
		return true
	}

	return false
}

func DistanceNum(randomNum, outNum int) {
	diff := 0

	if randomNum > outNum {
		diff = randomNum - outNum
		switch {
		case diff <= 5:
			color.Yellow("Секретное число больше👆 - 🔥 Горячо")
		case diff <= 15:
			color.Yellow("Секретное число больше👆 - 🙂 Тепло")
		default:
			color.Yellow("Секретное число больше👆 - ❄️ Холодно")
		}
	}
	if outNum > randomNum {
		diff = outNum - randomNum
		switch {
		case diff <= 5:
			color.Yellow("Секретное число меньше👇 - 🔥 Горячо")
		case diff <= 15:
			color.Yellow("Секретное число меньше👇 - 🙂 Тепло")
		default:
			color.Yellow("Секретное число меньше👇 - ❄️ Холодно")
		}
	}
}
