Feature: Query/Agent

  Query/Agent is successful when:
  - agent with identifier exists

  Query/Agent has the following outcomes:
  - query response returned

  Rule: The agent must exist

    Background:
      Given agent
      """
      {
        "address": "address1",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: agent exists
      When query agent
      """
      {
        "address": "address1"
      }
      """
      Then expect no error

    Scenario: agent does not exist
      When query agent
      """
      {
        "address": "address2"
      }
      """
      Then expect the error
      """
      agent with address address2: not found
      """

  Rule: The query response is returned

    Background:
      Given agent
      """
      {
        "address": "address1",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: query response returned
      When query agent
      """
      {
        "address": "address1"
      }
      """
      Then expect response
      """
      {
        "address": "address1",
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - response is never returned when query fails
