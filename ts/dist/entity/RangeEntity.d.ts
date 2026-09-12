import { ForeignExchangeRatesEntityBase } from '../ForeignExchangeRatesEntityBase';
import type { ForeignExchangeRatesSDK } from '../ForeignExchangeRatesSDK';
import type { Control } from '../types';
import type { Range, RangeLoadMatch } from '../ForeignExchangeRatesTypes';
declare class RangeEntity extends ForeignExchangeRatesEntityBase<Range> {
    constructor(client: ForeignExchangeRatesSDK, entopts: any);
    make(this: RangeEntity): RangeEntity;
    load(this: any, reqmatch?: RangeLoadMatch, ctrl?: Control): Promise<RangeEntity>;
}
export { RangeEntity };
