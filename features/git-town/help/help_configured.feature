Feature: show help screen when Git Town is configured

  As a user having forgotten how to use Git Town
  I want to see a helpful list of all commands
  So that I can refresh my memory quickly and move on to what I actually wanted to do.


  Background:
    Given the main branch name is "main"
    And my perennial branches are configured as "qa" and "staging"


  Scenario Outline:
    When I run `<COMMAND>`
    Then Git Town prints
      """
      Usage:
        git-town [command]
      """

    Examples:
      | COMMAND       |
      | git-town      |
      | git-town help |
