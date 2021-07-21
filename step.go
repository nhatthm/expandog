package expandog

import (
	"github.com/cucumber/godog"
	"github.com/cucumber/messages-go/v10"
)

// StepExpander expands variables in cucumber steps.
type StepExpander struct {
	expanders []interface{}
	expand    Expander
}

// RegisterContext register StepExpander to the test suite.
func (m *StepExpander) RegisterContext(s *godog.ScenarioContext) {
	s.BeforeScenario(func(sc *godog.Scenario) {
		m.expand = chainExpanders(m.expanders...)
	})

	s.BeforeStep(func(st *godog.Step) {
		expandStep(st, m.expand)
	})
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
	st.Text = expand(st.Text)

	if st.Argument == nil {
		return
	}

	switch msg := st.Argument.Message.(type) {
	case *messages.PickleStepArgument_DocString:
		msg.DocString.Content = expand(msg.DocString.Content)

	case *messages.PickleStepArgument_DataTable:
		for _, row := range msg.DataTable.Rows {
			for _, cell := range row.Cells {
				cell.Value = expand(cell.Value)
			}
		}
	}
}
