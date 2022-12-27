Feature: Content

  Scenario: valid content
    Given content
    """
    {
      "id": 1,
      "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
    }
    """
    When validate content
    Then expect no error

  Scenario: an error is returned if id is empty
    Given content
    """
    {}
    """
    When validate content
    Then expect the error
    """
    id: empty or zero is not allowed: parse error
    """

  Scenario: an error is returned if curator is empty
    Given content
    """
    {
      "id": 1
    }
    """
    When validate content
    Then expect the error
    """
    curator: empty address string is not allowed: parse error
    """

  Scenario: an error is returned if hash is empty
    Given content
    """
    {
      "id": 1,
      "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When validate content
    Then expect the error
    """
    hash: empty string is not allowed: parse error
    """

  Scenario: an error is returned if hash exceeds 128 characters
    Given content
    """
    {
      "id": 1,
      "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    And hash with length "129"
    When validate content
    Then expect the error
    """
    hash: exceeds max length 128: parse error
    """
