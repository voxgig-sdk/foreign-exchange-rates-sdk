<?php
declare(strict_types=1);

// ForeignExchangeRates SDK base feature

class ForeignExchangeRatesBaseFeature
{
    public string $version;
    public string $name;
    public bool $active;

    // Positions this feature when added via the client `extend` option:
    // "__before__" / "__after__" / "__replace__" name an already-added
    // feature (mirrors the ts feature `_options`). Declared so setting it
    // on an extension instance avoids the dynamic-property deprecation.
    public ?array $_options = null;

    public function __construct()
    {
        $this->version = '0.0.1';
        $this->name = 'base';
        $this->active = true;
    }

    public function get_version(): string { return $this->version; }
    public function get_name(): string { return $this->name; }
    public function get_active(): bool { return $this->active; }

    public function init(ForeignExchangeRatesContext $ctx, array $options): void {}
    public function PostConstruct(ForeignExchangeRatesContext $ctx): void {}
    public function PostConstructEntity(ForeignExchangeRatesContext $ctx): void {}
    public function SetData(ForeignExchangeRatesContext $ctx): void {}
    public function GetData(ForeignExchangeRatesContext $ctx): void {}
    public function GetMatch(ForeignExchangeRatesContext $ctx): void {}
    public function SetMatch(ForeignExchangeRatesContext $ctx): void {}
    public function PrePoint(ForeignExchangeRatesContext $ctx): void {}
    public function PreSpec(ForeignExchangeRatesContext $ctx): void {}
    public function PreRequest(ForeignExchangeRatesContext $ctx): void {}
    public function PreResponse(ForeignExchangeRatesContext $ctx): void {}
    public function PreResult(ForeignExchangeRatesContext $ctx): void {}
    public function PreDone(ForeignExchangeRatesContext $ctx): void {}
    public function PreUnexpected(ForeignExchangeRatesContext $ctx): void {}
}
