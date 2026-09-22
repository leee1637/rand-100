package random

import (
	"fmt"

	"github.com/fatih/color"
)

func CheckNum(randomNum, outNum int) (bool, error) {

	if randomNum == outNum {
		color.Green("Вы угадали число!!!")
		return true, nil
	}

	diff := 0

	if randomNum > outNum {
		diff = randomNum - outNum
		switch {
		case diff <= 5:
			color.Yellow("Секретное число больше👆 - 🔥 Горячо")
			return false, nil
		case diff <= 15:
			color.Yellow("Секретное число больше👆 - 🙂 Тепло")
			return false, nil
		default:
			color.Yellow("Секретное число больше👆 - ❄️ Холодно")
			return false, nil
		}
	}
	if outNum > randomNum {
		diff = outNum - randomNum
		switch {
		case diff <= 5:
			color.Yellow("Секретное число меньше👇 - 🔥 Горячо")
			return false, nil
		case diff <= 15:
			color.Yellow("Секретное число меньше👇 - 🙂 Тепло")
			return false, nil
		default:
			color.Yellow("Секретное число меньше👇 - ❄️ Холодно")
			return false, nil
		}
	}

	return false, fmt.Errorf("Не удалось проверить число")
}
