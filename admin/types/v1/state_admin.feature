Feature: Admin

  Scenario: valid admin
    Given admin
    """
    {
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When validate admin
    Then expect no error

  Scenario: an error is returned if address is empty
    Given admin
    """
    {}
    """
    When validate admin
    Then expect the error
    """
    address: empty address string is not allowed: parse error
    """
