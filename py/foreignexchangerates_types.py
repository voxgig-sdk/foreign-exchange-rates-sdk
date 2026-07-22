# Typed models for the ForeignExchangeRates SDK.
#
# GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
# params (op.<name>.points[].args.params[]). Field/param types come from the
# canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
# @voxgig/apidef VALID_CANON). Do not edit by hand.
#
# These are TypedDicts, not dataclasses: the SDK ops return/accept plain dicts
# at runtime, and a TypedDict IS a dict shape, so the types match the runtime.
# Optional (req:false) keys are modelled as TypedDict key-optionality
# (total=False), split into a required base + total=False subclass when a type
# has both required and optional keys.

from __future__ import annotations

from typing import TypedDict, Any


class Account(TypedDict, total=False):
    email: str
    key: str
    org: str
    usage: dict


class AccountLoadMatch(TypedDict, total=False):
    email: str
    key: str
    org: str
    usage: dict


class ConvertRequired(TypedDict):
    pair: list


class Convert(ConvertRequired, total=False):
    amount: float
    conversion: list
    converted: float
    to: str


class ConvertListMatch(TypedDict):
    amount: float
    to: str


class ConvertCreateDataRequired(TypedDict):
    pair: list


class ConvertCreateData(ConvertCreateDataRequired, total=False):
    amount: float
    conversion: list
    converted: float
    to: str


class Currency(TypedDict, total=False):
    decimal: int
    derived: bool
    name: str
    type: str


class CurrencyLoadMatch(TypedDict, total=False):
    decimal: int
    derived: bool
    name: str
    type: str


class Range(TypedDict, total=False):
    base: str
    end_date: str
    has_more: bool
    next_cursor: str
    rate: dict
    start_date: str


class RangeLoadMatch(TypedDict, total=False):
    base: str
    end_date: str
    has_more: bool
    next_cursor: str
    rate: dict
    start_date: str


class Rate(TypedDict, total=False):
    base: str
    data_updated_at: str
    derivation_bps_max: float
    derived: bool
    is_forward_filled: bool
    market_session: str
    notice: str
    pair: str
    quote: str
    rate: dict
    source: str
    timestamp: int


class RateLoadMatch(TypedDict, total=False):
    date: str
    id: str
