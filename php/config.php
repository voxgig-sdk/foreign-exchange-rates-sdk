<?php
declare(strict_types=1);

// ForeignExchangeRates SDK configuration

class ForeignExchangeRatesConfig
{
    /** @var array<string,mixed>|null */
    private static ?array $shared_config = null;

    /**
     * Return the process-wide config, built once on first use. The SDK reads
     * the config on every request and never writes to it, so one instance is
     * shared by every client rather than rebuilt per client.
     *
     * PHP arrays are copy-on-write, so callers that do mutate the result get
     * their own copy and cannot disturb the shared one.
     */
    public static function shared_config(): array
    {
        if (self::$shared_config === null) {
            self::$shared_config = self::make_config();
        }
        return self::$shared_config;
    }

    /**
     * Build a fresh, fully materialised config array. Every call rebuilds the
     * whole structure, so prefer shared_config unless you need a private copy.
     */
    public static function make_config(): array
    {
        return [
            "main" => [
                "name" => "ForeignExchangeRates",
                "slug" => "foreign-exchange-rates",
                "version" => "0.0.1",
                "target" => "php",
            ],
            "feature" => [
                "test" => [
          'options' => [
            'active' => false,
          ],
        ],
            ],
            "options" => [
                "base" => "https://api.exchangerate.dev",
                "auth" => [
                    "prefix" => "Bearer",
                ],
                "headers" => [
          'content-type' => 'application/json',
        ],
                "entity" => [
                    "account" => [],
                    "convert" => [],
                    "currency" => [],
                    "range" => [],
                    "rate" => [],
                ],
            ],
            "entity" => [
        'account' => [
          'fields' => [
            [
              'name' => 'calls_this_month',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'limit',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'resets_on',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'account',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/v1/account',
                  'parts' => [
                    'v1',
                    'account',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.usage`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'convert' => [
          'fields' => [
            [
              'name' => 'conversions',
              'type' => '`$ARRAY`',
            ],
            [
              'name' => 'from',
              'op' => [
                'create' => [
                  'req' => true,
                  'type' => '`$STRING`',
                ],
              ],
              'type' => '`$STRING`',
            ],
            [
              'name' => 'pairs',
              'req' => true,
              'short' => 'Array of [targetCurrency, amount] tuples.',
              'type' => '`$ARRAY`',
            ],
          ],
          'name' => 'convert',
          'op' => [
            'create' => [
              'input' => 'data',
              'name' => 'create',
              'points' => [
                [
                  'args' => [],
                  'kind' => 'http',
                  'method' => 'POST',
                  'orig' => '/v1/convert',
                  'parts' => [
                    'v1',
                    'convert',
                  ],
                  'select' => [],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 100,
                        'kind' => 'param',
                        'name' => 'amount',
                        'orig' => 'amount',
                        'reqd' => true,
                        'type' => '`$NUMBER`',
                      ],
                      [
                        'example' => 'USD',
                        'kind' => 'param',
                        'name' => 'from',
                        'orig' => 'from',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 'EUR',
                        'kind' => 'param',
                        'name' => 'to',
                        'orig' => 'to',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/v1/convert/{from}/{to}/{amount}',
                  'parts' => [
                    'v1',
                    'convert',
                    '{from}',
                    '{to}',
                    '{amount}',
                  ],
                  'select' => [
                    'exist' => [
                      'amount',
                      'from',
                      'to',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'convert',
              ],
            ],
          ],
        ],
        'currency' => [
          'fields' => [
            [
              'name' => 'decimals',
              'type' => '`$INTEGER`',
            ],
            [
              'name' => 'derived',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'name',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'type',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'currency',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'fiat',
                        'kind' => 'query',
                        'name' => 'type',
                        'orig' => 'type',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/v1/currencies',
                  'parts' => [
                    'v1',
                    'currencies',
                  ],
                  'select' => [
                    'exist' => [
                      'type',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'range' => [
          'fields' => [],
          'name' => 'range',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'USD',
                        'kind' => 'query',
                        'name' => 'base',
                        'orig' => 'base',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => '2020-01-31',
                        'kind' => 'query',
                        'name' => 'end_date',
                        'orig' => 'end_date',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 'json',
                        'kind' => 'query',
                        'name' => 'format',
                        'orig' => 'format',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => '2020-01-01',
                        'kind' => 'query',
                        'name' => 'start_date',
                        'orig' => 'start_date',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 'EUR,GBP',
                        'kind' => 'query',
                        'name' => 'symbol',
                        'orig' => 'symbol',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/v1/range',
                  'parts' => [
                    'v1',
                    'range',
                  ],
                  'select' => [
                    'exist' => [
                      'base',
                      'end_date',
                      'format',
                      'start_date',
                      'symbol',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.rates`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [],
          ],
        ],
        'rate' => [
          'fields' => [
            [
              'name' => 'base',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'derivation_bps_max',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'derived',
              'type' => '`$BOOLEAN`',
            ],
            [
              'name' => 'pair',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'quote',
              'type' => '`$STRING`',
            ],
            [
              'name' => 'rate',
              'type' => '`$NUMBER`',
            ],
            [
              'name' => 'source',
              'type' => '`$STRING`',
            ],
          ],
          'name' => 'rate',
          'op' => [
            'load' => [
              'input' => 'data',
              'name' => 'load',
              'points' => [
                [
                  'args' => [
                    'query' => [
                      [
                        'example' => 'USD',
                        'kind' => 'query',
                        'name' => 'base',
                        'orig' => 'base',
                        'type' => '`$STRING`',
                      ],
                      [
                        'example' => 'EUR,GBP',
                        'kind' => 'query',
                        'name' => 'symbol',
                        'orig' => 'symbol',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/v1/latest',
                  'parts' => [
                    'v1',
                    'latest',
                  ],
                  'select' => [
                    'exist' => [
                      'base',
                      'symbol',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.rates`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => '2020-01-15',
                        'kind' => 'param',
                        'name' => 'date',
                        'orig' => 'date',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                    'query' => [
                      [
                        'example' => 'EUR,GBP',
                        'kind' => 'query',
                        'name' => 'symbol',
                        'orig' => 'symbol',
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/v1/{date}',
                  'parts' => [
                    'v1',
                    '{date}',
                  ],
                  'select' => [
                    'exist' => [
                      'date',
                      'symbol',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body.rates`',
                  ],
                ],
                [
                  'args' => [
                    'params' => [
                      [
                        'example' => 'eur-usd',
                        'kind' => 'param',
                        'name' => 'id',
                        'orig' => 'slug',
                        'reqd' => true,
                        'type' => '`$STRING`',
                      ],
                    ],
                  ],
                  'kind' => 'http',
                  'method' => 'GET',
                  'orig' => '/v1/rate/{slug}',
                  'parts' => [
                    'v1',
                    'rate',
                    '{id}',
                  ],
                  'rename' => [
                    'param' => [
                      'slug' => 'id',
                    ],
                  ],
                  'select' => [
                    'exist' => [
                      'id',
                    ],
                  ],
                  'transform' => [
                    'req' => '`reqdata`',
                    'res' => '`body`',
                  ],
                ],
              ],
            ],
          ],
          'relations' => [
            'ancestors' => [
              [
                'v1',
              ],
            ],
          ],
        ],
      ],
        ];
    }


    public static function make_feature(string $name)
    {
        require_once __DIR__ . '/features.php';
        return ForeignExchangeRatesFeatures::make_feature($name);
    }
}
