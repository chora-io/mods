Feature: Query/Content

  Query/Content is successful when:
  - content with identifier exists

  Query/Content has the following outcomes:
  - query response returned

  Rule: The content must exist

    Background:
      Given content
      """
      {
        "id": 1,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: content exists
      When query content
      """
      {
        "id": 1
      }
      """
      Then expect no error

    Scenario: content does not exist
      When query content
      """
      {
        "id": 2
      }
      """
      Then expect the error
      """
      content with id 2: not found
      """

  Rule: The query response is returned

    Background:
      Given content
      """
      {
        "id": 1,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: query response returned
      When query content
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
