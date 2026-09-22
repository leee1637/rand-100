package random

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"random/internal/domain"
	"strings"
	"time"
)

func SaveGameResult(str string, countRand int) error {
	out, err := exec.Command("go", "env", "GOMOD").Output()
	if err != nil {
		return fmt.Errorf("не удалось определить корень проекта: %v", err)
	}

	modPath := strings.TrimSpace(string(out))

	var rootDir string

	if modPath == os.DevNull || modPath == "" {
		// Если go.mod не найден (например, проект не инициализирован через go mod init),
		// то в качестве запасного варианта берем текущую рабочую директорию
		rootDir, _ = os.Getwd()
	} else {
		rootDir = filepath.Dir(modPath)
	}

	filename := filepath.Join(rootDir, "results.json")

	var history []domain.GameResult

	fileData, err := os.ReadFile(filename)

	if err == nil {
		err = json.Unmarshal(fileData, &history)
		if err != nil {
			return fmt.Errorf("ошибка парсинга старого JSON: %v", err)
		}
	} else if !os.IsNotExist(err) {
		return fmt.Errorf("ошибка чтения диска: %v", err)
	}

	newResult := domain.GameResult{
		Date:     time.Now(),
		Result:   str,
		Attempts: countRand,
	}

	history = append(history, newResult)

	updatedData, err := json.MarshalIndent(history, "", "    ")
	if err != nil {
		return fmt.Errorf("ошибка кодирования в JSON: %v", err)
	}

	err = os.WriteFile(filename, updatedData, 0644)
	if err != nil {
		return fmt.Errorf("ошибка записи файла: %v", err)
	}

	return nil
}
