Feature: Msg/AddValidator

  Msg/AddValidator is successful when:
  - admin is the admin address

  Msg/AddValidator has the following outcomes:
  - Validator is added to state
  - EventAddValidator is emitted
  - MsgAddValidatorResponse is returned

  Rule: The admin must be the admin address

    Background:
      Given admin "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"

    Scenario: admin is admin address
      When msg add validator
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect no error

    Scenario: admin is not admin address
      When msg add validator
      """
      {
        "admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect the error
      """
      admin: expected chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: received chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: unauthorized
      """

  Rule: Validator is added to state

    Background:
      Given admin "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"

    Scenario: state validator added
      When msg add validator
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect state validator
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventAddValidator is emitted

    Scenario: event add emitted
      When msg add validator
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect event add
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """

    # No failing scenario - event is never emitted when message fails

  Rule: MsgAddValidatorResponse is returned

    Background:
      Given admin "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"

    Scenario: message response returned
      When msg add validator
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect response
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """

    # No failing scenario - response is never returned when message fails
