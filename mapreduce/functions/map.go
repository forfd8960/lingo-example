package functions

import (
	"fmt"

	"github.com/forfd8960/lingo-example/mapreduce/custresults"
	"github.com/forfd8960/lingo-example/mapreduce/custtypes"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/eval"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/parser"
)

type FunctionMap struct{}

func (f *FunctionMap) Desc() (string, string) {
	return fmt.Sprintf("%s%s %s%s",
			string(parser.TokLeftPar),
			f.Symbol(),
			"#'f (1 2 3)",
			string(parser.TokRightPar)),
		"map apply f on each num in list"
}

func (f *FunctionMap) Symbol() parser.TokLabel {
	return parser.TokLabel("map")
}

func (f *FunctionMap) Validate(env *eval.Environment, stack *eval.StackFrame) error {
	if stack.Size() < 2 {
		return eval.WrongNumberOfArgs(f.Symbol(), stack.Size(), 1)
	}

	mapfunI := stack.Items()[0]
	if mapfunI.Type() != custtypes.TypeMapf {
		return eval.WrongTypeOfArg(f.Symbol(), 1, mapfunI)
	}

	listDataI := stack.Items()[1]
	if listDataI.Type() != custtypes.TypeListData {
		return eval.WrongTypeOfArg(f.Symbol(), 2, listDataI)
	}

	return nil
}

func (f *FunctionMap) Evaluate(env *eval.Environment, stack *eval.StackFrame) (eval.Result, error) {
	var result = []int64{}
	return custresults.NewListResult(result), nil
}

func NewFunctionMap() (eval.Function, error) {
	fun := &FunctionMap{}
	_ = parser.HookToken(fun.Symbol())
	return fun, nil
}
