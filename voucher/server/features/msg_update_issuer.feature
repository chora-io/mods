Feature: Msg/UpdateIssuer

  UpdateIssuer is successful when:
  - issuer is the voucher issuer

  UpdateIssuer has the following outcomes:
  - message response returned
  - Content is updated in state
  - EventUpdateIssuer is emitted

  Rule: The issuer must be the voucher issuer

    Background:
      Given voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: issuer is voucher issuer
      When msg update issuer
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_issuer": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect no error

    Scenario: issuer is not voucher issuer
      When msg update issuer
      """
      {
        "id": 1,
        "issuer": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup",
        "new_issuer": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect the error
      """
      issuer chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup: voucher issuer chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38: unauthorized
      """

  Rule: The message response is returned

    Background:
      Given voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: message response returned
      When msg update issuer
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_issuer": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect response
      """
      {
        "id": 1
      }
      """

    # No failing scenario - response is never returned when message fails

  Rule: Content is updated in state

    Background:
      Given voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: state voucher updated
      When msg update issuer
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_issuer": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect state voucher
      """
      {
        "id": 1,
        "issuer": "hEyiXxUCaFQmkbuhO9r+QDscjIY=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    # No failing scenario - state is never updated when message fails

  Rule: EventUpdateIssuer emitted

    Background:
      Given voucher
      """
      {
        "id": 1,
        "issuer": "BTZfSbi0JKqguZ/tIAPUIhdAa7Y=",
        "metadata": "chora:13toVfvC2YxrrfSXWB5h2BGHiXZURsKxWUz72uDRDSPMCrYPguGUXSC.rdf"
      }
      """

    Scenario: event update issuer emitted
      When msg update issuer
      """
      {
        "id": 1,
        "issuer": "chora1q5m97jdcksj24g9enlkjqq75ygt5q6ak54jk38",
        "new_issuer": "chora1s3x2yhc4qf59gf53hwsnhkh7gqa3eryxnu6nup"
      }
      """
      Then expect event update issuer
      """
      {
        "id": 1
      }
      """

    # No failing scenario - event is never emitted when message fails
