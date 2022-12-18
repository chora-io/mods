Feature: Msg/CreateContent

  Msg/CreateContent is successful when:
  - always (an error is never returned)

  Msg/CreateContent has the following outcomes:
  - message response returned
  - Content is added to state
  - EventCreateContent is emitted

  Rule: An error is never returned

    Scenario: never error
      When msg create content
      """
      {
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect no error

    # No failing scenario - only internal errors can occur

  Rule: The message response is returned

    Scenario: message response returned
      When msg create content
      """
      {
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect response
      """
      {
        "id": 1
      }
      """

    # No failing scenario - response is never returned when message fails

  Rule: Content is added to state

    Scenario: state content added
      When msg create content
      """
      {
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect state content
      """
      {
        "id": 1,
        "creator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventCreateContent is emitted

    Scenario: event create content emitted
      When msg create content
      """
      {
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect event create content
      """
      {
        "id": 1
      }
      """

    # No failing scenario - event is never emitted when message fails
