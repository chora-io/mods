Feature: Msg/Remove

  Remove is successful when:
  - authority is the authority address
  - validator with address exists

  Remove has the following outcomes:
  - message response returned
  - Validator is removed from state
  - EventRemove is emitted

  Rule: The authority must be the authority address

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And validator
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: authority is authority address
      When msg remove
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no error

    Scenario: authority is not authority address
      When msg remove
      """
      {
        "authority": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect the error
      """
      authority: expected chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: received chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: unauthorized
      """

  Rule: The validator must exist

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And validator
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: validator exists
      When msg remove
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no error

    Scenario: validator does not exist
      When msg remove
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect the error
      """
      validator with address chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: not found: not found
      """

  Rule: The message response is returned

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And validator
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: message response returned
      When msg remove
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect response
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """

    # No failing scenario - response is never returned when message fails

  Rule: Validator is removed from state

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And validator
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: state validator removed
      When msg remove
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no validator with address "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"

    # No failing scenario - state is never updated when message fails

  Rule: EventRemove emitted

    Background:
      Given authority "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      And validator
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: event remove emitted
      When msg remove
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect event remove
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """

    # No failing scenario - event is never emitted when message fails
