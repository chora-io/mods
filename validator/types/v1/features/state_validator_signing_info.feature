Feature: ValidatorSigningInfo

  Scenario: valid validator signing info
    Given validator signing info
    """
    {
      "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "missedBlocksCount": 10
    }
    """
    When validate validator signing info
    Then expect no error

  Scenario: valid validator signing info with zero value
    Given validator signing info
    """
    {
      "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate validator signing info
    Then expect no error

  Scenario: an error is returned if address is empty
    Given validator signing info
    """
    {}
    """
    When validate validator signing info
    Then expect the error
    """
    address: empty address string is not allowed: parse error
    """

  Scenario: an error is returned if address is not a bech32 address
    Given validator signing info
    """
    {
      "address": "foo"
    }
    """
    When validate validator signing info
    Then expect the error
    """
    address: decoding bech32 failed: invalid bech32 string length 3: parse error
    """
