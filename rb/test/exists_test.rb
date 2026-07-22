# ForeignExchangeRates SDK exists test

require "minitest/autorun"
require_relative "../ForeignExchangeRates_sdk"

class ExistsTest < Minitest::Test
  def test_create_test_sdk
    testsdk = ForeignExchangeRatesSDK.test(nil, nil)
    assert !testsdk.nil?
  end
end
