Feature: Query/Authority

  Query/Authority is successful when:
  - authority with identifier exists

  Query/Authority has the following outcomes:
  - query response returned

  Rule: The query response is returned

    Background:
      Given authority
      """
      {
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
      }
      """

    Scenario: query response returned
      When query authority
      """
      {}
      """
      Then expect response
      """
      {
        "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """

    # No failing scenario - response is never returned when query fails
