Feature: Query/Validators

  Query/Validators is successful when:
  - always (an error is never returned)

  Query/Validators has the following outcomes:
  - query response returned

  Rule: An error is never returned

    Scenario: never error
      When query validators
      """
      {}
      """
      Then expect no error

  Rule: The query response is returned

    Scenario: response with no validators
      When query validators
      """
      {}
      """
      Then expect response
      """
      {
        "validators": [],
        "pagination": {}
      }
      """

    Scenario: response with one validator
      Given validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When query validators
      """
      {}
      """
      Then expect response
      """
      {
        "validators": [
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

    Scenario: response with two validators
      Given validator
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Given validator
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      When query validators
      """
      {}
      """
      Then expect response
      """
      {
        "validators": [
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
