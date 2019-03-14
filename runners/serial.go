package runners

import (
	"github.com/testdouble/diplomat/builders"
	"github.com/testdouble/diplomat/differs"
	"github.com/testdouble/diplomat/http"
)

// Serial runs all tests one at a time.
type Serial struct {
	Differ differs.Differ
	Client http.Client
}

// Run returns the results of running all tests in `spec`.
func (s *Serial) Run(test builders.Test) (TestResult, error) {
	response, err := s.Client.Do(test.Request)
	if err != nil {
		return TestResult{}, err
	}

	diff, err := s.Differ.Diff(test.Response, response)
	if err != nil {
		return TestResult{}, err
	}

	return TestResult{
		Name: test.Name,
		Diff: diff,
	}, nil
}

// RunAll returns the results of running all tests in the provided channel.
func (s *Serial) RunAll(tests chan builders.Test, errors chan error) chan TestResult {
	results := make(chan TestResult)

	go func() {
		for test := range tests {
			result, err := s.Run(test)
			if err != nil {
				errors <- err
				return
			}

			results <- result
		}

		close(results)
	}()

	return results
}
