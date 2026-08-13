# ForeignExchangeRates SDK utility: make_context

from projectname_sdk.core.context import ForeignExchangeRatesContext


def make_context_util(ctxmap, basectx):
    return ForeignExchangeRatesContext(ctxmap, basectx)
