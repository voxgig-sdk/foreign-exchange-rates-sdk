<?php
declare(strict_types=1);

// ForeignExchangeRates SDK utility: result_body

class ForeignExchangeRatesResultBody
{
    public static function call(ForeignExchangeRatesContext $ctx): ?ForeignExchangeRatesResult
    {
        $response = $ctx->response;
        $result = $ctx->result;
        if ($result && $response && $response->json_func && $response->body) {
            $result->body = ($response->json_func)();
        }
        return $result;
    }
}
