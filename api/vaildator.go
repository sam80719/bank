package api

import (
	"github.com/go-playground/validator/v10"
	"github.com/sam80719/bank/util"
)

var supportedCurrencies = map[string]bool{
	"USD": true,
	"EUR": true,
	"CAD": true,
	"TWD": true,
}

var validCurrency validator.Func = func(fieldLevel validator.FieldLevel) bool {
	if currency, ok := fieldLevel.Field().Interface().(string); ok {
		return util.IsSupportedCurrency(currency)
	}
	return false
}
