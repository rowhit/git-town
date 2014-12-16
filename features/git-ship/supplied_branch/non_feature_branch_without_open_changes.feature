Feature: git ship: don't ship non-feature branches (without open changes)

  Background:
    Given non-feature branch configuration "qa, production"
    And I am on the "feature" branch
    When I run `git ship production -m 'feature done'` while allowing errors


  Scenario: result
    Then I get the error "The branch 'production' is not a feature branch. Only feature branches can be shipped."
    And I am still on the "feature" branch
