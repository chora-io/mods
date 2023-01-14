Feature: Query/BalancesByAddress

  Query/BalancesByAddress is successful when:
  - always (an error is never returned)

  Query/BalancesByAddress has the following outcomes:
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
      When query balances by address
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect no error

    Scenario: balance does not exist
      When query balances by address
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no error

  Rule: The query response is returned

    Scenario: query response returned with no amounts
      When query balances by address
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "total_amounts": [],
        "pagination": {}
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
      When query balances by address
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "total_amounts": [
          {
            "id": 1,
            "total_amount": "1.25"
          }
        ],
        "pagination": {
          "total": 1
        }
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
      And balance
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
      And balance
      """
      {
        "id": 2,
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "amount": "1.25",
        "expiration": {
          "seconds": 1609459200
        }
      }
      """
      When query balances by address
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "total_amounts": [
          {
            "id": 1,
            "total_amount": "2.50"
          },
          {
            "id": 2,
            "total_amount": "1.25"
          }
        ],
        "pagination": {
          "total": 3
        }
      }
      """

    # No failing scenario - response is never returned when query fails
