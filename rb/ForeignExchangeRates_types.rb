# frozen_string_literal: true

# Typed models for the ForeignExchangeRates SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Member types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Ruby types are unenforced; these YARD
# annotations document the shapes. Do not edit by hand.

# Account entity data model.
#
# @!attribute [rw] calls_this_month
#   @return [Integer, nil]
#
# @!attribute [rw] limit
#   @return [Integer, nil]
#
# @!attribute [rw] resets_on
#   @return [String, nil]
Account = Struct.new(
  :calls_this_month,
  :limit,
  :resets_on,
  keyword_init: true
)

# Request payload for Account#load.
#
# @!attribute [rw] calls_this_month
#   @return [Integer, nil]
#
# @!attribute [rw] limit
#   @return [Integer, nil]
#
# @!attribute [rw] resets_on
#   @return [String, nil]
AccountLoadMatch = Struct.new(
  :calls_this_month,
  :limit,
  :resets_on,
  keyword_init: true
)

# Convert entity data model.
#
# @!attribute [rw] conversions
#   @return [Array, nil]
#
# @!attribute [rw] from
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] pairs
#   @return [Array]
Convert = Struct.new(
  :conversions,
  :from,
  :id,
  :pairs,
  keyword_init: true
)

# Request payload for Convert#load.
#
# @!attribute [rw] amount
#   @return [Float]
#
# @!attribute [rw] from
#   @return [String]
#
# @!attribute [rw] to
#   @return [String]
ConvertLoadMatch = Struct.new(
  :amount,
  :from,
  :to,
  keyword_init: true
)

# Request payload for Convert#create.
#
# @!attribute [rw] conversions
#   @return [Array, nil]
#
# @!attribute [rw] from
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] pairs
#   @return [Array]
ConvertCreateData = Struct.new(
  :conversions,
  :from,
  :id,
  :pairs,
  keyword_init: true
)

# Currency entity data model.
#
# @!attribute [rw] decimals
#   @return [Integer, nil]
#
# @!attribute [rw] derived
#   @return [Boolean, nil]
#
# @!attribute [rw] name
#   @return [String, nil]
#
# @!attribute [rw] type
#   @return [String, nil]
Currency = Struct.new(
  :decimals,
  :derived,
  :name,
  :type,
  keyword_init: true
)

# Request payload for Currency#load.
#
# @!attribute [rw] type
#   @return [String, nil]
CurrencyLoadMatch = Struct.new(
  :type,
  keyword_init: true
)

# Range entity data model.
class RangeType
end

# Request payload for Range#load.
#
# @!attribute [rw] base
#   @return [String, nil]
#
# @!attribute [rw] end_date
#   @return [String]
#
# @!attribute [rw] format
#   @return [String, nil]
#
# @!attribute [rw] start_date
#   @return [String]
#
# @!attribute [rw] symbol
#   @return [String, nil]
RangeLoadMatch = Struct.new(
  :base,
  :end_date,
  :format,
  :start_date,
  :symbol,
  keyword_init: true
)

# Rate entity data model.
#
# @!attribute [rw] base
#   @return [String, nil]
#
# @!attribute [rw] derivation_bps_max
#   @return [Float, nil]
#
# @!attribute [rw] derived
#   @return [Boolean, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
#
# @!attribute [rw] pair
#   @return [String, nil]
#
# @!attribute [rw] quote
#   @return [String, nil]
#
# @!attribute [rw] rate
#   @return [Float, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
Rate = Struct.new(
  :base,
  :derivation_bps_max,
  :derived,
  :id,
  :pair,
  :quote,
  :rate,
  :source,
  keyword_init: true
)

# Request payload for Rate#load.
#
# @!attribute [rw] date
#   @return [String]
#
# @!attribute [rw] symbol
#   @return [String, nil]
RateLoadMatch = Struct.new(
  :date,
  :symbol,
  keyword_init: true
)

