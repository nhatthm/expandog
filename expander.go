package expandog

import (
	"errors"
	"fmt"
	"os"
	"strings"
)

// ErrUnsupportedExpander indicates that the provided expander is not supported.
var ErrUnsupportedExpander = errors.New("unsupported expander")

// EnvExpander expands variables using env vars.
var EnvExpander Expander = os.ExpandEnv

// Pairs is a pair of old and new to be replaced.
type Pairs = map[string]string

// Expander expands the variables in a string.
type Expander func(string) string

// Replacer replace string.
type Replacer interface {
	Replace(string) string
}

// BeforeScenario expands variables from a provider that provide only before running a scenario.
func BeforeScenario(provide func() Pairs) Expander {
	return mapExpander(provide())
}

func placeholder(name string) string {
	return fmt.Sprintf("$%s", name)
}

// runtimeExpander expands variables from a provider.
func runtimeExpander(provide func() Pairs) Expander {
	return func(s string) string {
		return mapExpander(provide())(s)
	}
}

// mapExpander initiates a new variable expander from a map of values.
func mapExpander(pairs Pairs) Expander {
	oldNew := make([]string, 0, 2*len(pairs))

	for k, v := range pairs {
		oldNew = append(oldNew, placeholder(k), v)
	}

	return strings.NewReplacer(oldNew...).Replace
}

func chainExpanders(expanders ...interface{}) Expander {
	l := make([]Expander, 0, len(expanders))

	for _, e := range expanders {
		l = append(l, newExpander(e))
	}

	return func(s string) string {
		for _, expand := range l {
			s = expand(s)
		}

		return s
	}
}

func newExpander(e interface{}) Expander {
	switch e := e.(type) {
	case Pairs:
		return mapExpander(e)

	case func() map[string]string:
		return runtimeExpander(e)

	case Replacer:
		return e.Replace

	case Expander:
		return e

	case func(string) string:
		return e
	}

	panic(fmt.Errorf("%w: got %T", ErrUnsupportedExpander, e))
}
