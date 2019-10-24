# External Config Data Source

UCloud SDK provide external data source for loading client config and credential config.

The supported data source such as follow:

- Environment Variables
- Shared Config File
- STS Credentials on Metadata Server

There are some examples:

- [External Env & File](../../examples/external)
- [STS](../../examples/stscreds)

If STS token is expired, it will raise `client.CredentialExpiredError` .
