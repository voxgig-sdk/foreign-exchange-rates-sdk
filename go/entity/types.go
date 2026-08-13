// Typed models for the ForeignExchangeRates SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import (
	"encoding/json"

	"github.com/voxgig-sdk/foreign-exchange-rates-sdk/go/core"
)

// Account is the typed data model for the account entity.
type Account struct {
	CallsThisMonth *int `json:"calls_this_month,omitempty"`
	Limit *int `json:"limit,omitempty"`
	ResetsOn *string `json:"resets_on,omitempty"`
}

// AccountLoadMatch is the typed request payload for Account.LoadTyped.
type AccountLoadMatch struct {
	CallsThisMonth *int `json:"calls_this_month,omitempty"`
	Limit *int `json:"limit,omitempty"`
	ResetsOn *string `json:"resets_on,omitempty"`
}

// Convert is the typed data model for the convert entity.
type Convert struct {
	Amount *float64 `json:"amount,omitempty"`
	Conversions *[]any `json:"conversions,omitempty"`
	Converted *float64 `json:"converted,omitempty"`
	From *string `json:"from,omitempty"`
	Pairs []any `json:"pairs"`
	To *string `json:"to,omitempty"`
}

// ConvertListMatch is the typed request payload for Convert.ListTyped.
type ConvertListMatch struct {
	Amount float64 `json:"amount"`
	From string `json:"from"`
	To string `json:"to"`
}

// ConvertCreateData is the typed request payload for Convert.CreateTyped.
type ConvertCreateData struct {
	Amount *float64 `json:"amount,omitempty"`
	Conversions *[]any `json:"conversions,omitempty"`
	Converted *float64 `json:"converted,omitempty"`
	From *string `json:"from,omitempty"`
	Pairs []any `json:"pairs"`
	To *string `json:"to,omitempty"`
}

// Currency is the typed data model for the currency entity.
type Currency struct {
	Decimals *int `json:"decimals,omitempty"`
	Derived *bool `json:"derived,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// CurrencyLoadMatch is the typed request payload for Currency.LoadTyped.
type CurrencyLoadMatch struct {
	Decimals *int `json:"decimals,omitempty"`
	Derived *bool `json:"derived,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// Range is the typed data model for the range entity.
type Range struct {
}

// RangeLoadMatch is the typed request payload for Range.LoadTyped.
type RangeLoadMatch struct {
}

// Rate is the typed data model for the rate entity.
type Rate struct {
	Base *string `json:"base,omitempty"`
	DerivationBpsMax *float64 `json:"derivation_bps_max,omitempty"`
	Derived *bool `json:"derived,omitempty"`
	Pair *string `json:"pair,omitempty"`
	Quote *string `json:"quote,omitempty"`
	Rate *float64 `json:"rate,omitempty"`
	Source *string `json:"source,omitempty"`
}

// RateLoadMatch is the typed request payload for Rate.LoadTyped.
type RateLoadMatch struct {
	Date *string `json:"date,omitempty"`
	Id *string `json:"id,omitempty"`
}

// asMap turns a typed request/data struct into the map[string]any the
// runtime op pipeline consumes, honouring the json tags above.
func asMap(v any) map[string]any {
	out := map[string]any{}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// entityData unwraps an entity to its data map.
//
// Operations resolve to the ENTITY, not the raw data (see AGENTS.md), and an
// entity's fields are UNEXPORTED — marshalling one directly yields `{}`, so
// every typed accessor would silently hand back a zero-valued struct. The
// typed boundary therefore takes the data hop first.
func entityData(v any) any {
	if ent, ok := v.(core.Entity); ok {
		return ent.Data()
	}
	return v
}

// typedFrom decodes a runtime value (an entity, or the map[string]any the op
// pipeline produced) into a typed model T via a JSON round-trip. On any error
// it returns the zero value of T; the op's own (value, error) tuple carries
// the real error.
func typedFrom[T any](v any) T {
	var out T
	v = entityData(v)
	if v == nil {
		return out
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}

// typedSliceFrom decodes a runtime list value into a typed slice []T via a
// JSON round-trip, for list ops. `list` resolves to a slice of ENTITY
// instances, so each element takes the data hop.
func typedSliceFrom[T any](v any) []T {
	var out []T
	if v == nil {
		return out
	}
	if list, ok := v.([]any); ok {
		unwrapped := make([]any, 0, len(list))
		for _, item := range list {
			unwrapped = append(unwrapped, entityData(item))
		}
		v = unwrapped
	}
	b, err := json.Marshal(v)
	if err != nil {
		return out
	}
	_ = json.Unmarshal(b, &out)
	return out
}
