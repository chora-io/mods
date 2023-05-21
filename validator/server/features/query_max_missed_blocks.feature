Feature: Query/MaxMissedBlocks

  Query/MaxMissedBlocks is successful when:
  - always (an error is never returned)

  Query/MaxMissedBlocks has the following outcomes:
  - query response returned

  Rule: An error is never returned

    Scenario: never error
      Given max missed blocks
      """
      {
        "max_missed_blocks": 100
      }
      """
      When query max missed blocks
      Then expect no error

    Scenario: never error with zero value
      When query max missed blocks
      Then expect no error

    # No failing scenario - only internal errors can occur

  Rule: The query response is returned

    Scenario: query response returned
      Given max missed blocks
      """
      {
        "max_missed_blocks": 100
      }
      """
      When query max missed blocks
      Then expect response
      """
      {
        "max_missed_blocks": 100
      }
      """

    Scenario: query response returned with zero value
      Given max missed blocks
      """
      {
        "max_missed_blocks": 0
      }
      """
      When query max missed blocks
      Then expect response
      """
      {
        "max_missed_blocks": 0
      }
      """

    # No failing scenario - response is never returned when query fails
