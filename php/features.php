<?php
declare(strict_types=1);

// ForeignExchangeRates SDK feature factory

require_once __DIR__ . '/feature/BaseFeature.php';
require_once __DIR__ . '/feature/TestFeature.php';


class ForeignExchangeRatesFeatures
{
    public static function make_feature(string $name)
    {
        switch ($name) {
            case "base":
                return new ForeignExchangeRatesBaseFeature();
            case "test":
                return new ForeignExchangeRatesTestFeature();
            default:
                return new ForeignExchangeRatesBaseFeature();
        }
    }
}
