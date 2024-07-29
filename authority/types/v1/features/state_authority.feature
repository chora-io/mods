Feature: Authority

  Scenario: valid authority
    Given authority
    """
    {
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When validate authority
    Then expect no error

  Scenario: an error is returned if address is empty
    Given authority
    """
    {}
    """
    When validate authority
    Then expect the error
    """
    address: empty address string is not allowed: parse error
    """
