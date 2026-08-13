# ForeignExchangeRates SDK exists test

import pytest
from foreignexchangerates_sdk import ForeignExchangeRatesSDK


class TestExists:

    def test_should_create_test_sdk(self):
        testsdk = ForeignExchangeRatesSDK.test(None, None)
        assert testsdk is not None
