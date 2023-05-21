Feature: MsgUpdateMaxMissedBlocks

  Scenario: valid message
    Given message
    """
    {
      "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "maxMissedBlocks": 100
    }
    """
    When validate message
    Then expect no error

  Scenario: valid message with zero value
    Given message
    """
    {
      "authority": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate message
    Then expect no error

  Scenario: an error is returned if authority is empty
    Given message
    """
    {}
    """
    When validate message
    Then expect the error
    """
    authority: empty address string is not allowed: invalid address
    """

  Scenario: an error is returned if authority is not a bech32 address
    Given message
    """
    {
      "authority": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    authority: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """
