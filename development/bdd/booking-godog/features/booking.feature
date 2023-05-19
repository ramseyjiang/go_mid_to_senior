Feature: User books a hotel room

  Scenario: User successfully books a hotel room
    Given the user has selected a valid city and date range
    When they select a specific hotel room
    And they input valid personal and payment information
    Then their booking is confirmed
    And they receive a booking confirmation email
