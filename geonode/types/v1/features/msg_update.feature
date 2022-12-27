Feature: MsgUpdate

  Scenario: valid message
    Given message
    """
    {
      "id": 1,
      "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "new_metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
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

  Scenario: an error is returned if new metadata is empty
    Given message
    """
    {
      "id": 1,
      "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate message
    Then expect the error
    """
    new metadata: empty string is not allowed: invalid request
    """

  Scenario: an error is returned if new metadata exceeds 128 characters
    Given message
    """
    {
      "id": 1,
      "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    And new metadata with length "129"
    When validate message
    Then expect the error
    """
    new metadata: exceeds max length 128: invalid request
    """
