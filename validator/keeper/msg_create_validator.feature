Feature: Msg/CreateValidator

  Msg/CreateValidator is successful when:
  - always (an error is never returned)

  Msg/CreateValidator has the following outcomes:
  - Validator is added to state
  - EventCreateValidator is emitted
  - MsgCreateValidatorResponse is returned

  Rule: The operator must be the validator operator

    Scenario: never error
      When msg create validator
      """
      {
        "operator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect no error

  Rule: Validator is added to state

    Scenario: state validator added
      When msg create validator
      """
      {
        "operator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect state validator
      """
      {
        "operator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventCreateValidator is emitted

    Scenario: event add emitted
      When msg create validator
      """
      {
        "operator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect event add
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """

    # No failing scenario - event is never emitted when message fails

  Rule: MsgCreateValidatorResponse is returned

    Scenario: message response returned
      When msg create validator
      """
      {
        "operator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect response
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """

    # No failing scenario - response is never returned when message fails
