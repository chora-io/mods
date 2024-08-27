# Concepts

## Admin

The admin is the account that has permission to update module parameters. Module parameters define the settings of the module. The same account can be the admin for all modules within an application in which case there would be a single "network admin" (see the [admin](https://docs.chora.io/specs/mods/admin/) module for more information).

## Validator

A validator is responsible for producing and verifying blocks by running the blockchain application.

A validator can be "active" or "inactive" depending on their status and the set module parameters. An "active" validator is a validator that has an active node and the status of the validator meets the criteria defined by the module parameters whereas an "inactive" validators is a validator that does not have an active node and/or does not meet the criteria defined by the module parameters.

## Validator Operator

An account that registers a validator becomes the operator of the validator. 
