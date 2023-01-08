Feature: Query/Vouchers

  Query/Vouchers is successful when:
  - always (an error is never returned)

  Query/Vouchers has the following outcomes:
  - query response returned

  Rule: An error is never returned

    Scenario: never error
      When query vouchers
      """
      {}
      """
      Then expect no error

  Rule: The query response is returned

    Scenario: response with no vouchers
      When query vouchers
      """
      {}
      """
      Then expect response
      """
      {
        "vouchers": [],
        "pagination": {}
      }
      """

    Scenario: response with one voucher
      Given voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When query vouchers
      """
      {}
      """
      Then expect response
      """
      {
        "vouchers": [
          {
            "id": 1,
            "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          }
        ],
        "pagination": {
          "total": 1
        }
      }
      """

    Scenario: response with two vouchers
      Given voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Given voucher
      """
      {
        "id": 2,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      When query vouchers
      """
      {}
      """
      Then expect response
      """
      {
        "vouchers": [
          {
            "id": 1,
            "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          },
          {
            "id": 2,
            "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
          }
        ],
        "pagination": {
          "total": 2
        }
      }
      """

    # No failing scenario - response is never returned when query fails
