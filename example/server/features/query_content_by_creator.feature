Feature: Query/ContentByCreator

  Query/ContentByCreator is successful when:
  - creator is a bech32 address

  Query/ContentByCreator has the following outcomes:
  - query response returned

  Rule: The creator address must be a valid bech32 address

    Scenario: valid creator address
      When query content by creator
      """
      {
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect no error

    Scenario: invalid creator address
      When query content by creator
      """
      {
        "creator": "foo"
      }
      """
      Then expect the error
      """
      creator: decoding bech32 failed: invalid bech32 string length 3: invalid address
      """

  Rule: The query response is returned

    Background: content by different creator
      Given content
      """
      {
        "id": 1,
        "creator": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: response with no content by creator
      When query content by creator
      """
      {
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "content": [],
        "pagination": {}
      }
      """

    Scenario: response with one content by creator
      Given content
      """
      {
        "id": 2,
        "creator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When query content by creator
      """
      {
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "content": [
          {
            "id": 2,
            "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          }
        ],
        "pagination": {
          "total": 1
        }
      }
      """

    Scenario: response with two content by creator
      Given content
      """
      {
        "id": 2,
        "creator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Given content
      """
      {
        "id": 3,
        "creator": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "hash": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      When query content by creator
      """
      {
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "creator": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "content": [
          {
            "id": 2,
            "hash": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          },
          {
            "id": 3,
            "hash": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
          }
        ],
        "pagination": {
          "total": 2
        }
      }
      """

    # No failing scenario - response is never returned when query fails
