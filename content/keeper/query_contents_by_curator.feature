Feature: Query/ContentsByCurator

  Query/ContentsByCurator is successful when:
  - curator is a bech32 address

  Query/ContentsByCurator has the following outcomes:
  - query response returned

  Rule: The curator address must be a valid bech32 address

    Scenario: valid curator address
      When query contents by curator
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect no error

    Scenario: invalid curator address
      When query contents by curator
      """
      {
        "curator": "foo"
      }
      """
      Then expect the error
      """
      curator: decoding bech32 failed: invalid bech32 string length 3: invalid address
      """

  Rule: The query response is returned

    Scenario: response with no contents by curator
      Given content
      """
      {
        "curator": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When query contents by curator
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "contents": [],
        "pagination": {}
      }
      """

    Scenario: response with one content by curator
      Given content
      """
      {
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When query contents by curator
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "contents": [
          {
            "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          }
        ],
        "pagination": {
          "total": 1
        }
      }
      """

    Scenario: response with two contents by curator
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
      When query contents by curator
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "contents": [
          {
            "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          },
          {
            "hash": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
          }
        ],
        "pagination": {
          "total": 2
        }
      }
      """

    # No failing scenario - response is never returned when query fails
