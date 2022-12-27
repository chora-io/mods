Feature: Query/Node

  Query/Node is successful when:
  - node with identifier exists

  Query/Node has the following outcomes:
  - query response returned

  Rule: The node must exist

    Background:
      Given node
      """
      {
        "id": 1,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: node exists
      When query node
      """
      {
        "id": 1
      }
      """
      Then expect no error

    Scenario: node does not exist
      When query node
      """
      {
        "id": 2
      }
      """
      Then expect the error
      """
      node with id 2: not found
      """

  Rule: The query response is returned

    Background:
      Given node
      """
      {
        "id": 1,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: query response returned
      When query node
      """
      {
        "id": 1
      }
      """
      Then expect response
      """
      {
        "id": 1,
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - response is never returned when query fails
