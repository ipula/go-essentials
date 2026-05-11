package convertion

import (
	"errors"
	"strconv"
)

func StringToFloat64(strings []string) ([]float64, error) {

	var floats []float64

	for _, stringVal := range strings {
		floatPrice, err := strconv.ParseFloat(stringVal, 64)

		if err != nil {
			return nil, errors.New("converting string to float failed")
		}
		floats = append(floats, floatPrice)
	}
	return floats, nil
}
