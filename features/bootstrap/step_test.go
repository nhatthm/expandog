package bootstrap

import (
	"context"
	"testing"

	"github.com/cucumber/godog"
	"github.com/stretchr/testify/assert"
)

type stepCtx struct {
	test         *testing.T
	stepText     string
	stepArgument interface{}
	timestamp    string
}

var doNothing = func() error { return nil }

func (s *stepCtx) RegisterContext(t *testing.T, sc *godog.ScenarioContext) {
	t.Helper()

	s.test = t

	sc.StepContext().After(func(_ context.Context, st *godog.Step, _ godog.StepResultStatus, _ error) (context.Context, error) {
		s.stepText = st.Text

		if st.Argument == nil {
			return nil, nil
		}

		if st.Argument.DocString != nil {
			s.stepArgument = st.Argument.DocString.Content
		}

		if st.Argument.DataTable != nil {
			s.stepArgument = st.Argument.DataTable
		}

		return nil, nil
	})

	sc.Step("step text is:", s.assertStepText)
	sc.Step("step argument is a string:", s.assertStepArgumentString)
	sc.Step("step argument is a table:", s.assertStepArgumentTable)
	sc.Step(`current timestamp is "([^"]*)"`, s.setCurrentTimestamp)
	sc.Step(`timestamp = "([^"]*)"`, s.assertTimestamp)
	sc.Step(".*", doNothing)
}

func (s *stepCtx) assertStepText(text *godog.DocString) error {
	assert.Equal(s.test, text.Content, s.stepText)

	return nil
}

func (s *stepCtx) assertStepArgumentString(text *godog.DocString) error {
	assert.Equal(s.test, text.Content, s.stepArgument)

	return nil
}

func (s *stepCtx) assertStepArgumentTable(text *godog.Table) error {
	assert.Equal(s.test, text, s.stepArgument)

	return nil
}

func (s *stepCtx) setCurrentTimestamp(timestamp string) error {
	assert.NotEqual(s.test, "$TIMESTAMP", timestamp)
	assert.NotEqual(s.test, s.timestamp, timestamp)

	s.timestamp = timestamp

	return nil
}

func (s *stepCtx) assertTimestamp(timestamp string) error {
	assert.NotEqual(s.test, "$TIMESTAMP", timestamp)
	assert.Equal(s.test, s.timestamp, timestamp)

	return nil
}

func newStepCtx() *stepCtx {
	return &stepCtx{}
}
