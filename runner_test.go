package main

import (
	"testing"

	"github.com/testdouble/http-assertion-tool/loaders"
	"github.com/testdouble/http-assertion-tool/mocks"
	"github.com/testdouble/http-assertion-tool/parsers"
	"github.com/testdouble/http-assertion-tool/runners"
)

func TestRunnerRun(t *testing.T) {
	loader := &mocks.FileLoader{}
	parser := &mocks.SpecParser{}
	runner := &mocks.SpecRunner{}
	printer := &mocks.ResultsPrinter{}

	file := new(loaders.File)
	spec := new(parsers.Spec)
	result := new(runners.Result)

	loader.On("Load", "test-file").Return(file, nil)
	parser.On("Parse", file).Return(spec, nil)
	runner.On("Run", spec).Return(result, nil)
	printer.On("Print", result).Return(nil)

	subject := Runner{
		loader:  loader,
		parser:  parser,
		runner:  runner,
		printer: printer,
	}

	err := subject.Run("test-file")
	if err != nil {
		t.Fatalf("Failed with: %v", err)
	}

	loader.AssertExpectations(t)
	parser.AssertExpectations(t)
	runner.AssertExpectations(t)
	printer.AssertExpectations(t)
}
