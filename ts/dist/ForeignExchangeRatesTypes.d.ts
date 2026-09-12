export interface Account {
    calls_this_month?: number;
    limit?: number;
    resets_on?: string;
}
export interface AccountLoadMatch {
    calls_this_month?: number;
    limit?: number;
    resets_on?: string;
}
export interface Convert {
    conversions?: any[];
    from?: string;
    id?: string;
    pairs: any[];
}
export interface ConvertLoadMatch {
    amount: number;
    from: string;
    to: string;
}
export interface ConvertCreateData {
    conversions?: any[];
    from?: string;
    id?: string;
    pairs: any[];
}
export interface Currency {
    decimals?: number;
    derived?: boolean;
    name?: string;
    type?: string;
}
export interface CurrencyLoadMatch {
    type?: string;
}
export interface Range {
}
export interface RangeLoadMatch {
    base?: string;
    end_date: string;
    format?: string;
    start_date: string;
    symbol?: string;
}
export interface Rate {
    base?: string;
    derivation_bps_max?: number;
    derived?: boolean;
    id?: string;
    pair?: string;
    quote?: string;
    rate?: number;
    source?: string;
}
export interface RateLoadMatch {
    date: string;
    symbol?: string;
}
