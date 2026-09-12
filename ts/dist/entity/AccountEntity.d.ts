import { ForeignExchangeRatesEntityBase } from '../ForeignExchangeRatesEntityBase';
import type { ForeignExchangeRatesSDK } from '../ForeignExchangeRatesSDK';
import type { Control } from '../types';
import type { Account, AccountLoadMatch } from '../ForeignExchangeRatesTypes';
declare class AccountEntity extends ForeignExchangeRatesEntityBase<Account> {
    constructor(client: ForeignExchangeRatesSDK, entopts: any);
    make(this: AccountEntity): AccountEntity;
    load(this: any, reqmatch?: AccountLoadMatch, ctrl?: Control): Promise<AccountEntity>;
}
export { AccountEntity };
