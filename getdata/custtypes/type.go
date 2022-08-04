package custtypes

import (
	"strconv"
	"strings"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/parser"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/types"
)

var TypeSalaryId, TypeSalary = types.NewTypeWithProperties("salary", types.Collection)
var TokFloat = parser.HookToken(parser.TokLabel(TypeSalary.Name))

// recognize (true) als boolean
type FloatMatcher struct{}

func (i FloatMatcher) Match(s string) parser.TokLabel {
	if !strings.Contains(s, ".") {
		return parser.TokUnknown
	}

	if _, err := strconv.ParseFloat(s, 32); err == nil {
		return TokFloat.Label
	}
	return parser.TokUnknown
}
func (i FloatMatcher) Id() string {
	return string(TokFloat.Label)
}

func init() {
	parser.HookMatcher(FloatMatcher{})
}
