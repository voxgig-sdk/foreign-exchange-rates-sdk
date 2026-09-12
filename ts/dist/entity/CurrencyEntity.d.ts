import { ForeignExchangeRatesEntityBase } from '../ForeignExchangeRatesEntityBase';
import type { ForeignExchangeRatesSDK } from '../ForeignExchangeRatesSDK';
import type { Control } from '../types';
import type { Currency, CurrencyLoadMatch } from '../ForeignExchangeRatesTypes';
declare class CurrencyEntity extends ForeignExchangeRatesEntityBase<Currency> {
    constructor(client: ForeignExchangeRatesSDK, entopts: any);
    make(this: CurrencyEntity): CurrencyEntity;
    load(this: any, reqmatch?: CurrencyLoadMatch, ctrl?: Control): Promise<CurrencyEntity>;
}
export { CurrencyEntity };
