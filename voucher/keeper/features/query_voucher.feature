Feature: Query/Voucher

  Query/Voucher is successful when:
  - voucher with identifier exists

  Query/Voucher has the following outcomes:
  - query response returned

  Rule: The voucher must exist

    Background:
      Given voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: voucher exists
      When query voucher
      """
      {
        "id": 1
      }
      """
      Then expect no error

    Scenario: voucher does not exist
      When query voucher
      """
      {
        "id": 2
      }
      """
      Then expect the error
      """
      voucher with id 2: not found
      """

  Rule: The query response is returned

    Background:
      Given voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: query response returned
      When query voucher
      """
      {
        "id": 1
      }
      """
      Then expect response
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - response is never returned when query fails
