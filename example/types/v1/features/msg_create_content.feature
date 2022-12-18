Feature: MsgCreateContent

  Scenario: a valid message
    Given the message
    """
    {
      "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
    }
    """
    When the message is validated
    Then expect no error

  Scenario: an error is returned if creator is empty
    Given the message
    """
    {}
    """
    When the message is validated
    Then expect the error
    """
    creator cannot be empty: invalid request
    """

  Scenario: an error is returned if creator is not a bech32 address
    Given the message
    """
    {
      "creator": "foo"
    }
    """
    When the message is validated
    Then expect the error
    """
    creator: decoding bech32 failed: invalid bech32 string length 3: invalid address
    """

  Scenario: an error is returned if hash is empty
    Given the message
    """
    {
      "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When the message is validated
    Then expect the error
    """
    hash: cannot be empty: invalid request
    """

  Scenario: an error is returned if hash exceeds 128 characters
    Given the message
    """
    {
      "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    And hash with length "129"
    When the message is validated
    Then expect the error
    """
    hash: exceeds max length 128: invalid request
    """
