# Overview

The `admin` module enables a network to have a dynamic network admin. A network admin is an account that has permission to update module parameters.

Depending on how each module is configured in a blockchain application, the admin account can have full control over module parameters, i.e. all modules are configured to use the admin module, or partial control over module parameters, i.e. only some modules are configured to use the admin module.

A network could also have multiple network admins depending on how the blockchain application is configured, i.e. the blockchain application could be configured to run multiple instances of the admin module and permissions to update module parameters are distributed among admins.

## Contents

1. [Concepts](01_concepts.md)
2. [Msg Service](02_msg.md)
3. [Query Service](03_query.md)
4. [State](04_state.md)
5. [Events](05_events.md)
