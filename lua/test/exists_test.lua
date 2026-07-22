-- ForeignExchangeRates SDK exists test

local sdk = require("foreign-exchange-rates_sdk")

describe("ForeignExchangeRatesSDK", function()
  it("should create test SDK", function()
    local testsdk = sdk.test(nil, nil)
    assert.is_not_nil(testsdk)
  end)
end)
