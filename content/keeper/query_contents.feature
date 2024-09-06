Feature: Query/Contents

  Query/Contents is successful when:
  - always (an error is never returned)

  Query/Contents has the following outcomes:
  - query response returned

  Rule: An error is never returned

    Scenario: never error
      When query contents
      """
      {}
      """
      Then expect no error

  Rule: The query response is returned

    Scenario: response with no contents
      When query contents
      """
      {}
      """
      Then expect response
      """
      {
        "contents": [],
        "pagination": {}
      }
      """

    Scenario: response with one content
      Given content
      """
      {
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When query contents
      """
      {}
      """
      Then expect response
      """
      {
        "contents": [
          {
            "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          }
        ],
        "pagination": {
          "total": 1
        }
      }
      """

    Scenario: response with two contents
      Given content
      """
      {
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Given content
      """
      {
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      When query contents
      """
      {}
      """
      Then expect response
      """
      {
        "contents": [
          {
            "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          },
          {
            "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "hash": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
          }
        ],
        "pagination": {
          "total": 2
        }
      }
      """

    # No failing scenario - response is never returned when query fails
