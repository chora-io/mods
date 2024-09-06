Feature: Msg/RemoveValidator

  RemoveValidator is successful when:
  - validator with address exists

  RemoveValidator has the following outcomes:
  - Validator is removed from state
  - ValidatorSigningInfo is removed from state
  - EventRemoveValidator is emitted
  - MsgRemoveValidatorResponse is returned

  Rule: The validator must exist

    Background:
      Given validator
      """
      {
        "operator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      And validator signing info
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "missed_blocks": [],
        "missed_blocks_count": 0
      }
      """

    Scenario: validator exists
      When msg remove validator
      """
      {
        "operator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no error

    Scenario: validator does not exist
      When msg remove validator
      """
      {
        "operator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect the error
      """
      validator with address chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: not found: not found
      """

  Rule: Validator is removed from state

    Background:
      Given validator
      """
      {
        "operator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      And validator signing info
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "missed_blocks": [],
        "missed_blocks_count": 0
      }
      """

    Scenario: state validator removed
      When msg remove validator
      """
      {
        "operator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no validator with address "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"

    # No failing scenario - state is never updated when message fails

  Rule: ValidatorSigningInfo is removed from state

    Background:
      Given validator
      """
      {
        "operator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      And validator signing info
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "missed_blocks": [],
        "missed_blocks_count": 0
      }
      """

    Scenario: state validator removed
      When msg remove validator
      """
      {
        "operator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no validator signing info with address "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"

    # No failing scenario - state is never updated when message fails

  Rule: EventRemoveValidator emitted

    Background:
      Given validator
      """
      {
        "operator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      And validator signing info
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "missed_blocks": [],
        "missed_blocks_count": 0
      }
      """

    Scenario: event remove emitted
      When msg remove validator
      """
      {
        "operator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
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

  Rule: MsgRemoveValidatorResponse is returned

    Background:
      Given validator
      """
      {
        "operator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      And validator signing info
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "missed_blocks": [],
        "missed_blocks_count": 0
      }
      """

    Scenario: message response returned
      When msg remove validator
      """
      {
        "operator": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
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
