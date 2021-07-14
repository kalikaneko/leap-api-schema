# leap-api-schema

Validate LEAP api json files.

## Build

go build

## Usage

```bash
❯ ./leap-api-schema
Usage: leap-api-schema <EIPURL>

❯ ./leap-api-schema https://float.bitmask.net/3/eip-service.json
/locations/Amsterdam: {"country_code":"NL"... "name" value is required

❯ ./leap-api-schema https://black.riseup.net/3/config/eip-service.json
OK: https://black.riseup.net/3/config/eip-service.json
```

