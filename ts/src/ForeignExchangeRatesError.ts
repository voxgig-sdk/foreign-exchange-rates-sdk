
import { Context } from './Context'


class ForeignExchangeRatesError extends Error {

  isForeignExchangeRatesError = true

  sdk = 'ForeignExchangeRates'

  code: string
  ctx: Context

  constructor(code: string, msg: string, ctx: Context) {
    super(msg)
    this.code = code
    this.ctx = ctx
  }

}

export {
  ForeignExchangeRatesError
}

