Feature: Msg/UpdateAgentMetadata

  UpdateAgentMetadata is successful when:
  - admin is the agent admin

  UpdateAgentMetadata has the following outcomes:
  - message response returned
  - Content is updated in state
  - EventUpdateAgentMetadata is emitted

  Rule: The admin must be the agent admin

    Background:
      Given agent
      """
      {
        "address": "address",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: admin is agent admin
      When msg update agent metadata
      """
      {
        "address": "address",
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect no error

    Scenario: admin is not agent admin
      When msg update agent metadata
      """
      {
        "address": "address",
        "admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "new_metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect the error
      """
      admin chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: agent admin chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: unauthorized
      """

  Rule: The message response is returned

    Background:
      Given agent
      """
      {
        "address": "address",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: message response returned
      When msg update agent metadata
      """
      {
        "address": "address",
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect response
      """
      {
        "address": "address"
      }
      """

    # No failing scenario - response is never returned when message fails

  Rule: Content is updated in state

    Background:
      Given agent
      """
      {
        "address": "address",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: state agent updated
      When msg update agent metadata
      """
      {
        "address": "address",
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect state agent
      """
      {
        "address": "address",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventUpdateAgentMetadata emitted

    Background:
      Given agent
      """
      {
        "address": "address",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: event update agent metadata emitted
      When msg update agent metadata
      """
      {
        "address": "address",
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect event update agent metadata
      """
      {
        "address": "address"
      }
      """

    # No failing scenario - event is never emitted when message fails
