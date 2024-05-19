Feature: Customer Getter

  Scenario: Retrieve details of an existing customer
    Given the customer with document "52411797044" exists
    When I retrieve the customer details
    Then I should get the customer details
