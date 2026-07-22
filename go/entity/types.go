// Typed models for the ForeignExchangeRates SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
package entity

import "encoding/json"

// Account is the typed data model for the account entity.
type Account struct {
	Email *string `json:"email,omitempty"`
	Key *string `json:"key,omitempty"`
	Org *string `json:"org,omitempty"`
	Usage *map[string]any `json:"usage,omitempty"`
}

// AccountLoadMatch is the typed request payload for Account.LoadTyped.
type AccountLoadMatch struct {
	Email *string `json:"email,omitempty"`
	Key *string `json:"key,omitempty"`
	Org *string `json:"org,omitempty"`
	Usage *map[string]any `json:"usage,omitempty"`
}

// Convert is the typed data model for the convert entity.
type Convert struct {
	Amount *float64 `json:"amount,omitempty"`
	Conversion *[]any `json:"conversion,omitempty"`
	Converted *float64 `json:"converted,omitempty"`
	From *string `json:"from,omitempty"`
	Pair []any `json:"pair"`
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
	Conversion *[]any `json:"conversion,omitempty"`
	Converted *float64 `json:"converted,omitempty"`
	From *string `json:"from,omitempty"`
	Pair []any `json:"pair"`
	To *string `json:"to,omitempty"`
}

// Currency is the typed data model for the currency entity.
type Currency struct {
	Decimal *int `json:"decimal,omitempty"`
	Derived *bool `json:"derived,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// CurrencyLoadMatch is the typed request payload for Currency.LoadTyped.
type CurrencyLoadMatch struct {
	Decimal *int `json:"decimal,omitempty"`
	Derived *bool `json:"derived,omitempty"`
	Name *string `json:"name,omitempty"`
	Type *string `json:"type,omitempty"`
}

// Range is the typed data model for the range entity.
type Range struct {
	Base *string `json:"base,omitempty"`
	EndDate *string `json:"end_date,omitempty"`
	HasMore *bool `json:"has_more,omitempty"`
	NextCursor *string `json:"next_cursor,omitempty"`
	Rate *map[string]any `json:"rate,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
}

// RangeLoadMatch is the typed request payload for Range.LoadTyped.
type RangeLoadMatch struct {
	Base *string `json:"base,omitempty"`
	EndDate *string `json:"end_date,omitempty"`
	HasMore *bool `json:"has_more,omitempty"`
	NextCursor *string `json:"next_cursor,omitempty"`
	Rate *map[string]any `json:"rate,omitempty"`
	StartDate *string `json:"start_date,omitempty"`
}

// Rate is the typed data model for the rate entity.
type Rate struct {
	Base *string `json:"base,omitempty"`
	DataUpdatedAt *string `json:"data_updated_at,omitempty"`
	DerivationBpsMax *float64 `json:"derivation_bps_max,omitempty"`
	Derived *bool `json:"derived,omitempty"`
	IsForwardFilled *bool `json:"is_forward_filled,omitempty"`
	MarketSession *string `json:"market_session,omitempty"`
	Notice *string `json:"notice,omitempty"`
	Pair *string `json:"pair,omitempty"`
	Quote *string `json:"quote,omitempty"`
	Rate *map[string]any `json:"rate,omitempty"`
	Source *string `json:"source,omitempty"`
	Timestamp *int `json:"timestamp,omitempty"`
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

// typedFrom decodes a runtime value (a map[string]any produced by the op
// pipeline) into a typed model T via a JSON round-trip. On any error it
// returns the zero value of T; the op's own (value, error) tuple carries the
// real error.
func typedFrom[T any](v any) T {
	var out T
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

// typedSliceFrom decodes a runtime list value ([]any of maps) into a typed
// slice []T via a JSON round-trip, for list ops.
func typedSliceFrom[T any](v any) []T {
	var out []T
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
