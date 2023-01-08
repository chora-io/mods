Feature: Content

  Scenario: valid voucher
    Given voucher
    """
    {
      "id": 1,
      "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
    }
    """
    When validate voucher
    Then expect no error

  Scenario: an error is returned if id is empty
    Given voucher
    """
    {}
    """
    When validate voucher
    Then expect the error
    """
    id: empty or zero is not allowed: parse error
    """

  Scenario: an error is returned if issuer is empty
    Given voucher
    """
    {
      "id": 1
    }
    """
    When validate voucher
    Then expect the error
    """
    issuer: empty address string is not allowed: parse error
    """

  Scenario: an error is returned if metadata is empty
    Given voucher
    """
    {
      "id": 1,
      "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When validate voucher
    Then expect the error
    """
    metadata: empty string is not allowed: parse error
    """

  Scenario: an error is returned if metadata exceeds 128 characters
    Given voucher
    """
    {
      "id": 1,
      "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    And metadata with length "129"
    When validate voucher
    Then expect the error
    """
    metadata: exceeds max length 128: parse error
    """
