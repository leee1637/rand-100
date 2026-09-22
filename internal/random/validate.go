package random

import "fmt"

func ValidateNum(num, max int) error {
	if num > max || num < 0 {
		return fmt.Errorf("Число не может быть больше %v или меньше 0", max)
	}

	return nil
}
