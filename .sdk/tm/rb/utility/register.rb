# ForeignExchangeRates SDK utility registration
require_relative '../core/utility_type'
require_relative 'clean'
require_relative 'done'
require_relative 'make_error'
require_relative 'feature_add'
require_relative 'feature_hook'
require_relative 'feature_init'
require_relative 'fetcher'
require_relative 'make_fetch_def'
require_relative 'make_context'
require_relative 'make_options'
require_relative 'make_request'
require_relative 'make_response'
require_relative 'make_result'
require_relative 'make_point'
require_relative 'make_spec'
require_relative 'make_url'
require_relative 'param'
require_relative 'prepare_auth'
require_relative 'prepare_body'
require_relative 'prepare_headers'
require_relative 'prepare_method'
require_relative 'prepare_params'
require_relative 'prepare_path'
require_relative 'prepare_query'
require_relative 'result_basic'
require_relative 'result_body'
require_relative 'result_headers'
require_relative 'transform_request'
require_relative 'transform_response'

ForeignExchangeRatesUtility.registrar = ->(u) {
  u.clean = ForeignExchangeRatesUtilities::Clean
  u.done = ForeignExchangeRatesUtilities::Done
  u.make_error = ForeignExchangeRatesUtilities::MakeError
  u.feature_add = ForeignExchangeRatesUtilities::FeatureAdd
  u.feature_hook = ForeignExchangeRatesUtilities::FeatureHook
  u.feature_init = ForeignExchangeRatesUtilities::FeatureInit
  u.fetcher = ForeignExchangeRatesUtilities::Fetcher
  u.make_fetch_def = ForeignExchangeRatesUtilities::MakeFetchDef
  u.make_context = ForeignExchangeRatesUtilities::MakeContext
  u.make_options = ForeignExchangeRatesUtilities::MakeOptions
  u.make_request = ForeignExchangeRatesUtilities::MakeRequest
  u.make_response = ForeignExchangeRatesUtilities::MakeResponse
  u.make_result = ForeignExchangeRatesUtilities::MakeResult
  u.make_point = ForeignExchangeRatesUtilities::MakePoint
  u.make_spec = ForeignExchangeRatesUtilities::MakeSpec
  u.make_url = ForeignExchangeRatesUtilities::MakeUrl
  u.param = ForeignExchangeRatesUtilities::Param
  u.prepare_auth = ForeignExchangeRatesUtilities::PrepareAuth
  u.prepare_body = ForeignExchangeRatesUtilities::PrepareBody
  u.prepare_headers = ForeignExchangeRatesUtilities::PrepareHeaders
  u.prepare_method = ForeignExchangeRatesUtilities::PrepareMethod
  u.prepare_params = ForeignExchangeRatesUtilities::PrepareParams
  u.prepare_path = ForeignExchangeRatesUtilities::PreparePath
  u.prepare_query = ForeignExchangeRatesUtilities::PrepareQuery
  u.result_basic = ForeignExchangeRatesUtilities::ResultBasic
  u.result_body = ForeignExchangeRatesUtilities::ResultBody
  u.result_headers = ForeignExchangeRatesUtilities::ResultHeaders
  u.transform_request = ForeignExchangeRatesUtilities::TransformRequest
  u.transform_response = ForeignExchangeRatesUtilities::TransformResponse
}
