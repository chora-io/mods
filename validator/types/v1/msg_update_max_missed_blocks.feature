Feature: MsgUpdatePolicy

  Scenario: valid message
    Given message
    """
    {
      "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "signedBlocksWindow": 100,
      "minSignedPerWindow": 100
    }
    """
    When validate message
    Then expect no error

  Scenario: valid message with zero values
    Given message
    """
    {
      "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate message
    Then expect no error

  Scenario: an error is returned if admin is empty
    Given message
    """
    {}
    """
    When validate message
    Then expect the error
    """
    admin: empty address string is not allowed: invalid address
    """

  Scenario: an error is returned if admin is not a bech32 address
    Given message
    """
    {
      "admin": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    admin: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """
