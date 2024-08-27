Feature: Query/Governor

  Query/Governor is successful when:
  - governor with address exists

  Query/Governor has the following outcomes:
  - query response returned

  Rule: The governor must exist

    Background:
      Given governor
      """
      {
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: governor exists
      When query governor
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect no error

    Scenario: governor does not exist
      When query governor
      """
      {
        "address": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect the error
      """
      governor with address chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: not found
      """

  Rule: The query response is returned

    Background:
      Given governor
      """
      {
        "address": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: query response returned
      When query governor
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38"
      }
      """
      Then expect response
      """
      {
        "address": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - response is never returned when query fails
