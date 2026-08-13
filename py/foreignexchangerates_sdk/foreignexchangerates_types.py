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
    calls_this_month: int
    limit: int
    resets_on: str


class AccountLoadMatch(TypedDict, total=False):
    calls_this_month: int
    limit: int
    resets_on: str


class ConvertRequired(TypedDict):
    pairs: list


class Convert(ConvertRequired, total=False):
    amount: float
    conversions: list
    converted: float
    to: str


class ConvertListMatch(TypedDict):
    amount: float
    to: str


class ConvertCreateDataRequired(TypedDict):
    pairs: list


class ConvertCreateData(ConvertCreateDataRequired, total=False):
    amount: float
    conversions: list
    converted: float
    to: str


class Currency(TypedDict, total=False):
    decimals: int
    derived: bool
    name: str
    type: str


class CurrencyLoadMatch(TypedDict, total=False):
    decimals: int
    derived: bool
    name: str
    type: str


class Range(TypedDict):
    pass


class RangeLoadMatch(TypedDict):
    pass


class Rate(TypedDict, total=False):
    base: str
    derivation_bps_max: float
    derived: bool
    pair: str
    quote: str
    rate: float
    source: str


class RateLoadMatch(TypedDict, total=False):
    date: str
    id: str
