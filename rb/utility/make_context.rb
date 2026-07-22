# ForeignExchangeRates SDK utility: make_context
require_relative '../core/context'
module ForeignExchangeRatesUtilities
  MakeContext = ->(ctxmap, basectx) {
    ForeignExchangeRatesContext.new(ctxmap, basectx)
  }
end
