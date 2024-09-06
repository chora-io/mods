Feature: MsgDelete

  Scenario: valid message
    Given message
    """
    {
      "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
    }
    """
    When validate message
    Then expect no error

  Scenario: an error is returned if curator is empty
    Given message
    """
    {}
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
      "curator": "foo"
    }
    """
    When validate message
    Then expect the error
    """
    curator: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """

  Scenario: an error is returned if hash is empty
    Given message
    """
    {
      "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate message
    Then expect the error
    """
    hash: empty string is not allowed: invalid request
    """

  Scenario: an error is returned if hash exceeds 128 characters
    Given message
    """
    {
      "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    And hash with length "129"
    When validate message
    Then expect the error
    """
    hash: exceeds max length 128: invalid request
    """
