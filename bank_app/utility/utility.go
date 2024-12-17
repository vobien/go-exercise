package utility

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func WriteFloatToFile(value float64, filename string) {
	textValue := fmt.Sprintf("%f", value)
	os.WriteFile(filename, []byte(textValue), 0644)
}

func ReadFloatFromFile(filename string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return defaultValue, errors.New(fmt.Sprintf("Error reading value from file: %s", err))
	}

	val, err := strconv.ParseFloat(string(data), 64)
	if err != nil {
		return defaultValue, errors.New(fmt.Sprintf("Error parsing stored value from file: %s", err))
	}

	return val, nil
}
