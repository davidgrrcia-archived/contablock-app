package utils

import (
	"fmt"
	"math"
	"math/big"
	"strconv"
	"time"
)

func TransformBigIntToFloat64(num *big.Int) (*float64, error) {
	if num == nil {
		return nil, nil
	}

	string := num.String()

	float, err := strconv.ParseFloat(string, 64)
	if err != nil {
		fmt.Println("Error:", err)
		return nil, err
	}

	return &float, err
}

func TransformFloat64ToBigInt(num *float64) *big.Int {
	if num == nil {
		zero := 0.00
		num = &zero
	}

	bigFloat := big.NewFloat(*num)
	bigInt, _ := bigFloat.Int(nil)

	return bigInt
}

func GetDecimalPart(number float64) float64 {
	wholePart := math.Floor(number)
	decimalPart := number - wholePart
	decimalInt := decimalPart * math.Pow(10, 5)

	return decimalInt
}

func TransformDateToBigInt(date time.Time) *big.Int {
	return big.NewInt(date.Unix())
}

func TransformBigIntToDate(num *big.Int) time.Time {
	return time.Unix(num.Int64(), 0)
}
