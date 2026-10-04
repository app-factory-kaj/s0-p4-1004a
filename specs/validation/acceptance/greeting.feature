Feature: Greeting

  @story-1
  Rule: A caller supplying a name receives a greeting addressed to that name

    Scenario: Greeting a named caller
      Given the greeter service is available
      When a Caller sends a name of "Ada" to the greeter
      Then the Caller receives a JSON greeting addressed to "Ada"

  @story-2
  Rule: A caller who omits the name still receives a predictable greeting

    Scenario: Greeting a caller with no name supplied
      Given the greeter service is available
      When a Caller requests a greeting without supplying a name
      Then the Caller receives a JSON greeting that does not fail
