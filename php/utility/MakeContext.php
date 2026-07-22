<?php
declare(strict_types=1);

// ForeignExchangeRates SDK utility: make_context

require_once __DIR__ . '/../core/Context.php';

class ForeignExchangeRatesMakeContext
{
    public static function call(array $ctxmap, ?ForeignExchangeRatesContext $basectx): ForeignExchangeRatesContext
    {
        return new ForeignExchangeRatesContext($ctxmap, $basectx);
    }
}
