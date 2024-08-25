Feature: Msg/UpdatePolicy

  UpdatePolicy is successful when:
  - admin is the admin address

  UpdatePolicy has the following outcomes:
  - Policy is updated in state
  - EventUpdatePolicy is emitted
  - MsgUpdatePolicyResponse is returned

  Rule: The admin must be the admin address

    Background:
      Given admin "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And policy
      """
      {
        "signedBlocksWindow": 100
      }
      """

    Scenario: admin is admin address
      When msg update policy
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "signedBlocksWindow": 100
      }
      """
      Then expect no error

    Scenario: admin is not admin address
      When msg update policy
      """
      {
        "admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "signedBlocksWindow": 100
      }
      """
      Then expect the error
      """
      admin: expected chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: received chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: unauthorized
      """

  Rule: Policy is updated in state

    Background:
      Given admin "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And policy
      """
      {
        "signedBlocksWindow": 100
      }
      """

    Scenario: policy updated
      When msg update policy
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "signedBlocksWindow": 200
      }
      """
      Then expect policy
      """
      {
        "signedBlocksWindow": 200
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventUpdatePolicy emitted

    Background:
      Given admin "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And policy
      """
      {
        "signedBlocksWindow": 100
      }
      """

    Scenario: event remove emitted
      When msg update policy
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "signedBlocksWindow": 200
      }
      """
      Then expect event update policy
      """
      {
        "signedBlocksWindow": 200
      }
      """

    # No failing scenario - event is never emitted when message fails

  Rule: MsgUpdatePolicyResponse returned

    Background:
      Given admin "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And policy
      """
      {
        "signedBlocksWindow": 100
      }
      """

    Scenario: message response returned
      When msg update policy
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "signedBlocksWindow": 200
      }
      """
      Then expect response
      """
      {
        "signedBlocksWindow": 200
      }
      """

    # No failing scenario - response is never returned when message fails
