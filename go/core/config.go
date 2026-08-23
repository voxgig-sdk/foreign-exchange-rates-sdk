package core

import (
	"sync"
)

// MakeConfig builds a fresh, fully materialised config map. Every call
// rebuilds the whole structure, so prefer SharedConfig unless you need a
// private copy you intend to mutate.
func MakeConfig() map[string]any {
	return map[string]any{
		"main": map[string]any{
			"name": "ForeignExchangeRates",
			"slug": "foreign-exchange-rates",
			"version": "0.0.1",
			"target": "go",
		},
		"feature": map[string]any{
			"test": map[string]any{
				"options": map[string]any{
					"active": false,
				},
			},
		},
		"options": map[string]any{
			"base": "https://api.exchangerate.dev",
			"auth": map[string]any{
				"prefix": "Bearer",
			},
			"headers": map[string]any{
				"content-type": "application/json",
			},
			"entity": map[string]any{
				"account": map[string]any{},
				"convert": map[string]any{},
				"currency": map[string]any{},
				"range": map[string]any{},
				"rate": map[string]any{},
			},
		},
		"entity": map[string]any{
			"account": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "calls_this_month",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "limit",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "resets_on",
						"type": "`$STRING`",
					},
				},
				"name": "account",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "GET",
								"orig": "/v1/account",
								"parts": []any{
									"v1",
									"account",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.usage`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"convert": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "conversions",
						"type": "`$ARRAY`",
					},
					map[string]any{
						"name": "from",
						"op": map[string]any{
							"create": map[string]any{
								"req": true,
								"type": "`$STRING`",
							},
						},
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "pairs",
						"req": true,
						"short": "Array of [targetCurrency, amount] tuples.",
						"type": "`$ARRAY`",
					},
				},
				"name": "convert",
				"op": map[string]any{
					"create": map[string]any{
						"input": "data",
						"name": "create",
						"points": []any{
							map[string]any{
								"args": map[string]any{},
								"kind": "http",
								"method": "POST",
								"orig": "/v1/convert",
								"parts": []any{
									"v1",
									"convert",
								},
								"select": map[string]any{},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": 100,
											"kind": "param",
											"name": "amount",
											"orig": "amount",
											"reqd": true,
											"type": "`$NUMBER`",
										},
										map[string]any{
											"example": "USD",
											"kind": "param",
											"name": "from",
											"orig": "from",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "EUR",
											"kind": "param",
											"name": "to",
											"orig": "to",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/v1/convert/{from}/{to}/{amount}",
								"parts": []any{
									"v1",
									"convert",
									"{from}",
									"{to}",
									"{amount}",
								},
								"select": map[string]any{
									"exist": []any{
										"amount",
										"from",
										"to",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"convert",
						},
					},
				},
			},
			"currency": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "decimals",
						"type": "`$INTEGER`",
					},
					map[string]any{
						"name": "derived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "name",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "type",
						"type": "`$STRING`",
					},
				},
				"name": "currency",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "fiat",
											"kind": "query",
											"name": "type",
											"orig": "type",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/v1/currencies",
								"parts": []any{
									"v1",
									"currencies",
								},
								"select": map[string]any{
									"exist": []any{
										"type",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"range": map[string]any{
				"fields": []any{},
				"name": "range",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "USD",
											"kind": "query",
											"name": "base",
											"orig": "base",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "2020-01-31",
											"kind": "query",
											"name": "end_date",
											"orig": "end_date",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "json",
											"kind": "query",
											"name": "format",
											"orig": "format",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "2020-01-01",
											"kind": "query",
											"name": "start_date",
											"orig": "start_date",
											"reqd": true,
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "EUR,GBP",
											"kind": "query",
											"name": "symbol",
											"orig": "symbol",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/v1/range",
								"parts": []any{
									"v1",
									"range",
								},
								"select": map[string]any{
									"exist": []any{
										"base",
										"end_date",
										"format",
										"start_date",
										"symbol",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.rates`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{},
				},
			},
			"rate": map[string]any{
				"fields": []any{
					map[string]any{
						"name": "base",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "derivation_bps_max",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "derived",
						"type": "`$BOOLEAN`",
					},
					map[string]any{
						"name": "pair",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "quote",
						"type": "`$STRING`",
					},
					map[string]any{
						"name": "rate",
						"type": "`$NUMBER`",
					},
					map[string]any{
						"name": "source",
						"type": "`$STRING`",
					},
				},
				"name": "rate",
				"op": map[string]any{
					"load": map[string]any{
						"input": "data",
						"name": "load",
						"points": []any{
							map[string]any{
								"args": map[string]any{
									"query": []any{
										map[string]any{
											"example": "USD",
											"kind": "query",
											"name": "base",
											"orig": "base",
											"type": "`$STRING`",
										},
										map[string]any{
											"example": "EUR,GBP",
											"kind": "query",
											"name": "symbol",
											"orig": "symbol",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/v1/latest",
								"parts": []any{
									"v1",
									"latest",
								},
								"select": map[string]any{
									"exist": []any{
										"base",
										"symbol",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.rates`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "2020-01-15",
											"kind": "param",
											"name": "date",
											"orig": "date",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
									"query": []any{
										map[string]any{
											"example": "EUR,GBP",
											"kind": "query",
											"name": "symbol",
											"orig": "symbol",
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/v1/{date}",
								"parts": []any{
									"v1",
									"{date}",
								},
								"select": map[string]any{
									"exist": []any{
										"date",
										"symbol",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body.rates`",
								},
							},
							map[string]any{
								"args": map[string]any{
									"params": []any{
										map[string]any{
											"example": "eur-usd",
											"kind": "param",
											"name": "id",
											"orig": "slug",
											"reqd": true,
											"type": "`$STRING`",
										},
									},
								},
								"kind": "http",
								"method": "GET",
								"orig": "/v1/rate/{slug}",
								"parts": []any{
									"v1",
									"rate",
									"{id}",
								},
								"rename": map[string]any{
									"param": map[string]any{
										"slug": "id",
									},
								},
								"select": map[string]any{
									"exist": []any{
										"id",
									},
								},
								"transform": map[string]any{
									"req": "`reqdata`",
									"res": "`body`",
								},
							},
						},
					},
				},
				"relations": map[string]any{
					"ancestors": []any{
						[]any{
							"v1",
						},
					},
				},
			},
		},
	}
}

var (
	sharedConfigOnce sync.Once
	sharedConfigVal  map[string]any
)

// SharedConfig returns the process-wide config, built once on first use.
// The SDK reads the config on every request and never writes to it, so one
// instance is shared by every client rather than rebuilt per client.
//
// The returned map is shared: treat it as read-only. Callers that need to
// mutate should use MakeConfig, which always returns a fresh copy.
func SharedConfig() map[string]any {
	sharedConfigOnce.Do(func() {
		sharedConfigVal = MakeConfig()
	})
	return sharedConfigVal
}

func makeFeature(name string) Feature {
	switch name {
	case "test":
		if NewTestFeatureFunc != nil {
			return NewTestFeatureFunc()
		}
	default:
		if NewBaseFeatureFunc != nil {
			return NewBaseFeatureFunc()
		}
	}
	return nil
}
