# Msg Service

The `proposal` module provides a message service for interacting with the state of the module.

## Draft Definitions

`MsgSubmitProposal` - Any account can submit a proposal. A proposal will execute a message on behalf the admin or a module account depending on how the module is configured.

`MsgSubmitProposalDeposit` - Any account can submit a proposal deposit. The deposit will be returned to the account after the voting period has ended (unless the result is "no with veto").

`MsgWithdrawProposal` - Only the account that submitted the proposal (the "proposer") or the admin can withdraw a proposal. Any deposits will be returned to the accounts that made the deposits.

...

`MsgSubmitVote` - Any account with voting power can vote on a proposal. Voting power is determined by the set module parameters (e.g. quantity of tokens locked in the network, accounts defined by another module - e.g. governor or validator, or an allowlist managed by the proposal module itself).

`MsgUpdateVote` - The voter account can update their vote before the voting period has ended.

## Proto Definitions

~~The messages are defined in proto files available to view on [Buf Schema Registry](https://buf.build/chora/proposal).~~

<!-- listed alphabetically -->
