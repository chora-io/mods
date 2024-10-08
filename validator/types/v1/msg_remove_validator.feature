Feature: MsgRemoveValidator

  Scenario: valid message
    Given message
    """
    {
      "operator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
    }
    """
    When validate message
    Then expect no error

  Scenario: an error is returned if operator is empty
    Given message
    """
    {}
    """
    When validate message
    Then expect the error
    """
    operator: empty address string is not allowed: invalid address
    """

  Scenario: an error is returned if operator is not a bech32 address
    Given message
    """
    {
      "operator": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    operator: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """

  Scenario: an error is returned if address is empty
    Given message
    """
    {
      "operator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate message
    Then expect the error
    """
    address: empty address string is not allowed: invalid address
    """

  Scenario: an error is returned if address is not a bech32 address
    Given message
    """
    {
      "operator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "address": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    address: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """
