Feature: ValidatorMissedBlocks

  Scenario: valid validator missed blocks
    Given validator missed blocks
    """
    {
      "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "missedBlocks": 10
    }
    """
    When validate validator missed blocks
    Then expect no error

  Scenario: valid validator missed blocks with zero value
    Given validator missed blocks
    """
    {
      "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate validator missed blocks
    Then expect no error

  Scenario: an error is returned if address is empty
    Given validator missed blocks
    """
    {}
    """
    When validate validator missed blocks
    Then expect the error
    """
    address: empty address string is not allowed: parse error
    """

  Scenario: an error is returned if address is not a bech32 address
    Given validator missed blocks
    """
    {
      "address": "foo"
    }
    """
    When validate validator missed blocks
    Then expect the error
    """
    address: decoding bech32 failed: invalid bech32 string length 3: parse error
    """
