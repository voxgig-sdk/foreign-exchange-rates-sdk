import { AccountEntity } from './entity/AccountEntity';
import { ConvertEntity } from './entity/ConvertEntity';
import { CurrencyEntity } from './entity/CurrencyEntity';
import { RangeEntity } from './entity/RangeEntity';
import { RateEntity } from './entity/RateEntity';
export type * from './ForeignExchangeRatesTypes';
import { inspect } from 'node:util';
import type { Context, Feature } from './types';
import { config } from './Config';
import { ForeignExchangeRatesEntityBase } from './ForeignExchangeRatesEntityBase';
import { Utility } from './utility/Utility';
import { BaseFeature } from './feature/base/BaseFeature';
declare const stdutil: Utility;
declare class ForeignExchangeRatesSDK {
    _mode: string;
    _options: any;
    _utility: Utility;
    _features: Feature[];
    _rootctx: Context;
    constructor(options?: any);
    options(): any;
    utility(): any;
    prepare(fetchargs?: any): Promise<any>;
    direct(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    _rawRequest(fetchargs?: any): Promise<Error | {
        ok: boolean;
        status: number;
        headers: any;
        data: any;
        err?: undefined;
    } | {
        ok: boolean;
        err: any;
        status?: undefined;
        headers?: undefined;
        data?: undefined;
    }>;
    graphql(query: string, variables?: any, ctrl?: any): Promise<any>;
    Account(entopts?: Record<string, any>): AccountEntity;
    Convert(entopts?: Record<string, any>): ConvertEntity;
    Currency(entopts?: Record<string, any>): CurrencyEntity;
    Range(entopts?: Record<string, any>): RangeEntity;
    Rate(entopts?: Record<string, any>): RateEntity;
    static test(testoptsarg?: any, sdkoptsarg?: any): ForeignExchangeRatesSDK;
    tester(testopts?: any, sdkopts?: any): ForeignExchangeRatesSDK;
    toJSON(): {
        name: string;
    };
    toString(): string;
    [inspect.custom](): string;
}
declare const SDK: typeof ForeignExchangeRatesSDK;
export { stdutil, config, BaseFeature, ForeignExchangeRatesEntityBase, ForeignExchangeRatesSDK, SDK, };
