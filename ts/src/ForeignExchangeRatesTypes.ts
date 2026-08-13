// Typed models for the ForeignExchangeRates SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Account {
  calls_this_month?: number
  limit?: number
  resets_on?: string
}

export interface AccountLoadMatch {
  calls_this_month?: number
  limit?: number
  resets_on?: string
}

export interface Convert {
  amount?: number
  conversions?: any[]
  converted?: number
  from?: string
  pairs: any[]
  to?: string
}

export interface ConvertListMatch {
  amount: number
  from: string
  to: string
}

export interface ConvertCreateData {
  amount?: number
  conversions?: any[]
  converted?: number
  from?: string
  pairs: any[]
  to?: string
}

export interface Currency {
  decimals?: number
  derived?: boolean
  name?: string
  type?: string
}

export interface CurrencyLoadMatch {
  decimals?: number
  derived?: boolean
  name?: string
  type?: string
}

export interface Range {
}

export interface RangeLoadMatch {
}

export interface Rate {
  base?: string
  derivation_bps_max?: number
  derived?: boolean
  pair?: string
  quote?: string
  rate?: number
  source?: string
}

export interface RateLoadMatch {
  date?: string
  id?: string
}

