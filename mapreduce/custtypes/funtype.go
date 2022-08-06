package custtypes

import (
	"fmt"
	"strings"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/parser"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/types"
)

var TypeMapfId, TypeMapf = types.NewTypeWithProperties("Mapf", types.Symbol)
var TokMapf = parser.HookToken(parser.TokLabel(TypeMapf.Name))

type MapfMatcher struct{}

func (i MapfMatcher) Match(s string) parser.TokLabel {
	s = strings.TrimSpace(s)

	if len(s) < 3 {
		return parser.TokUnknown
	}

	sl := []rune(s)

	if sl[0] != '#' {
		return parser.TokUnknown
	}

	if sl[1] != '\'' {
		return parser.TokUnknown
	}

	return TokListData.Label
}
func (i MapfMatcher) Id() string {
	return fmt.Sprintf("%d", TypeMapfId)
}

func init() {
	parser.HookMatcher(MapfMatcher{})
}
