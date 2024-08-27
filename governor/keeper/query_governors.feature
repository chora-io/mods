Feature: Query/Governors

  Query/Governors is successful when:
  - always (an error is never returned)

  Query/Governors has the following outcomes:
  - query response returned

  Rule: An error is never returned

    Scenario: never error
      When query governors
      """
      {}
      """
      Then expect no error

  Rule: The query response is returned

    Scenario: response with no governors
      When query governors
      """
      {}
      """
      Then expect response
      """
      {
        "governors": [],
        "pagination": {}
      }
      """

    Scenario: response with one governor
      Given governor
      """
      {
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When query governors
      """
      {}
      """
      Then expect response
      """
      {
        "governors": [
          {
            "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          }
        ],
        "pagination": {
          "total": 1
        }
      }
      """

    Scenario: response with two governors
      Given governor
      """
      {
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Given governor
      """
      {
        "address": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      When query governors
      """
      {}
      """
      Then expect response
      """
      {
        "governors": [
          {
            "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          },
          {
            "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
            "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
          }
        ],
        "pagination": {
          "total": 2
        }
      }
      """

    # No failing scenario - response is never returned when query fails
