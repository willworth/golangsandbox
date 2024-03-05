package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644) //file permissions notation
}

func GetFloatFromFile(fileName string) (float64, error) {
	data, err := os.ReadFile(fileName)

	if err != nil {
		return 1000, errors.New("Failed to read file")
	}

	valueText := string(data)
	// balance, _ := strconv.ParseFloat(balanceText, 64)
	value, err := strconv.ParseFloat(valueText, 64)

	if err != nil {
		return 1000, errors.New("Failed to parse stored value")

	}

	// underscore indicates we know there's a var, but we don't
	// want to work with it, or have unused var name errors
	return value, nil
}
