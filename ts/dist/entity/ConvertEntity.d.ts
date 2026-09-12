import { ForeignExchangeRatesEntityBase } from '../ForeignExchangeRatesEntityBase';
import type { ForeignExchangeRatesSDK } from '../ForeignExchangeRatesSDK';
import type { Control } from '../types';
import type { Convert, ConvertLoadMatch, ConvertCreateData } from '../ForeignExchangeRatesTypes';
declare class ConvertEntity extends ForeignExchangeRatesEntityBase<Convert> {
    constructor(client: ForeignExchangeRatesSDK, entopts: any);
    make(this: ConvertEntity): ConvertEntity;
    load(this: any, reqmatch?: ConvertLoadMatch, ctrl?: Control): Promise<ConvertEntity>;
    create(this: any, reqdata?: ConvertCreateData, ctrl?: Control): Promise<ConvertEntity>;
}
export { ConvertEntity };
