Feature: Msg/UpdatePolicy

  UpdatePolicy is successful when:
  - authority is the authority address

  UpdatePolicy has the following outcomes:
  - Policy is updated in state
  - EventUpdatePolicy is emitted
  - MsgUpdatePolicyResponse is returned

  Rule: The authority must be the authority address

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And policy
      """
      {
        "signedBlocksWindow": 100
      }
      """

    Scenario: authority is authority address
      When msg update policy
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "signedBlocksWindow": 100
      }
      """
      Then expect no error

    Scenario: authority is not authority address
      When msg update policy
      """
      {
        "authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "signedBlocksWindow": 100
      }
      """
      Then expect the error
      """
      authority: expected chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: received chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: unauthorized
      """

  Rule: Policy is updated in state

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
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
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
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
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
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
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
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
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
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
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
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
