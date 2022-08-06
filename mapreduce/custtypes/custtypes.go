package custtypes

import (
	"fmt"
	"strings"
	"unicode"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/parser"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/types"
)

var TypeListDataId, TypeListData = types.NewTypeWithProperties("ListData", types.Collection)
var TokListData = parser.HookToken(parser.TokLabel(TypeListData.Name))

type ListDataMatcher struct{}

func (i ListDataMatcher) Match(s string) parser.TokLabel {
	s = strings.TrimSpace(s)

	sl := strings.Split(s, " ")
	if len(sl) < 2 {
		return parser.TokUnknown
	}

	if sl[0] != "(" {
		return parser.TokUnknown
	}

	for _, n := range sl[1 : len(sl)-1] {
		n = strings.TrimSpace(n)
		rs := []rune(n)
		if !unicode.IsDigit(rs[0]) {
			return parser.TokUnknown
		}
	}

	if sl[len(sl)-1] != ")" {
		return parser.TokUnknown
	}

	return TokListData.Label
}
func (i ListDataMatcher) Id() string {
	return fmt.Sprintf("%d", TypeListDataId)
}

func init() {
	parser.HookMatcher(ListDataMatcher{})
}
