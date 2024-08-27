Feature: Query/Agents

  Query/Agents is successful when:
  - always (an error is never returned)

  Query/Agents has the following outcomes:
  - query response returned

  Rule: An error is never returned

    Scenario: never error
      When query agents
      """
      {}
      """
      Then expect no error

  Rule: The query response is returned

    Scenario: response with no agents
      When query agents
      """
      {}
      """
      Then expect response
      """
      {
        "agents": [],
        "pagination": {}
      }
      """

    Scenario: response with one agent
      Given agent
      """
      {
        "address": "G+ksLYTNBuzyqdTij+Xkx1ztGDzOMACTUcjF6iEkiH0=",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      When query agents
      """
      {}
      """
      Then expect response
      """
      {
        "agents": [
          {
            "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4",
            "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          }
        ],
        "pagination": {
          "total": 1
        }
      }
      """

    Scenario: response with two agents
      Given agent
      """
      {
        "address": "G+ksLYTNBuzyqdTij+Xkx1ztGDzOMACTUcjF6iEkiH0=",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """
      Given agent
      """
      {
        "address": "lxDqH19hS9D1Nom8zAvt3vUFS4136FKa6bluwHhbetE=",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
      }
      """
      When query agents
      """
      {}
      """
      Then expect response
      """
      {
        "agents": [
          {
            "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4",
            "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
          },
          {
            "address": "chora1jugw586lv99apafk3x7vczldmm6s2judwl599xhfh9hvq7zm0tgs7jccnn",
            "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
            "metadata": "chora:13toVfwypkE1AwUzQmuBHk28WWwCa5QCynCrBuoYgMvN2iroywJ5Vi1.rdf"
          }
        ],
        "pagination": {
          "total": 2
        }
      }
      """

    # No failing scenario - response is never returned when query fails
