Feature: Query/ContentByCurator

  Query/ContentByCurator is successful when:
  - curator is a bech32 address

  Query/ContentByCurator has the following outcomes:
  - query response returned

  Rule: The curator address must be a valid bech32 address

    Scenario: valid curator address
      When query content by curator
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect no error

    Scenario: invalid curator address
      When query content by curator
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

    Background: content by different curator
      Given content
      """
      {
        "id": 1,
        "curator": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: response with no content by curator
      When query content by curator
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "content": [],
        "pagination": {}
      }
      """

    Scenario: response with one content by curator
      Given content
      """
      {
        "id": 2,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When query content by curator
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "content": [
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

    Scenario: response with two content by curator
      Given content
      """
      {
        "id": 2,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Given content
      """
      {
        "id": 3,
        "curator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      When query content by curator
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "curator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "content": [
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
