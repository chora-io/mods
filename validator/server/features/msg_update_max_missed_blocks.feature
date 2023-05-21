Feature: Msg/UpdateMaxMissedBlocks

  UpdateMaxMissedBlocks is successful when:
  - authority is the authority address

  UpdateMaxMissedBlocks has the following outcomes:
  - MaxMissedBlocks is updated in state
  - EventUpdateMaxMissedBlocks is emitted
  - MsgUpdateMaxMissedBlocksResponse is returned

  Rule: The authority must be the authority address

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And max missed blocks
      """
      {
        "max_missed_blocks": 100
      }
      """

    Scenario: authority is authority address
      When msg update max missed blocks
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "max_missed_blocks": 100
      }
      """
      Then expect no error

    Scenario: authority is not authority address
      When msg update max missed blocks
      """
      {
        "authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "max_missed_blocks": 100
      }
      """
      Then expect the error
      """
      authority: expected chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: received chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: unauthorized
      """

  Rule: MaxMissedBlocks is updated in state

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And max missed blocks
      """
      {
        "max_missed_blocks": 100
      }
      """

    Scenario: max missed blocks updated
      When msg update max missed blocks
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "max_missed_blocks": 200
      }
      """
      Then expect max missed blocks
      """
      {
        "max_missed_blocks": 200
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventUpdateMaxMissedBlocks emitted

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And max missed blocks
      """
      {
        "max_missed_blocks": 100
      }
      """

    Scenario: event remove emitted
      When msg update max missed blocks
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "max_missed_blocks": 200
      }
      """
      Then expect event update max missed blocks
      """
      {
        "max_missed_blocks": 200
      }
      """

    # No failing scenario - event is never emitted when message fails

  Rule: MsgUpdateMaxMissedBlocksResponse returned

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And max missed blocks
      """
      {
        "max_missed_blocks": 100
      }
      """

    Scenario: message response returned
      When msg update max missed blocks
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "max_missed_blocks": 200
      }
      """
      Then expect response
      """
      {
        "max_missed_blocks": 200
      }
      """

    # No failing scenario - response is never returned when message fails
