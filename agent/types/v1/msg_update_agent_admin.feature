Feature: MsgUpdateAgentAdmin

  Scenario: valid message
    Given message
    """
    {
      "address": "address",
      "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "new_admin": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
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

  Scenario: an error is returned if admin is empty
    Given message
    """
    {
      "address": "address"
    }
    """
    When validate message
    Then expect the error
    """
    admin: decoding bech32 failed: invalid bech32 string length 7: invalid address
    """

  Scenario: an error is returned if admin is not a bech32 address
    Given message
    """
    {
      "address": "address",
      "admin": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    admin: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """

  Scenario: an error is returned if new admin is empty
    Given message
    """
    {
      "address": "address",
      "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate message
    Then expect the error
    """
    new admin: empty address string is not allowed: invalid address
    """

  Scenario: an error is returned if new admin is not a bech32 address
    Given message
    """
    {
      "address": "address",
      "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "new_admin": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    new admin: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """
