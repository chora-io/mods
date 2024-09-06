Feature: MsgRemoveSubject

  Scenario: valid message
    Given message
    """
    {
      "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4",
      "steward": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate message
    Then expect no error

  Scenario: an error is returned if address is empty
    Given message
    """
    {}
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
      "address": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    address: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """

  Scenario: an error is returned if steward is empty
    Given message
    """
    {
      "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4"
    }
    """
    When validate message
    Then expect the error
    """
    steward: empty address string is not allowed: invalid address
    """

  Scenario: an error is returned if steward is not a bech32 address
    Given message
    """
    {
      "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4",
      "steward": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    steward: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """
