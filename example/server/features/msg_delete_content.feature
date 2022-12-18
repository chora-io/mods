Feature: Msg/DeleteContent

  DeleteContent is successful when:
  - creator is the content creator

  DeleteContent has the following outcomes:
  - message response returned
  - Content is removed from state
  - EventDeleteContent is emitted

  Rule: The creator must be the content creator

    Background:
      Given content
      """
      {
        "id": 1,
        "creator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: creator is content creator
      When msg delete content
      """
      {
        "id": 1,
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect no error

    Scenario: creator is not content creator
      When msg delete content
      """
      {
        "id": 1,
        "creator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect the error
      """
      creator chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: content creator chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: unauthorized
      """

  Rule: The message response is returned

    Background:
      Given content
      """
      {
        "id": 1,
        "creator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: message response returned
      When msg delete content
      """
      {
        "id": 1,
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "id": 1
      }
      """

    # No failing scenario - response is never returned when message fails

  Rule: Content is removed from state

    Background:
      Given content
      """
      {
        "id": 1,
        "creator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: state content removed
      When msg delete content
      """
      {
        "id": 1,
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect no state content with id "1"

    # No failing scenario - state is never updated when message fails

  Rule: EventDeleteContent emitted

    Background:
      Given content
      """
      {
        "id": 1,
        "creator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: event delete content emitted
      When msg delete content
      """
      {
        "id": 1,
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect event delete content
      """
      {
        "id": 1
      }
      """

    # No failing scenario - event is never emitted when message fails
