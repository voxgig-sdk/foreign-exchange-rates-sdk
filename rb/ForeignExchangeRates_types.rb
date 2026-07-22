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
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] key
#   @return [String, nil]
#
# @!attribute [rw] org
#   @return [String, nil]
#
# @!attribute [rw] usage
#   @return [Hash, nil]
Account = Struct.new(
  :email,
  :key,
  :org,
  :usage,
  keyword_init: true
)

# Request payload for Account#load.
#
# @!attribute [rw] email
#   @return [String, nil]
#
# @!attribute [rw] key
#   @return [String, nil]
#
# @!attribute [rw] org
#   @return [String, nil]
#
# @!attribute [rw] usage
#   @return [Hash, nil]
AccountLoadMatch = Struct.new(
  :email,
  :key,
  :org,
  :usage,
  keyword_init: true
)

# Convert entity data model.
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] conversion
#   @return [Array, nil]
#
# @!attribute [rw] converted
#   @return [Float, nil]
#
# @!attribute [rw] from
#   @return [String, nil]
#
# @!attribute [rw] pair
#   @return [Array]
#
# @!attribute [rw] to
#   @return [String, nil]
Convert = Struct.new(
  :amount,
  :conversion,
  :converted,
  :from,
  :pair,
  :to,
  keyword_init: true
)

# Request payload for Convert#list.
#
# @!attribute [rw] amount
#   @return [Float]
#
# @!attribute [rw] from
#   @return [String]
#
# @!attribute [rw] to
#   @return [String]
ConvertListMatch = Struct.new(
  :amount,
  :from,
  :to,
  keyword_init: true
)

# Request payload for Convert#create.
#
# @!attribute [rw] amount
#   @return [Float, nil]
#
# @!attribute [rw] conversion
#   @return [Array, nil]
#
# @!attribute [rw] converted
#   @return [Float, nil]
#
# @!attribute [rw] from
#   @return [String, nil]
#
# @!attribute [rw] pair
#   @return [Array]
#
# @!attribute [rw] to
#   @return [String, nil]
ConvertCreateData = Struct.new(
  :amount,
  :conversion,
  :converted,
  :from,
  :pair,
  :to,
  keyword_init: true
)

# Currency entity data model.
#
# @!attribute [rw] decimal
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
  :decimal,
  :derived,
  :name,
  :type,
  keyword_init: true
)

# Request payload for Currency#load.
#
# @!attribute [rw] decimal
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
CurrencyLoadMatch = Struct.new(
  :decimal,
  :derived,
  :name,
  :type,
  keyword_init: true
)

# Range entity data model.
#
# @!attribute [rw] base
#   @return [String, nil]
#
# @!attribute [rw] end_date
#   @return [String, nil]
#
# @!attribute [rw] has_more
#   @return [Boolean, nil]
#
# @!attribute [rw] next_cursor
#   @return [String, nil]
#
# @!attribute [rw] rate
#   @return [Hash, nil]
#
# @!attribute [rw] start_date
#   @return [String, nil]
Range = Struct.new(
  :base,
  :end_date,
  :has_more,
  :next_cursor,
  :rate,
  :start_date,
  keyword_init: true
)

# Request payload for Range#load.
#
# @!attribute [rw] base
#   @return [String, nil]
#
# @!attribute [rw] end_date
#   @return [String, nil]
#
# @!attribute [rw] has_more
#   @return [Boolean, nil]
#
# @!attribute [rw] next_cursor
#   @return [String, nil]
#
# @!attribute [rw] rate
#   @return [Hash, nil]
#
# @!attribute [rw] start_date
#   @return [String, nil]
RangeLoadMatch = Struct.new(
  :base,
  :end_date,
  :has_more,
  :next_cursor,
  :rate,
  :start_date,
  keyword_init: true
)

# Rate entity data model.
#
# @!attribute [rw] base
#   @return [String, nil]
#
# @!attribute [rw] data_updated_at
#   @return [String, nil]
#
# @!attribute [rw] derivation_bps_max
#   @return [Float, nil]
#
# @!attribute [rw] derived
#   @return [Boolean, nil]
#
# @!attribute [rw] is_forward_filled
#   @return [Boolean, nil]
#
# @!attribute [rw] market_session
#   @return [String, nil]
#
# @!attribute [rw] notice
#   @return [String, nil]
#
# @!attribute [rw] pair
#   @return [String, nil]
#
# @!attribute [rw] quote
#   @return [String, nil]
#
# @!attribute [rw] rate
#   @return [Hash, nil]
#
# @!attribute [rw] source
#   @return [String, nil]
#
# @!attribute [rw] timestamp
#   @return [Integer, nil]
Rate = Struct.new(
  :base,
  :data_updated_at,
  :derivation_bps_max,
  :derived,
  :is_forward_filled,
  :market_session,
  :notice,
  :pair,
  :quote,
  :rate,
  :source,
  :timestamp,
  keyword_init: true
)

# Request payload for Rate#load.
#
# @!attribute [rw] date
#   @return [String, nil]
#
# @!attribute [rw] id
#   @return [String, nil]
RateLoadMatch = Struct.new(
  :date,
  :id,
  keyword_init: true
)

