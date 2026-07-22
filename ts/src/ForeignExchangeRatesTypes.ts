// Typed models for the ForeignExchangeRates SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.

export interface Account {
  email?: string
  key?: string
  org?: string
  usage?: Record<string, any>
}

export interface AccountLoadMatch {
  email?: string
  key?: string
  org?: string
  usage?: Record<string, any>
}

export interface Convert {
  amount?: number
  conversion?: any[]
  converted?: number
  from?: string
  pair: any[]
  to?: string
}

export interface ConvertListMatch {
  amount: number
  from: string
  to: string
}

export interface ConvertCreateData {
  amount?: number
  conversion?: any[]
  converted?: number
  from?: string
  pair: any[]
  to?: string
}

export interface Currency {
  decimal?: number
  derived?: boolean
  name?: string
  type?: string
}

export interface CurrencyLoadMatch {
  decimal?: number
  derived?: boolean
  name?: string
  type?: string
}

export interface Range {
  base?: string
  end_date?: string
  has_more?: boolean
  next_cursor?: string
  rate?: Record<string, any>
  start_date?: string
}

export interface RangeLoadMatch {
  base?: string
  end_date?: string
  has_more?: boolean
  next_cursor?: string
  rate?: Record<string, any>
  start_date?: string
}

export interface Rate {
  base?: string
  data_updated_at?: string
  derivation_bps_max?: number
  derived?: boolean
  is_forward_filled?: boolean
  market_session?: string
  notice?: string
  pair?: string
  quote?: string
  rate?: Record<string, any>
  source?: string
  timestamp?: number
}

export interface RateLoadMatch {
  date?: string
  id?: string
}

