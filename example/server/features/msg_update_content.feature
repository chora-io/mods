Feature: Msg/UpdateContent

  UpdateContent is successful when:
  - creator is the content creator

  UpdateContent has the following outcomes:
  - message response returned
  - Content is updated in state
  - EventUpdateContent is emitted

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
      When msg update content
      """
      {
        "id": 1,
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_hash": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect no error

    Scenario: creator is not content creator
      When msg update content
      """
      {
        "id": 1,
        "creator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "new_hash": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
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
      When msg update content
      """
      {
        "id": 1,
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_hash": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect response
      """
      {
        "id": 1
      }
      """

    # No failing scenario - response is never returned when message fails

  Rule: Content is updated in state

    Background:
      Given content
      """
      {
        "id": 1,
        "creator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: state content updated
      When msg update content
      """
      {
        "id": 1,
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_hash": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect state content
      """
      {
        "id": 1,
        "creator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventUpdateContent emitted

    Background:
      Given content
      """
      {
        "id": 1,
        "creator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: event update content emitted
      When msg update content
      """
      {
        "id": 1,
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_hash": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect event update content
      """
      {
        "id": 1
      }
      """

    # No failing scenario - event is never emitted when message fails
