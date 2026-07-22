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
    public ?string $email = null;
    public ?string $key = null;
    public ?string $org = null;
    public ?array $usage = null;
}

/** Request payload for Account#load. */
class AccountLoadMatch
{
    public ?string $email = null;
    public ?string $key = null;
    public ?string $org = null;
    public ?array $usage = null;
}

/** Convert entity data model. */
class Convert
{
    public ?float $amount = null;
    public ?array $conversion = null;
    public ?float $converted = null;
    public ?string $from = null;
    public array $pair;
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
    public ?array $conversion = null;
    public ?float $converted = null;
    public ?string $from = null;
    public array $pair;
    public ?string $to = null;
}

/** Currency entity data model. */
class Currency
{
    public ?int $decimal = null;
    public ?bool $derived = null;
    public ?string $name = null;
    public ?string $type = null;
}

/** Request payload for Currency#load. */
class CurrencyLoadMatch
{
    public ?int $decimal = null;
    public ?bool $derived = null;
    public ?string $name = null;
    public ?string $type = null;
}

/** Range entity data model. */
class Range
{
    public ?string $base = null;
    public ?string $end_date = null;
    public ?bool $has_more = null;
    public ?string $next_cursor = null;
    public ?array $rate = null;
    public ?string $start_date = null;
}

/** Request payload for Range#load. */
class RangeLoadMatch
{
    public ?string $base = null;
    public ?string $end_date = null;
    public ?bool $has_more = null;
    public ?string $next_cursor = null;
    public ?array $rate = null;
    public ?string $start_date = null;
}

/** Rate entity data model. */
class Rate
{
    public ?string $base = null;
    public ?string $data_updated_at = null;
    public ?float $derivation_bps_max = null;
    public ?bool $derived = null;
    public ?bool $is_forward_filled = null;
    public ?string $market_session = null;
    public ?string $notice = null;
    public ?string $pair = null;
    public ?string $quote = null;
    public ?array $rate = null;
    public ?string $source = null;
    public ?int $timestamp = null;
}

/** Request payload for Rate#load. */
class RateLoadMatch
{
    public ?string $date = null;
    public ?string $id = null;
}

