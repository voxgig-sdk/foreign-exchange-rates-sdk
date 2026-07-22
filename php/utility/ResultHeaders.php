<?php
declare(strict_types=1);

// ForeignExchangeRates SDK utility: result_headers

class ForeignExchangeRatesResultHeaders
{
    public static function call(ForeignExchangeRatesContext $ctx): ?ForeignExchangeRatesResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result) {
            if ($response && is_array($response->headers)) {
                $result->headers = $response->headers;
            } else {
                $result->headers = [];
            }
        }
        return $result;
    }
}
