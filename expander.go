package expandog

import (
	"github.com/godogx/expandvars"
)

// ErrUnsupportedExpander indicates that the provided expander is not supported.
//
// Deprecated: Use expandvars.ErrUnsupportedExpander instead.
var ErrUnsupportedExpander = expandvars.ErrUnsupportedExpander

// EnvExpander expands variables using env vars.
//
// Deprecated: Use expandvars.EnvExpander instead.
var EnvExpander = expandvars.EnvExpander

// Pairs is a pair of old and new to be replaced.
//
// Deprecated: Use expandvars.Pairs instead.
type Pairs = expandvars.Pairs

// Expander expands the variables in a string.
//
// Deprecated: Use expandvars.Expander instead.
type Expander = expandvars.Expander

// Replacer replace string.
//
// Deprecated: Use expandvars.Replacer instead.
type Replacer = expandvars.Replacer

// BeforeScenario expands variables from a provider that will be called only once before every scenario.
//
// Deprecated: Use expandvars.BeforeScenario instead.
func BeforeScenario(provide func() Pairs) func() Expander {
	return expandvars.BeforeScenario(provide)
}
