-- Typed models for the ForeignExchangeRates SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Account
---@field email? string
---@field key? string
---@field org? string
---@field usage? table

---@class AccountLoadMatch
---@field email? string
---@field key? string
---@field org? string
---@field usage? table

---@class Convert
---@field amount? number
---@field conversion? table
---@field converted? number
---@field from? string
---@field pair table
---@field to? string

---@class ConvertListMatch
---@field amount number
---@field from string
---@field to string

---@class ConvertCreateData
---@field amount? number
---@field conversion? table
---@field converted? number
---@field from? string
---@field pair table
---@field to? string

---@class Currency
---@field decimal? number
---@field derived? boolean
---@field name? string
---@field type? string

---@class CurrencyLoadMatch
---@field decimal? number
---@field derived? boolean
---@field name? string
---@field type? string

---@class Range
---@field base? string
---@field end_date? string
---@field has_more? boolean
---@field next_cursor? string
---@field rate? table
---@field start_date? string

---@class RangeLoadMatch
---@field base? string
---@field end_date? string
---@field has_more? boolean
---@field next_cursor? string
---@field rate? table
---@field start_date? string

---@class Rate
---@field base? string
---@field data_updated_at? string
---@field derivation_bps_max? number
---@field derived? boolean
---@field is_forward_filled? boolean
---@field market_session? string
---@field notice? string
---@field pair? string
---@field quote? string
---@field rate? table
---@field source? string
---@field timestamp? number

---@class RateLoadMatch
---@field date? string
---@field id? string

local M = {}

return M
