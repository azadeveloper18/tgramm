## API credentials

Kgramm does not use an application registered at Telegram.

The server function `CheckApiIdAndHash` ignores both arguments and returns nil, so the pair is a client build value, not a server secret.

Shipped pair:

- `api_id`: `10000`
- `api_hash`: `cdd52c2108d4da8b66dffc45aa689a61`

`configure.bat x64` uses that pair unless you override it. The published Telegram test pair `17349` / `344583e45741c457fe1862106095a5eb` is rejected.
