Feature: MsgIssue

  Scenario: valid message
    Given message
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
    When validate message
    Then expect no error

  Scenario: an error is returned if id is empty
    Given message
    """
    {}
    """
    When validate message
    Then expect the error
    """
    id: empty or zero is not allowed: invalid request
    """

  Scenario: an error is returned if issuer is empty
    Given message
    """
    {
      "id": 1
    }
    """
    When validate message
    Then expect the error
    """
    issuer: empty address string is not allowed: invalid address
    """

  Scenario: an error is returned if issuer is not a bech32 address
    Given message
    """
    {
      "id": 1,
      "issuer": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    issuer: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """

  Scenario: an error is returned if recipient is empty
    Given message
    """
    {
      "id": 1,
      "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate message
    Then expect the error
    """
    recipient: empty address string is not allowed: invalid address
    """

  Scenario: an error is returned if recipient is not a bech32 address
    Given message
    """
    {
      "id": 1,
      "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "recipient": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    recipient: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """

  Scenario: an error is returned if amount is empty
    Given message
    """
    {
      "id": 1,
      "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
    }
    """
    When validate message
    Then expect the error
    """
    amount: empty string is not allowed: invalid request
    """

  Scenario: an error is returned if amount is zero
    Given message
    """
    {
      "id": 1,
      "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
      "amount": "0"
    }
    """
    When validate message
    Then expect the error
    """
    amount: expected a positive decimal, got 0: invalid decimal string: invalid request
    """

  Scenario: an error is returned if amount is negative
    Given message
    """
    {
      "id": 1,
      "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
      "amount": "-1.25"
    }
    """
    When validate message
    Then expect the error
    """
    amount: expected a positive decimal, got -1.25: invalid decimal string: invalid request
    """

  Scenario: an error is returned if expiration is empty
    Given message
    """
    {
      "id": 1,
      "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
      "amount": "1.25"
    }
    """
    When validate message
    Then expect the error
    """
    expiration: empty timestamp is not allowed: invalid request
    """

  Scenario: an error is returned if metadata is empty
    Given message
    """
    {
      "id": 1,
      "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
      "amount": "1.25",
      "expiration": "2021-01-01T00:00:00Z"
    }
    """
    When validate message
    Then expect the error
    """
    metadata: empty string is not allowed: invalid request
    """

  Scenario: an error is returned if metadata exceeds 128 characters
    Given message
    """
    {
      "id": 1,
      "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "recipient": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
      "amount": "1.25",
      "expiration": "2021-01-01T00:00:00Z"
    }
    """
    And metadata with length "129"
    When validate message
    Then expect the error
    """
    metadata: exceeds max length 128: invalid request
    """
