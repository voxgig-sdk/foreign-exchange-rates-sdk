<?php
declare(strict_types=1);

// ForeignExchangeRates SDK utility: prepare_body

class ForeignExchangeRatesPrepareBody
{
    public static function call(ForeignExchangeRatesContext $ctx): mixed
    {
        if ($ctx->op->input === 'data') {
            return ($ctx->utility->transform_request)($ctx);
        }
        return null;
    }
}
