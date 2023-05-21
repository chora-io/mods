Feature: Msg/Add

  Msg/Add is successful when:
  - authority is the authority address

  Msg/Add has the following outcomes:
  - message response returned
  - Validator is added to state
  - EventAdd is emitted

  Rule: The authority must be the authority address

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"

    Scenario: authority is authority address
      When msg add
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect no error

    Scenario: authority is not authority address
      When msg add
      """
      {
        "authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect the error
      """
      authority: expected chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: received chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: unauthorized
      """

  Rule: The message response is returned

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"

    Scenario: message response returned
      When msg add
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
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

  Rule: Validator is added to state

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"

    Scenario: state validator added
      When msg add
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
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

  Rule: EventAdd is emitted

    Scenario: event add emitted
      When msg add
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
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
