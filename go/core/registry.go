package core

var UtilityRegistrar func(u *Utility)

var NewBaseFeatureFunc func() Feature

var NewTestFeatureFunc func() Feature

var NewAccountEntityFunc func(client *ForeignExchangeRatesSDK, entopts map[string]any) ForeignExchangeRatesEntity

var NewConvertEntityFunc func(client *ForeignExchangeRatesSDK, entopts map[string]any) ForeignExchangeRatesEntity

var NewCurrencyEntityFunc func(client *ForeignExchangeRatesSDK, entopts map[string]any) ForeignExchangeRatesEntity

var NewRangeEntityFunc func(client *ForeignExchangeRatesSDK, entopts map[string]any) ForeignExchangeRatesEntity

var NewRateEntityFunc func(client *ForeignExchangeRatesSDK, entopts map[string]any) ForeignExchangeRatesEntity

