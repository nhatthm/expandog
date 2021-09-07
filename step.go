package expandog

import (
	"github.com/cucumber/godog"
	"github.com/godogx/expandvars"
)

// StepExpander expands variables in cucumber steps.
//
// Deprecated: Use expandvars.StepExpander instead.
type StepExpander = expandvars.StepExpander

// NewStepExpander initiates a new variable expanders for cucumber steps.
//
// Deprecated: Use expandvars.NewStepExpander instead.
func NewStepExpander(expanders ...interface{}) *StepExpander {
	return expandvars.NewStepExpander(expanders...)
}

// ExpandStep expands variables in the step definition using an expander.
//
// Deprecated: Use expandvars.ExpandStep instead.
func ExpandStep(st *godog.Step, expanders ...interface{}) {
	expandvars.ExpandStep(st, expanders...)
}
