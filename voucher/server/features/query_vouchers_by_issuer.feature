Feature: Query/VouchersByIssuer

  Query/VouchersByIssuer is successful when:
  - issuer is a bech32 address

  Query/VouchersByIssuer has the following outcomes:
  - query response returned

  Rule: The issuer address must be a valid bech32 address

    Scenario: valid issuer address
      When query vouchers by issuer
      """
      {
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect no error

    Scenario: invalid issuer address
      When query vouchers by issuer
      """
      {
        "issuer": "foo"
      }
      """
      Then expect the error
      """
      issuer: decoding bech32 failed: invalid bech32 string length 3: invalid address
      """

  Rule: The query response is returned

    Background: voucher by different issuer
      Given voucher
      """
      {
        "id": 1,
        "issuer": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: response with no vouchers by issuer
      When query vouchers by issuer
      """
      {
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "vouchers": [],
        "pagination": {}
      }
      """

    Scenario: response with one voucher by issuer
      Given voucher
      """
      {
        "id": 2,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When query vouchers by issuer
      """
      {
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "vouchers": [
          {
            "id": 2,
            "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          }
        ],
        "pagination": {
          "total": 1
        }
      }
      """

    Scenario: response with two vouchers by issuer
      Given voucher
      """
      {
        "id": 2,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Given voucher
      """
      {
        "id": 3,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      When query vouchers by issuer
      """
      {
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "vouchers": [
          {
            "id": 2,
            "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          },
          {
            "id": 3,
            "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
          }
        ],
        "pagination": {
          "total": 2
        }
      }
      """

    # No failing scenario - response is never returned when query fails
