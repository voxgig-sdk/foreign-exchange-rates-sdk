package voxgigforeignexchangeratessdk

import (
	"github.com/voxgig-sdk/foreign-exchange-rates-sdk/go/core"
	"github.com/voxgig-sdk/foreign-exchange-rates-sdk/go/entity"
	"github.com/voxgig-sdk/foreign-exchange-rates-sdk/go/feature"
	_ "github.com/voxgig-sdk/foreign-exchange-rates-sdk/go/utility"
)

// Type aliases preserve external API.
type ForeignExchangeRatesSDK = core.ForeignExchangeRatesSDK
type Context = core.Context
type Utility = core.Utility
type Feature = core.Feature
type Entity = core.Entity
type ForeignExchangeRatesEntity = core.ForeignExchangeRatesEntity
type FetcherFunc = core.FetcherFunc
type Spec = core.Spec
type Result = core.Result
type Response = core.Response
type Operation = core.Operation
type Control = core.Control
type ForeignExchangeRatesError = core.ForeignExchangeRatesError

// BaseFeature from feature package.
type BaseFeature = feature.BaseFeature

func init() {
	core.NewBaseFeatureFunc = func() core.Feature {
		return feature.NewBaseFeature()
	}
	core.NewTestFeatureFunc = func() core.Feature {
		return feature.NewTestFeature()
	}
	core.NewAccountEntityFunc = func(client *core.ForeignExchangeRatesSDK, entopts map[string]any) core.ForeignExchangeRatesEntity {
		return entity.NewAccountEntity(client, entopts)
	}
	core.NewConvertEntityFunc = func(client *core.ForeignExchangeRatesSDK, entopts map[string]any) core.ForeignExchangeRatesEntity {
		return entity.NewConvertEntity(client, entopts)
	}
	core.NewCurrencyEntityFunc = func(client *core.ForeignExchangeRatesSDK, entopts map[string]any) core.ForeignExchangeRatesEntity {
		return entity.NewCurrencyEntity(client, entopts)
	}
	core.NewRangeEntityFunc = func(client *core.ForeignExchangeRatesSDK, entopts map[string]any) core.ForeignExchangeRatesEntity {
		return entity.NewRangeEntity(client, entopts)
	}
	core.NewRateEntityFunc = func(client *core.ForeignExchangeRatesSDK, entopts map[string]any) core.ForeignExchangeRatesEntity {
		return entity.NewRateEntity(client, entopts)
	}
}

// Constructor re-exports.
var NewForeignExchangeRatesSDK = core.NewForeignExchangeRatesSDK
var TestSDK = core.TestSDK
var NewContext = core.NewContext
var NewSpec = core.NewSpec
var NewResult = core.NewResult
var NewResponse = core.NewResponse
var NewOperation = core.NewOperation
var MakeConfig = core.MakeConfig

// No-arg convenience constructors. Go has no default-argument syntax,
// so these aliases let callers write `sdk.New()` / `sdk.Test()`
// instead of `sdk.NewForeignExchangeRatesSDK(nil)` / `sdk.TestSDK(nil, nil)`
// for the common no-options case.
func New() *ForeignExchangeRatesSDK  { return NewForeignExchangeRatesSDK(nil) }
func Test() *ForeignExchangeRatesSDK { return TestSDK(nil, nil) }
var NewBaseFeature = feature.NewBaseFeature
var NewTestFeature = feature.NewTestFeature
