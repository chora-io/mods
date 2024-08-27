Feature: Msg/UpdateValidator

  UpdateValidator is successful when:
  - validator with address exists

  UpdateValidator has the following outcomes:
  - Validator is updated in state
  - EventUpdateValidator is emitted
  - MsgUpdateValidatorResponse is returned

  Rule: The validator must exist

    Background:
      Given validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: validator exists
      When msg update validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect no error

    Scenario: validator does not exist
      When msg update validator
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "new_metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect the error
      """
      validator with address chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: not found: not found
      """

  Rule: Validator is updated in state

    Background:
      Given validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: state validator updated
      When msg update validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect state validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventUpdateValidator emitted

    Background:
      Given validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: event update validator emitted
      When msg update validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect event update validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """

    # No failing scenario - event is never emitted when message fails

  Rule: MsgUpdateValidatorResponse is returned

    Background:
      Given validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: message response returned
      When msg update validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      Then expect response
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """

    # No failing scenario - response is never returned when message fails
