# State

The `validator` module uses the `orm` module as an abstraction layer over the `KVStore` that enables the creation of database tables with primary and secondary keys.

For more information about the `orm` module, see [Cosmos SDK ADR 055: ORM](https://docs.cosmos.network/main/architecture/adr-055-orm).

## Proto Definitions

The state is defined in the proto files available to view on [Buf Schema Registry](https://buf.build/chora/validator).

<!-- listed alphabetically -->

- [SigningPolicy](https://buf.build/chora/validator/docs/main:chora.validator.v1#chora.validator.v1.SigningPolicy)
- [Validator](https://buf.build/chora/validator/docs/main:chora.validator.v1#chora.validator.v1.Validator)
- [ValidatorSigningInfo](https://buf.build/chora/validator/docs/main:chora.validator.v1#chora.validator.v1.ValidatorSigningInfo)
