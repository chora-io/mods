Feature: Query/Policy

  Query/Policy is successful when:
  - always (an error is never returned)

  Query/Policy has the following outcomes:
  - query response returned

  Rule: An error is never returned

    Scenario: never error
      Given policy
      """
      {
        "signedBlocksWindow": 100
      }
      """
      When query policy
      Then expect no error

    Scenario: never error with zero value
      When query policy
      Then expect no error

    # No failing scenario - only internal errors can occur

  Rule: The query response is returned

    Scenario: query response returned
      Given policy
      """
      {
        "signedBlocksWindow": 100
      }
      """
      When query policy
      Then expect response
      """
      {
        "signedBlocksWindow": 100
      }
      """

    Scenario: query response returned with zero value
      Given policy
      """
      {
        "signedBlocksWindow": 0
      }
      """
      When query policy
      Then expect response
      """
      {
        "signedBlocksWindow": 0
      }
      """

    # No failing scenario - response is never returned when query fails
