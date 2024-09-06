Feature: Query/SubjectsBySteward

  Query/SubjectsBySteward is successful when:
  - subject is a bech32 address

  Query/SubjectsBySteward has the following outcomes:
  - query response returned

  Rule: The subject address must be a valid bech32 address

    Scenario: valid subject address
      When query subjects by steward
      """
      {
        "steward": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect no error

    Scenario: invalid subject address
      When query subjects by steward
      """
      {
        "steward": "foo"
      }
      """
      Then expect the error
      """
      steward: decoding bech32 failed: invalid bech32 string length 3: invalid address
      """

  Rule: The query response is returned

    Background: subject by different subject
      Given subject
      """
      {
        "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4",
        "steward": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: response with no subjects by steward
      When query subjects by steward
      """
      {
        "steward": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "steward": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "subjects": [],
        "pagination": {}
      }
      """

    Scenario: response with one subject by subject
      Given subject
      """
      {
        "address": "G+ksLYTNBuzyqdTij+Xkx1ztGDzOMACTUcjF6iEkiH0=",
        "steward": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When query subjects by steward
      """
      {
        "steward": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "steward": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "subjects": [
          {
            "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4",
            "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          }
        ],
        "pagination": {
          "total": 1
        }
      }
      """

    Scenario: response with two subjects by steward
      Given subject
      """
      {
        "address": "G+ksLYTNBuzyqdTij+Xkx1ztGDzOMACTUcjF6iEkiH0=",
        "steward": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Given subject
      """
      {
        "address": "lxDqH19hS9D1Nom8zAvt3vUFS4136FKa6bluwHhbetE=",
        "steward": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      When query subjects by steward
      """
      {
        "steward": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "steward": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "subjects": [
          {
            "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4",
            "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          },
          {
            "address": "chora1jugw586lv99apafk3x7vczldmm6s2judwl599xhfh9hvq7zm0tgs7jccnn",
            "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
          }
        ],
        "pagination": {
          "total": 2
        }
      }
      """

    # No failing scenario - response is never returned when query fails
