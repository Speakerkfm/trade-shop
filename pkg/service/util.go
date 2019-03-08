package service

import (
	"fmt"
	"strconv"
)

func formatFloat(f float64) (float64, error) {
	return strconv.ParseFloat(fmt.Sprintf("%.2f", f), 64)
}
