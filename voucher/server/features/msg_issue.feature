Feature: Msg/Issue

  Msg/Issue is successful when:
  - issuer is the voucher issuer
  - expiration must be in the future

  Msg/Issue has the following outcomes:
  - message response returned
  - Balance is added to or updated in state
  - EventIssue is emitted

  Rule: The issuer must be the voucher issuer

    Background:
      Given block time "2020-01-01T00:00:00Z"
      And voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: issuer is voucher issuer
      When msg issue
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "amount": "1.25",
        "expiration": "2021-01-01T00:00:00Z"
      }
      """
      Then expect no error

    Scenario: issuer is not voucher issuer
      When msg issue
      """
      {
        "id": 1,
        "issuer": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "amount": "1.25",
        "expiration": "2021-01-01T00:00:00Z"
      }
      """
      Then expect the error
      """
      issuer chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: voucher issuer chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: unauthorized
      """

  Rule: The expiration must be in the future

    Background:
      Given block time "2020-01-01T00:00:00Z"
      And voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: expiration is greater than current block time
      When msg issue
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "amount": "1.25",
        "expiration": "2021-01-01T00:00:00Z"
      }
      """
      Then expect no error

    Scenario: expiration is equal to current block time
      When msg issue
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "amount": "1.25",
        "expiration": "2020-01-01T00:00:00Z"
      }
      """
      Then expect the error
      """
      expiration must be in the future: received 2020-01-01 00:00:00 +0000 UTC: invalid request
      """

    Scenario: expiration is less than current block time
      When msg issue
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "amount": "1.25",
        "expiration": "2019-01-01T00:00:00Z"
      }
      """
      Then expect the error
      """
      expiration must be in the future: received 2019-01-01 00:00:00 +0000 UTC: invalid request
      """

  Rule: The message response is returned

    Background:
      Given block time "2020-01-01T00:00:00Z"
      And voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: message response returned
      When msg issue
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "amount": "1.25",
        "expiration": "2021-01-01T00:00:00Z"
      }
      """
      Then expect response
      """
      {
        "id": 1
      }
      """

    # No failing scenario - response is never returned when message fails

  Rule: Balance is added to or updated in state

    Background:
      Given block time "2020-01-01T00:00:00Z"
      And voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: state balance added
      When msg issue
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "amount": "1.25",
        "expiration": "2021-01-01T00:00:00Z"
      }
      """
      Then expect state balance
      """
      {
        "id": 1,
        "address": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "amount": "1.25",
        "expiration": {
          "seconds": 1609459200
        }
      }
      """

    Scenario: state balance updated
      Given balance
      """
      {
        "id": 1,
        "address": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "amount": "1.25",
        "expiration": {
          "seconds": 1609459200
        }
      }
      """
      When msg issue
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "amount": "1.25",
        "expiration": "2021-01-01T00:00:00Z"
      }
      """
      Then expect state balance
      """
      {
        "id": 1,
        "address": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "amount": "2.50",
        "expiration": {
          "seconds": 1609459200
        }
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventIssue is emitted

    Background:
      Given block time "2020-01-01T00:00:00Z"
      And voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: event issue emitted
      When msg issue
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "amount": "1.25",
        "expiration": "2021-01-01T00:00:00Z"
      }
      """
      Then expect event issue
      """
      {
        "id": 1
      }
      """

    Scenario: event issue emitted with metadata
      When msg issue
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "amount": "1.25",
        "expiration": "2021-01-01T00:00:00Z",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Then expect event issue
      """
      {
        "id": 1,
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - event is never emitted when message fails
