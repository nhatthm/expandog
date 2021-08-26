package expandog

import (
	"context"

	"github.com/cucumber/godog"
)

// StepExpander expands variables in cucumber steps.
type StepExpander struct {
	expanders []interface{}
	expand    Expander
}

// RegisterExpander registers only the expander to the test suite to let it work.
// There will be no registration of step definition in this method.
func (m *StepExpander) RegisterExpander(s *godog.ScenarioContext) {
	s.Before(func(ctx context.Context, sc *godog.Scenario) (context.Context, error) {
		m.expand = chainExpanders(m.expanders...)

		return nil, nil
	})

	s.StepContext().Before(func(_ context.Context, st *godog.Step) (context.Context, error) {
		expandStep(st, m.expand)

		return nil, nil
	})
}

// RegisterContext register everything to the test suite.
func (m *StepExpander) RegisterContext(s *godog.ScenarioContext) {
	m.RegisterExpander(s)
}

// NewStepExpander initiates a new variable expanders for cucumber steps.
func NewStepExpander(expanders ...interface{}) *StepExpander {
	return &StepExpander{
		expanders: expanders,
	}
}

// ExpandStep expands variables in the step definition using an expander.
func ExpandStep(st *godog.Step, expanders ...interface{}) {
	expandStep(st, chainExpanders(expanders...))
}

func expandStep(st *godog.Step, expand Expander) {
	st.Text = doExpand(expand, st.Text)

	if st.Argument == nil {
		return
	}

	if st.Argument.DocString != nil {
		st.Argument.DocString.Content = doExpand(expand, st.Argument.DocString.Content)
	}

	if st.Argument.DataTable != nil {
		for _, row := range st.Argument.DataTable.Rows {
			for _, cell := range row.Cells {
				cell.Value = doExpand(expand, cell.Value)
			}
		}
	}
}
