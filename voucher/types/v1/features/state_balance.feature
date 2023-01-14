Feature: Balance

  Scenario: valid balance
    Given balance
    """
    {
      "id": 1,
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "amount": "1.25",
      "expiration": "2021-01-01T00:00:00Z"
    }
    """
    When validate balance
    Then expect no error

  Scenario: an error is returned if id is empty
    Given balance
    """
    {}
    """
    When validate balance
    Then expect the error
    """
    id: empty or zero is not allowed: parse error
    """

  Scenario: an error is returned if address is empty
    Given balance
    """
    {
      "id": 1
    }
    """
    When validate balance
    Then expect the error
    """
    address: empty address string is not allowed: parse error
    """

  Scenario: an error is returned if amount is empty
    Given balance
    """
    {
      "id": 1,
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y="
    }
    """
    When validate balance
    Then expect the error
    """
    amount: empty string is not allowed: parse error
    """

  Scenario: an error is returned if amount is zero
    Given balance
    """
    {
      "id": 1,
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "amount": "0"
    }
    """
    When validate balance
    Then expect the error
    """
    amount: expected a positive decimal, got 0: invalid decimal string: parse error
    """

  Scenario: an error is returned if amount is negative
    Given balance
    """
    {
      "id": 1,
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "amount": "-1.25"
    }
    """
    When validate balance
    Then expect the error
    """
    amount: expected a positive decimal, got -1.25: invalid decimal string: parse error
    """

  Scenario: an error is returned if expiration is empty
    Given balance
    """
    {
      "id": 1,
      "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
      "amount": "1.25"
    }
    """
    When validate balance
    Then expect the error
    """
    expiration: empty timestamp is not allowed: parse error
    """
