package parsers

import (
	"strings"

	"github.com/testdouble/http-assertion-tool/loaders"
)

// The Markdown parser parses all lines inside of code fences (```).
type Markdown struct {
	plainText PlainTextParser
}

const (
	inRichText = iota
	inCodeFence
)

// Parse parses all lines in `body`.
func (m *Markdown) Parse(body *loaders.Body) (*Spec, error) {
	mode := inRichText
	thisTestName := ""
	nextTestName := ""
	state := newParserState()
	state.finalizer = func(test *Test) {
		if len(thisTestName) == 0 {
			fallbackTestName(test)
		} else {
			test.Name = thisTestName
		}

		thisTestName = nextTestName
		nextTestName = ""
	}

	for _, line := range body.Lines {
		trimmedLine := strings.TrimSpace(line)

		switch {
		case strings.HasPrefix(trimmedLine, "```"):
			if mode == inRichText {
				mode = inCodeFence
			} else {
				mode = inRichText
			}
		case mode == inCodeFence:
			state.addLine(line)
		case strings.HasPrefix(trimmedLine, "#"):
			name := strings.TrimSpace(strings.SplitN(trimmedLine, " ", 2)[1])

			if state.mode == modeAwaitingRequest {
				thisTestName = name
			} else {
				nextTestName = name
			}
		}
	}

	return state.finalize()
}
