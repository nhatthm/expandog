package expandog

import (
	"os"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPlaceholder(t *testing.T) {
	t.Parallel()

	actual := placeholder("VAR")
	expected := "$VAR"

	assert.Equal(t, expected, actual)
}

func TestRuntimeExpander(t *testing.T) {
	t.Parallel()

	expand := runtimeExpander(func() Pairs {
		return Pairs{
			"HUSBAND": "John",
			"WIFE":    "Jane",
		}
	})

	content := "$HUSBAND & $WIFE"
	actual := expand(content)
	expected := "John & Jane"

	assert.Equal(t, expected, actual)
}

func TestMapExpander(t *testing.T) {
	t.Parallel()

	expand := mapExpander(Pairs{
		"HUSBAND": "John",
		"WIFE":    "Jane",
	})

	content := "$HUSBAND & $WIFE"
	actual := expand(content)
	expected := "John & Jane"

	assert.Equal(t, expected, actual)
}

func TestChainExpanders(t *testing.T) {
	t.Parallel()

	expand := chainExpanders(
		// Raw replacer.
		strings.NewReplacer("$TO", "Berlin"),
		// Our expanders.
		Pairs{
			"HUSBAND": "John",
		},
		func() Pairs {
			return Pairs{
				"WIFE": "Jane",
			}
		},
		func() Expander {
			return func(s string) string {
				return strings.ReplaceAll(s, "$DURATION", "and stay there for 3 days")
			}
		},
		func(s string) string {
			return strings.ReplaceAll(s, "$FROM", "Paris")
		},
		Expander(func(s string) string {
			return strings.ReplaceAll(s, "$TRANSPORT", "by bus")
		}),
		// Os.
		EnvExpander,
	)

	// Set os env.
	assert.NoError(t, os.Setenv("DATE", "Thursday"))

	defer func() {
		_ = os.Unsetenv("DATE") // nolint:errcheck
	}()

	content := "On $DATE, $HUSBAND & $WIFE are going from $FROM to $TO $TRANSPORT $DURATION"
	actual := expand(content)
	expected := "On Thursday, John & Jane are going from Paris to Berlin by bus and stay there for 3 days"

	assert.Equal(t, expected, actual)
}

func TestChainExpanders_UnsupportedType(t *testing.T) {
	t.Parallel()

	assert.PanicsWithError(t, "unsupported expander: got int", func() {
		chainExpanders(Pairs{
			"HUSBAND": "John",
		}, 42)
	})
}
