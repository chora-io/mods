Feature: Msg/UpdateContentCurator

  UpdateContentCurator is successful when:
  - curator is the content curator

  UpdateContentCurator has the following outcomes:
  - message response returned
  - Content is updated in state
  - EventUpdateContentCurator is emitted

  Rule: The curator must be the content curator

    Background:
      Given content
      """
      {
        "id": 1,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: curator is content curator
      When msg update content curator
      """
      {
        "id": 1,
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_curator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no error

    Scenario: curator is not content curator
      When msg update content curator
      """
      {
        "id": 1,
        "curator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "new_curator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect the error
      """
      curator chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: content curator chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: unauthorized
      """

  Rule: The message response is returned

    Background:
      Given content
      """
      {
        "id": 1,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: message response returned
      When msg update content curator
      """
      {
        "id": 1,
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_curator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
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
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: state content updated
      When msg update content curator
      """
      {
        "id": 1,
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_curator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect state content
      """
      {
        "id": 1,
        "curator": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventUpdateContentCurator emitted

    Background:
      Given content
      """
      {
        "id": 1,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: event update curator emitted
      When msg update content curator
      """
      {
        "id": 1,
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_curator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect event update content curator
      """
      {
        "id": 1
      }
      """

    # No failing scenario - event is never emitted when message fails
