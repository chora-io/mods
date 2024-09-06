Feature: Msg/UpdateAdmin

  UpdateAdmin is successful when:
  - admin is the admin account

  UpdateAdmin has the following outcomes:
  - message response returned
  - Admin is updated in state
  - EventUpdateAdmin is emitted

  Rule: The admin must be the admin account

    Background:
      Given admin
      """
      {
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
      }
      """

    Scenario: admin is admin account
      When msg update admin
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no error

    Scenario: admin is not admin account
      When msg update admin
      """
      {
        "admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "new_admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect the error
      """
      admin chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: admin account chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: unauthorized
      """

  Rule: The message response is returned

    Background:
      Given admin
      """
      {
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
      }
      """

    Scenario: message response returned
      When msg update admin
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect response
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """

    # No failing scenario - response is never returned when message fails

  Rule: Admin is updated in state

    Background:
      Given admin
      """
      {
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
      }
      """

    Scenario: state admin updated
      When msg update admin
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect state admin
      """
      {
        "admin": "hEyiXxUCaFQmkbuhO9r+QDscjIY="
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventUpdateAdmin emitted

    Background:
      Given admin
      """
      {
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
      }
      """

    Scenario: event update admin emitted
      When msg update admin
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect event update
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """

    # No failing scenario - event is never emitted when message fails
