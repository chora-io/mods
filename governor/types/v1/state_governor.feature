Feature: Governor

  Scenario: valid governor
    Given governor
    """
    {
      "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
      "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
    }
    """
    When validate governor
    Then expect no error

  Scenario: an error is returned if address is empty
    Given governor
    """
    {}
    """
    When validate governor
    Then expect the error
    """
    address: empty address string is not allowed: parse error
    """

  Scenario: an error is returned if address is not a bech32 address
    Given governor
    """
    {
      "address": "foo"
    }
    """
    When validate governor
    Then expect the error
    """
    address: decoding bech32 failed: invalid bech32 string length 3: parse error
    """

  Scenario: an error is returned if metadata is empty
    Given governor
    """
    {
      "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    When validate governor
    Then expect the error
    """
    metadata: empty string is not allowed: parse error
    """

  Scenario: an error is returned if metadata exceeds 128 characters
    Given governor
    """
    {
      "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
    }
    """
    And metadata with length "129"
    When validate governor
    Then expect the error
    """
    metadata: exceeds max length 128: parse error
    """
