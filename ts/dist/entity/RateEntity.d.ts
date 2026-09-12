import { ForeignExchangeRatesEntityBase } from '../ForeignExchangeRatesEntityBase';
import type { ForeignExchangeRatesSDK } from '../ForeignExchangeRatesSDK';
import type { Control } from '../types';
import type { Rate, RateLoadMatch } from '../ForeignExchangeRatesTypes';
declare class RateEntity extends ForeignExchangeRatesEntityBase<Rate> {
    constructor(client: ForeignExchangeRatesSDK, entopts: any);
    make(this: RateEntity): RateEntity;
    load(this: any, reqmatch?: RateLoadMatch, ctrl?: Control): Promise<RateEntity>;
}
export { RateEntity };
