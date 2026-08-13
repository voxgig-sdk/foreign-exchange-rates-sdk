<?php
declare(strict_types=1);

// Typed models for the ForeignExchangeRates SDK.
//
// GENERATED from the API model: main.kit.entity.<e>.fields[] and per-op
// params (op.<name>.points[].args.params[]). Field/param types come from the
// canonical type sentinels via @voxgig/sdkgen canonToType (source of truth:
// @voxgig/apidef VALID_CANON). Do not edit by hand.
//
// These are documentation-grade value objects (PHP 8 typed properties),
// registered on the composer classmap autoload. The SDK boundary exchanges
// assoc-arrays; these classes name the shapes for tooling and typed callers.

/** Account entity data model. */
class Account
{
    public ?int $calls_this_month = null;
    public ?int $limit = null;
    public ?string $resets_on = null;
}

/** Request payload for Account#load. */
class AccountLoadMatch
{
    public ?int $calls_this_month = null;
    public ?int $limit = null;
    public ?string $resets_on = null;
}

/** Convert entity data model. */
class Convert
{
    public ?float $amount = null;
    public ?array $conversions = null;
    public ?float $converted = null;
    public ?string $from = null;
    public array $pairs;
    public ?string $to = null;
}

/** Request payload for Convert#list. */
class ConvertListMatch
{
    public float $amount;
    public string $from;
    public string $to;
}

/** Request payload for Convert#create. */
class ConvertCreateData
{
    public ?float $amount = null;
    public ?array $conversions = null;
    public ?float $converted = null;
    public ?string $from = null;
    public array $pairs;
    public ?string $to = null;
}

/** Currency entity data model. */
class Currency
{
    public ?int $decimals = null;
    public ?bool $derived = null;
    public ?string $name = null;
    public ?string $type = null;
}

/** Request payload for Currency#load. */
class CurrencyLoadMatch
{
    public ?int $decimals = null;
    public ?bool $derived = null;
    public ?string $name = null;
    public ?string $type = null;
}

/** Range entity data model. */
class Range
{
}

/** Request payload for Range#load. */
class RangeLoadMatch
{
}

/** Rate entity data model. */
class Rate
{
    public ?string $base = null;
    public ?float $derivation_bps_max = null;
    public ?bool $derived = null;
    public ?string $pair = null;
    public ?string $quote = null;
    public ?float $rate = null;
    public ?string $source = null;
}

/** Request payload for Rate#load. */
class RateLoadMatch
{
    public ?string $date = null;
    public ?string $id = null;
}

