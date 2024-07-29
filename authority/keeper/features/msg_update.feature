Feature: Msg/Update

  Update is successful when:
  - authority is the authority account

  Update has the following outcomes:
  - message response returned
  - Authority is updated in state
  - EventUpdate is emitted

  Rule: The authority must be the authority account

    Background:
      Given authority
      """
      {
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
      }
      """

    Scenario: authority is authority account
      When msg update
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no error

    Scenario: authority is not authority account
      When msg update
      """
      {
        "authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "new_authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect the error
      """
      authority chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: authority account chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: unauthorized
      """

  Rule: The message response is returned

    Background:
      Given authority
      """
      {
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
      }
      """

    Scenario: message response returned
      When msg update
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect response
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """

    # No failing scenario - response is never returned when message fails

  Rule: Authority is updated in state

    Background:
      Given authority
      """
      {
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
      }
      """

    Scenario: state authority updated
      When msg update
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect state authority
      """
      {
        "address": "hEyiXxUCaFQmkbuhO9r+QDscjIY="
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventUpdate emitted

    Background:
      Given authority
      """
      {
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
      }
      """

    Scenario: event update authority emitted
      When msg update
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect event update
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """

    # No failing scenario - event is never emitted when message fails
