# ForeignExchangeRates SDK feature factory

require_relative 'feature/base_feature'
require_relative 'feature/test_feature'


module ForeignExchangeRatesFeatures
  def self.make_feature(name)
    case name
    when "base"
      ForeignExchangeRatesBaseFeature.new
    when "test"
      ForeignExchangeRatesTestFeature.new
    else
      ForeignExchangeRatesBaseFeature.new
    end
  end
end
