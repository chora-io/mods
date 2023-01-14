Feature: Query/BalancesByVoucher

  Query/BalancesByVoucher is successful when:
  - always (an error is never returned)

  Query/BalancesByVoucher has the following outcomes:
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
      When query balances by voucher
      """
      {
        "id": 1
      }
      """
      Then expect no error

    Scenario: balance does not exist
      When query balances by voucher
      """
      {
        "id": 2
      }
      """
      Then expect no error

  Rule: The query response is returned

    Scenario: query response returned with no amounts
      When query balances by voucher
      """
      {
        "id": 1
      }
      """
      Then expect response
      """
      {
        "id": 1,
        "total_amounts": [],
        "pagination": {}
      }
      """

    Scenario: query response returned with one address
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
      When query balances by voucher
      """
      {
        "id": 1
      }
      """
      Then expect response
      """
      {
        "id": 1,
        "total_amounts": [
          {
            "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "total_amount": "1.25"
          }
        ],
        "pagination": {
          "total": 1
        }
      }
      """

    Scenario: query response returned with multiple addresses
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
        "id": 1,
        "address": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "amount": "1.25",
        "expiration": {
          "seconds": 1640995200
        }
      }
      """
      When query balances by voucher
      """
      {
        "id": 1
      }
      """
      Then expect response
      """
      {
        "id": 1,
        "total_amounts": [
          {
            "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "total_amount": "2.50"
          },
          {
            "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
            "total_amount": "1.25"
          }
        ],
        "pagination": {
          "total": 3
        }
      }
      """

    # No failing scenario - response is never returned when query fails
