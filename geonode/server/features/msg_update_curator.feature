Feature: Msg/UpdateCurator

  UpdateCurator is successful when:
  - curator is the node curator

  UpdateCurator has the following outcomes:
  - message response returned
  - Content is updated in state
  - EventUpdateCurator is emitted

  Rule: The curator must be the node curator

    Background:
      Given node
      """
      {
        "id": 1,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: curator is node curator
      When msg update curator
      """
      {
        "id": 1,
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_curator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no error

    Scenario: curator is not node curator
      When msg update curator
      """
      {
        "id": 1,
        "curator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "new_curator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect the error
      """
      curator chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: node curator chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: unauthorized
      """

  Rule: The message response is returned

    Background:
      Given node
      """
      {
        "id": 1,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: message response returned
      When msg update curator
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
      Given node
      """
      {
        "id": 1,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: state node updated
      When msg update curator
      """
      {
        "id": 1,
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_curator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect state node
      """
      {
        "id": 1,
        "curator": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventUpdateCurator emitted

    Background:
      Given node
      """
      {
        "id": 1,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: event update curator emitted
      When msg update curator
      """
      {
        "id": 1,
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_curator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect event update curator
      """
      {
        "id": 1
      }
      """

    # No failing scenario - event is never emitted when message fails
