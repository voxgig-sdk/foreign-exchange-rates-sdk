package core

type ForeignExchangeRatesError struct {
	IsForeignExchangeRatesError bool
	Sdk              string
	Code             string
	Msg              string
	Ctx              *Context
	Result           any
	Spec             any
}

func NewForeignExchangeRatesError(code string, msg string, ctx *Context) *ForeignExchangeRatesError {
	return &ForeignExchangeRatesError{
		IsForeignExchangeRatesError: true,
		Sdk:              "ForeignExchangeRates",
		Code:             code,
		Msg:              msg,
		Ctx:              ctx,
	}
}

func (e *ForeignExchangeRatesError) Error() string {
	return e.Msg
}
