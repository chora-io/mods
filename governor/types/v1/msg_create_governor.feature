Feature: MsgCreateGovernor

  Scenario: valid message
    Given message
    """
    {
      "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
      "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
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

  Scenario: an error is returned if metadata is empty
    Given message
    """
    {
      "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate message
    Then expect the error
    """
    metadata: empty string is not allowed: invalid request
    """

  Scenario: an error is returned if metadata exceeds 128 characters
    Given message
    """
    {
      "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    And metadata with length "129"
    When validate message
    Then expect the error
    """
    metadata: exceeds max length 128: invalid request
    """
