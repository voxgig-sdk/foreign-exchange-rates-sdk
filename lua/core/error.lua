-- ForeignExchangeRates SDK error

local ForeignExchangeRatesError = {}
ForeignExchangeRatesError.__index = ForeignExchangeRatesError


function ForeignExchangeRatesError.new(code, msg, ctx)
  local self = setmetatable({}, ForeignExchangeRatesError)
  self.is_sdk_error = true
  self.sdk = "ForeignExchangeRates"
  self.code = code or ""
  self.msg = msg or ""
  self.ctx = ctx
  self.result = nil
  self.spec = nil
  return self
end


function ForeignExchangeRatesError:error()
  return self.msg
end


function ForeignExchangeRatesError:__tostring()
  return self.msg
end


return ForeignExchangeRatesError
