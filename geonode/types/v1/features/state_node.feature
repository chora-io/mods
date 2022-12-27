Feature: Content

  Scenario: valid node
    Given node
    """
    {
      "id": 1,
      "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
    }
    """
    When validate node
    Then expect no error

  Scenario: an error is returned if id is empty
    Given node
    """
    {}
    """
    When validate node
    Then expect the error
    """
    id: empty or zero is not allowed: parse error
    """

  Scenario: an error is returned if curator is empty
    Given node
    """
    {
      "id": 1
    }
    """
    When validate node
    Then expect the error
    """
    curator: empty address string is not allowed: parse error
    """

  Scenario: an error is returned if metadata is empty
    Given node
    """
    {
      "id": 1,
      "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When validate node
    Then expect the error
    """
    metadata: empty string is not allowed: parse error
    """

  Scenario: an error is returned if metadata exceeds 128 characters
    Given node
    """
    {
      "id": 1,
      "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    And metadata with length "129"
    When validate node
    Then expect the error
    """
    metadata: exceeds max length 128: parse error
    """
