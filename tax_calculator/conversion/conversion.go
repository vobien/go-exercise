package conversion

import (
	"strconv"
	"errors"
)

func StringToFloat(strings []string) ([]float64, error) {
	var prices []float64

	for _, str := range strings {
		p, err := strconv.ParseFloat(str, 64)

		if err != nil {
			return nil, errors.New("error parsing float")
		}

		prices = append(prices, p)
	}

	return prices, nil
}