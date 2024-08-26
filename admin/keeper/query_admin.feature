Feature: Query/Admin

  Query/Admin is successful when:
  - admin with identifier exists

  Query/Admin has the following outcomes:
  - query response returned

  Rule: The query response is returned

    Background:
      Given admin
      """
      {
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
      }
      """

    Scenario: query response returned
      When query admin
      """
      {}
      """
      Then expect response
      """
      {
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """

    # No failing scenario - response is never returned when query fails
