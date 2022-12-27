Feature: MsgDelete

  Scenario: valid message
    Given message
    """
    {
      "id": 1,
      "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
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

  Scenario: an error is returned if curator is empty
    Given message
    """
    {
      "id": 1
    }
    """
    When validate message
    Then expect the error
    """
    curator: empty address string is not allowed: invalid address
    """

  Scenario: an error is returned if curator is not a bech32 address
    Given message
    """
    {
      "id": 1,
      "curator": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    curator: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """
