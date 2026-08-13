-- Typed models for the ForeignExchangeRates SDK (LuaLS annotations).
--
-- GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
-- params (op.<name>.points[].args.params[]). Field/param types come from the
-- canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
-- @voxgig/apidef VALID_CANON). Annotations only — no runtime effect. Do not
-- edit by hand.

---@class Account
---@field calls_this_month? number
---@field limit? number
---@field resets_on? string

---@class AccountLoadMatch
---@field calls_this_month? number
---@field limit? number
---@field resets_on? string

---@class Convert
---@field amount? number
---@field conversions? table
---@field converted? number
---@field from? string
---@field pairs table
---@field to? string

---@class ConvertListMatch
---@field amount number
---@field from string
---@field to string

---@class ConvertCreateData
---@field amount? number
---@field conversions? table
---@field converted? number
---@field from? string
---@field pairs table
---@field to? string

---@class Currency
---@field decimals? number
---@field derived? boolean
---@field name? string
---@field type? string

---@class CurrencyLoadMatch
---@field decimals? number
---@field derived? boolean
---@field name? string
---@field type? string

---@class Range

---@class RangeLoadMatch

---@class Rate
---@field base? string
---@field derivation_bps_max? number
---@field derived? boolean
---@field pair? string
---@field quote? string
---@field rate? number
---@field source? string

---@class RateLoadMatch
---@field date? string
---@field id? string

local M = {}

return M
