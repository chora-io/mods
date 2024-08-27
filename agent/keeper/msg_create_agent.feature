Feature: Msg/CreateAgent

  Msg/CreateAgent is successful when:
  - always (an error is never returned)

  Msg/CreateAgent has the following outcomes:
  - message response returned
  - Agent is added to state
  - EventCreateAgent is emitted

  Rule: An error is never returned

    Scenario: never error
      Given agent sequence
      """
      {
        "sequence": "0"
      }
      """
      When msg create agent
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect no error

    # No failing scenario - only internal errors can occur

  Rule: The message response is returned

    Scenario: message response returned
      Given agent sequence
      """
      {
        "sequence": "0"
      }
      """
      When msg create agent
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect response
      """
      {
        "address": "chora140dhknrxj0vjsn4serghtw7wydm2a6mykxmkl3lmecegk4pp32dqf6sw5n"
      }
      """

    # No failing scenario - response is never returned when message fails

  Rule: Agent is added to state

    Scenario: state agent added
      Given agent sequence
      """
      {
        "sequence": "0"
      }
      """
      When msg create agent
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect state agent
      """
      {
        "address": "q9t7TGaT2ShOsMjRdbvOI3au62Sxt2/H+84yi1Qhipo=",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventCreateAgent is emitted

    Scenario: event create emitted
      Given agent sequence
      """
      {
        "sequence": "0"
      }
      """
      When msg create agent
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect event create agent
      """
      {
        "address": "chora140dhknrxj0vjsn4serghtw7wydm2a6mykxmkl3lmecegk4pp32dqf6sw5n"
      }
      """

    # No failing scenario - event is never emitted when message fails
