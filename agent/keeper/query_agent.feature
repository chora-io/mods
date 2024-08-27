Feature: Query/Agent

  Query/Agent is successful when:
  - agent with identifier exists

  Query/Agent has the following outcomes:
  - query response returned

  Rule: The agent must exist

    Background:
      Given agent
      """
      {
        "address": "G+ksLYTNBuzyqdTij+Xkx1ztGDzOMACTUcjF6iEkiH0=",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: agent exists
      When query agent
      """
      {
        "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4"
      }
      """
      Then expect no error

    Scenario: agent does not exist
      When query agent
      """
      {
        "address": "chora1jugw586lv99apafk3x7vczldmm6s2judwl599xhfh9hvq7zm0tgs7jccnn"
      }
      """
      Then expect the error
      """
      agent with address chora1jugw586lv99apafk3x7vczldmm6s2judwl599xhfh9hvq7zm0tgs7jccnn: not found
      """

  Rule: The query response is returned

    Background:
      Given agent
      """
      {
        "address": "G+ksLYTNBuzyqdTij+Xkx1ztGDzOMACTUcjF6iEkiH0=",
        "admin": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: query response returned
      When query agent
      """
      {
        "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4"
      }
      """
      Then expect response
      """
      {
        "address": "chora1r05jctvye5rweu4f6n3gle0ycaww6xpueccqpy63erz75gfy3p7snu2hw4",
        "admin": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - response is never returned when query fails
