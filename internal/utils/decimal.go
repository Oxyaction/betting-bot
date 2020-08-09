package utils

import (
	"regexp"

	"github.com/shopspring/decimal"
)

func DecimalFromText(text string) (decimal.Decimal, error) {
	reg, err := regexp.Compile("[^.0-9]+")
	if err != nil {
		panic(err)
	}
	replacedString := reg.ReplaceAllString(text, "")
	return decimal.NewFromString(replacedString)
}
