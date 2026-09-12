# ForeignExchangeRates SDK configuration


# The sekreto plugin DEFINITIONS the model selected per feature, imported
# above by name from the modules the catalogue's active `plugin.def`
# entries declare. Handed to each feature (secrets builds its Sekreto
# with them): a provider kind not listed here is unknown to that SDK.
FEATURE_PLUGINS = {
}


_shared_config = None


def shared_config():
    """Return the process-wide config, built once on first use.

    The SDK reads the config on every request and never writes to it, so one
    instance is shared by every client rather than rebuilt per client.

    The returned dict is shared: treat it as read-only. Callers that need to
    mutate should use make_config, which always returns a fresh copy.
    """
    global _shared_config
    if _shared_config is None:
        _shared_config = make_config()
    return _shared_config


def make_config():
    """Build a fresh, fully materialised config dict.

    Every call rebuilds the whole structure, so prefer shared_config unless
    you need a private copy you intend to mutate.
    """
    return {
        "main": {
            "name": "ForeignExchangeRates",
            "slug": "foreign-exchange-rates",
            "version": "0.0.1",
            "target": "py",
        },
        "feature": {
            "test": {
        "options": {
          "active": False,
        },
        "transport": "base",
      },
        },
        "options": {
            "base": "https://api.exchangerate.dev",
            "auth": {
                "prefix": "Bearer",
            },
            "headers": {
        "content-type": "application/json",
      },
            "entity": {
                "account": {},
                "convert": {},
                "currency": {},
                "range": {},
                "rate": {},
            },
        },
        "entity": {
      "account": {
        "fields": [
          {
            "name": "calls_this_month",
            "type": "`$INTEGER`",
          },
          {
            "name": "limit",
            "type": "`$INTEGER`",
          },
          {
            "format": "date",
            "name": "resets_on",
            "type": "`$STRING`",
          },
        ],
        "name": "account",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {},
                "kind": "http",
                "method": "GET",
                "orig": "/v1/account",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "account",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.usage`",
                },
                "parts": [
                  "v1",
                  "account",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "convert": {
        "fields": [
          {
            "name": "conversions",
            "type": "`$ARRAY`",
          },
          {
            "name": "from",
            "op": {
              "create": {
                "req": True,
                "type": "`$STRING`",
              },
            },
            "type": "`$STRING`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "pairs",
            "req": True,
            "short": "Array of [targetCurrency, amount] tuples.",
            "type": "`$ARRAY`",
          },
        ],
        "id": {
          "field": "id",
          "from": {
            "from": "from",
          },
          "name": "id",
          "parts": [
            "from",
            "to",
            "amount",
          ],
          "sep": "/",
        },
        "name": "convert",
        "op": {
          "create": {
            "input": "data",
            "name": "create",
            "points": [
              {
                "args": {},
                "kind": "http",
                "method": "POST",
                "orig": "/v1/convert",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "convert",
                  },
                ],
                "select": {},
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "v1",
                  "convert",
                ],
              },
            ],
          },
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "params": [
                    {
                      "example": 100,
                      "kind": "param",
                      "name": "amount",
                      "orig": "amount",
                      "reqd": True,
                      "type": "`$NUMBER`",
                    },
                    {
                      "example": "USD",
                      "kind": "param",
                      "name": "from",
                      "orig": "from",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "EUR",
                      "kind": "param",
                      "name": "to",
                      "orig": "to",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/v1/convert/{from}/{to}/{amount}",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "convert",
                  },
                  {
                    "var": "from",
                  },
                  {
                    "var": "to",
                  },
                  {
                    "var": "amount",
                  },
                ],
                "select": {
                  "exist": [
                    "amount",
                    "from",
                    "to",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "v1",
                  "convert",
                  "{from}",
                  "{to}",
                  "{amount}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "convert",
            ],
          ],
        },
      },
      "currency": {
        "fields": [
          {
            "name": "decimals",
            "type": "`$INTEGER`",
          },
          {
            "name": "derived",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "name",
            "type": "`$STRING`",
          },
          {
            "name": "type",
            "type": "`$STRING`",
          },
        ],
        "name": "currency",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "query": [
                    {
                      "example": "fiat",
                      "kind": "query",
                      "name": "type",
                      "orig": "type",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/v1/currencies",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "currencies",
                  },
                ],
                "select": {
                  "exist": [
                    "type",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "v1",
                  "currencies",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "range": {
        "fields": [],
        "name": "range",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "query": [
                    {
                      "example": "USD",
                      "kind": "query",
                      "name": "base",
                      "orig": "base",
                      "type": "`$STRING`",
                    },
                    {
                      "example": "2020-01-31",
                      "kind": "query",
                      "name": "end_date",
                      "orig": "end_date",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "json",
                      "kind": "query",
                      "name": "format",
                      "orig": "format",
                      "type": "`$STRING`",
                    },
                    {
                      "example": "2020-01-01",
                      "kind": "query",
                      "name": "start_date",
                      "orig": "start_date",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                    {
                      "example": "EUR,GBP",
                      "kind": "query",
                      "name": "symbol",
                      "orig": "symbol",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/v1/range",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "range",
                  },
                ],
                "select": {
                  "exist": [
                    "base",
                    "end_date",
                    "format",
                    "start_date",
                    "symbol",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.rates`",
                },
                "parts": [
                  "v1",
                  "range",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [],
        },
      },
      "rate": {
        "fields": [
          {
            "name": "base",
            "type": "`$STRING`",
          },
          {
            "name": "derivation_bps_max",
            "type": "`$NUMBER`",
          },
          {
            "name": "derived",
            "type": "`$BOOLEAN`",
          },
          {
            "name": "id",
            "type": "`$STRING`",
          },
          {
            "name": "pair",
            "type": "`$STRING`",
          },
          {
            "name": "quote",
            "type": "`$STRING`",
          },
          {
            "name": "rate",
            "type": "`$NUMBER`",
          },
          {
            "name": "source",
            "type": "`$STRING`",
          },
        ],
        "id": {
          "field": "id",
          "name": "id",
        },
        "name": "rate",
        "op": {
          "load": {
            "input": "data",
            "name": "load",
            "points": [
              {
                "args": {
                  "query": [
                    {
                      "example": "USD",
                      "kind": "query",
                      "name": "base",
                      "orig": "base",
                      "type": "`$STRING`",
                    },
                    {
                      "example": "EUR,GBP",
                      "kind": "query",
                      "name": "symbol",
                      "orig": "symbol",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/v1/latest",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "latest",
                  },
                ],
                "select": {
                  "exist": [
                    "base",
                    "symbol",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.rates`",
                },
                "parts": [
                  "v1",
                  "latest",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "example": "2020-01-15",
                      "kind": "param",
                      "name": "date",
                      "orig": "date",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                  "query": [
                    {
                      "example": "EUR,GBP",
                      "kind": "query",
                      "name": "symbol",
                      "orig": "symbol",
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/v1/{date}",
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "var": "date",
                  },
                ],
                "select": {
                  "exist": [
                    "date",
                    "symbol",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body.rates`",
                },
                "parts": [
                  "v1",
                  "{date}",
                ],
              },
              {
                "args": {
                  "params": [
                    {
                      "example": "eur-usd",
                      "kind": "param",
                      "name": "id",
                      "orig": "slug",
                      "reqd": True,
                      "type": "`$STRING`",
                    },
                  ],
                },
                "kind": "http",
                "method": "GET",
                "orig": "/v1/rate/{slug}",
                "rename": {
                  "param": {
                    "slug": "id",
                  },
                },
                "segments": [
                  {
                    "lit": "v1",
                  },
                  {
                    "lit": "rate",
                  },
                  {
                    "var": "id",
                  },
                ],
                "select": {
                  "exist": [
                    "id",
                  ],
                },
                "transform": {
                  "req": "`reqdata`",
                  "res": "`body`",
                },
                "parts": [
                  "v1",
                  "rate",
                  "{id}",
                ],
              },
            ],
          },
        },
        "relations": {
          "ancestors": [
            [
              "v1",
            ],
          ],
        },
      },
    },
    }
