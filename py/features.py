# ForeignExchangeRates SDK feature factory

from feature.base_feature import ForeignExchangeRatesBaseFeature
from feature.test_feature import ForeignExchangeRatesTestFeature


def _make_feature(name):
    features = {
        "base": lambda: ForeignExchangeRatesBaseFeature(),
        "test": lambda: ForeignExchangeRatesTestFeature(),
    }
    factory = features.get(name)
    if factory is not None:
        return factory()
    return features["base"]()
