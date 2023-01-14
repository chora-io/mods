Feature: Query/Balance

  Query/Balance is successful when:
  - always (an error is never returned)

  Query/Balance has the following outcomes:
  - query response returned

  Rule: An error is never returned

    Background:
      Given balance
      """
      {
        "id": 1,
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "amount": "1.25",
        "expiration": {
          "seconds": 1609459200
        }
      }
      """

    Scenario: balance exists
      When query balance
      """
      {
        "id": 1,
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect no error

    Scenario: balance does not exist
      When query balance
      """
      {
        "id": 2,
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect no error

  Rule: The query response is returned

    Scenario: query response returned with no amounts
      When query balance
      """
      {
        "id": 1,
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "id": 1,
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "total_amount": "0",
        "amounts": []
      }
      """

    Scenario: query response returned with one amount
      Given balance
      """
      {
        "id": 1,
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "amount": "1.25",
        "expiration": {
          "seconds": 1609459200
        }
      }
      """
      When query balance
      """
      {
        "id": 1,
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "id": 1,
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "total_amount": "1.25",
        "amounts": [
          {
            "amount": "1.25",
            "expiration": "2021-01-01T00:00:00Z"
          }
        ]
      }
      """

    Scenario: query response returned with multiple amounts
      Given balance
      """
      {
        "id": 1,
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "amount": "1.25",
        "expiration": {
          "seconds": 1609459200
        }
      }
      """
      Given balance
      """
      {
        "id": 1,
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "amount": "1.25",
        "expiration": {
          "seconds": 1640995200
        }
      }
      """
      When query balance
      """
      {
        "id": 1,
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "id": 1,
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "total_amount": "2.50",
        "amounts": [
          {
            "amount": "1.25",
            "expiration": "2021-01-01T00:00:00Z"
          },
          {
            "amount": "1.25",
            "expiration": "2022-01-01T00:00:00Z"
          }
        ]
      }
      """

    # No failing scenario - response is never returned when query fails
