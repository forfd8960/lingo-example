package functions

import (
	"fmt"

	"github.com/forfd8960/lingo-example/mapreduce/custresults"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/eval"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/parser"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/types"
)

type FunctionList struct{}

func (f *FunctionList) Desc() (string, string) {
	return fmt.Sprintf("%s%s %s%s",
			string(parser.TokLeftPar),
			f.Symbol(),
			"1 2 3 4 5 ...",
			string(parser.TokRightPar)),
		"get list of values"
}

func (f *FunctionList) Symbol() parser.TokLabel {
	return parser.TokLabel("list")
}

func (f *FunctionList) Validate(env *eval.Environment, stack *eval.StackFrame) error {
	if stack.Size() < 1 {
		return eval.WrongNumberOfArgs(f.Symbol(), stack.Size(), 1)
	}

	for idx, item := range stack.Items() {
		if item.Type() != types.TypeInt {
			return eval.WrongTypeOfArg(f.Symbol(), idx+1, item)
		}
	}
	return nil
}

func (f *FunctionList) Evaluate(env *eval.Environment, stack *eval.StackFrame) (eval.Result, error) {
	var result = []int64{}

	for !stack.Empty() {
		item := stack.Pop().(*eval.IntResult).Val
		result = append(result, int64(item))
	}

	var rs = []int64{}
	for i := len(result) - 1; i >= 0; i-- {
		rs = append(rs, result[i])
	}

	if len(rs) == 0 {
		return nil, fmt.Errorf("argument empty")
	}

	fmt.Println("result: ", rs)

	return custresults.NewListResult(rs), nil
}

func NewFunctionList() (eval.Function, error) {
	fun := &FunctionList{}
	_ = parser.HookToken(fun.Symbol())
	return fun, nil
}
