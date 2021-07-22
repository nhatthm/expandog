package expandog_test

import (
	"os"
	"strings"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"

	"github.com/nhatthm/expandog"
)

func TestExpandStep(t *testing.T) {
	t.Parallel()

	expanders := []interface{}{
		// Raw replacer.
		strings.NewReplacer("$TO", "Berlin"),
		// Our expanders.
		expandog.Pairs{
			"HUSBAND": "John",
		},
		func() expandog.Pairs {
			return expandog.Pairs{
				"WIFE": "Jane",
			}
		},
		func() expandog.Expander {
			return func(s string) string {
				return strings.ReplaceAll(s, "$DURATION", "and stay there for 3 days")
			}
		},
		func(s string) string {
			return strings.ReplaceAll(s, "$FROM", "Paris")
		},
		expandog.Expander(func(s string) string {
			return strings.ReplaceAll(s, "$TRANSPORT", "by bus")
		}),
		// Os.
		expandog.EnvExpander,
	}

	// Set os env.
	assert.NoError(t, os.Setenv("GREETINGS", "Hi Dave"))

	defer func() {
		_ = os.Unsetenv("GREETINGS") // nolint:errcheck
	}()

	step := &godog.Step{Text: "$GREETINGS, $HUSBAND & $WIFE are going from $FROM to $TO $TRANSPORT $DURATION"}
	expected := "Hi Dave, John & Jane are going from Paris to Berlin by bus and stay there for 3 days"

	expandog.ExpandStep(step, expanders...)

	assert.Equal(t, expected, step.Text)
}
