Feature:
    Expand variables

    Scenario: Env var is replaced
        Given env var FOOBAR is replaced in step text: $FOOBAR

        Then step text is:
        """
        env var FOOBAR is replaced in step text: foobar
        """

        Given env var FOOBAR is replaced in step argument (string)
        """
        FOOBAR=$FOOBAR
        """

        Then step argument is a string:
        """
        FOOBAR=foobar
        """

        Given env var FOOBAR is replaced in step argument (table)
            | col 1   | col 2   | col 3   |
            | value 1 | $FOOBAR | value 3 |

        Then step argument is a table:
            | col 1   | col 2  | col 3   |
            | value 1 | foobar | value 3 |

    Scenario: Map var is replaced
        Given map var NAME is replaced in step text: $NAME

        Then step text is:
        """
        map var NAME is replaced in step text: John
        """

        Given map var NAME is replaced in step argument (string)
        """
        NAME=$NAME
        """

        Then step argument is a string:
        """
        NAME=John
        """

        Given map var NAME is replaced in step argument (table)
            | col 1   | col 2 | col 3   |
            | value 1 | $NAME | value 3 |

        Then step argument is a table:
            | col 1   | col 2 | col 3   |
            | value 1 | John  | value 3 |

    Scenario: Scenario Provider runs only once every scenario
        Given current timestamp is "$TIMESTAMP"

        Then First call, timestamp = "$TIMESTAMP"
        Then Second call, timestamp = "$TIMESTAMP"

    Scenario: Scenario Provider runs once again in another scenario
        Given current timestamp is "$TIMESTAMP"

        Then First call, timestamp = "$TIMESTAMP"
        Then Second call, timestamp = "$TIMESTAMP"
