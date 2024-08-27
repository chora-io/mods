# Msg Service

The `governor` module provides a message service for interacting with the state of the module.

## Draft Definitions

`MsgUpdateParameters` - Only the admin can update module parameters.

...

`MsgCreateGovernor` - Any account can create (and therefore become) a governor. A governor can receive delegations from any account. If rewards are enabled, a governor can receive rewards that are either minted by the network or withdrawn from an account managed by the network. A governor can manage their own commission rates but only within the parameters defined by the network.

`MsgRetireGovernor` - A governor can retire from their position at any time, forfeiting their voting power and releasing their delegations. If there are any unclaimed rewards remaining in the governor account, the rewards will be automatically transferred from the governor account to the owner account.

`MsgRemoveGovernor` - A governor account can be removed by the network admin, enabling the network to remove a governor by force if necessary. Any unclaimed rewards can be sent to the owner account, sent to another account, or burned upon removal (to reduce the supply of the reward token). 

## Proto Definitions

~~The messages are defined in proto files available to view on [Buf Schema Registry](https://buf.build/chora/governor).~~

<!-- listed alphabetically -->
