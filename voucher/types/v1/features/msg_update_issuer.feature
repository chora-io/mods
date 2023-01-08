Feature: MsgUpdateIssuer

  Scenario: valid message
    Given message
    """
    {
      "id": 1,
      "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "new_issuer": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
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

  Scenario: an error is returned if new issuer is empty
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
    new issuer: empty address string is not allowed: invalid address
    """

  Scenario: an error is returned if new issuer is not a bech32 address
    Given message
    """
    {
      "id": 1,
      "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "new_issuer": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    new issuer: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """
