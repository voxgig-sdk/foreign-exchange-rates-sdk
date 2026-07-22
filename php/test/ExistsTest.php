<?php
declare(strict_types=1);

// ForeignExchangeRates SDK exists test

require_once __DIR__ . '/../foreignexchangerates_sdk.php';

use PHPUnit\Framework\TestCase;

class ExistsTest extends TestCase
{
    public function test_create_test_sdk(): void
    {
        $testsdk = ForeignExchangeRatesSDK::test(null, null);
        $this->assertNotNull($testsdk);
    }
}
