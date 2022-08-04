package functions

import (
	"fmt"

	"github.com/forfd8960/lingo-example/getdata/custresults"
	"github.com/forfd8960/lingo-example/getdata/salary"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/eval"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/parser"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/types"
)

type FunctionGetSalary struct{}

func (f *FunctionGetSalary) Desc() (string, string) {
	return fmt.Sprintf("%s%s %s%s",
			string(parser.TokLeftPar),
			f.Symbol(),
			"min max",
			string(parser.TokRightPar)),
		"get city salary data"
}

func (f *FunctionGetSalary) Symbol() parser.TokLabel {
	return parser.TokLabel("getsalary")
}

func (f *FunctionGetSalary) Validate(env *eval.Environment, stack *eval.StackFrame) error {
	if stack.Size() != 1 {
		return eval.WrongNumberOfArgs(f.Symbol(), stack.Size(), 1)
	}

	arg := stack.Items()[0]

	if arg.Type() != types.TypeString {
		return eval.WrongTypeOfArg(f.Symbol(), 1, arg)
	}

	if len(arg.Value().(string)) == 0 {
		return fmt.Errorf(eval.ErrorMessage(f.Symbol(), "city name is empty"))
	}
	return nil
}

func (f *FunctionGetSalary) Evaluate(env *eval.Environment, stack *eval.StackFrame) (eval.Result, error) {
	var result = map[string]interface{}{}

	var cityName string
	for !stack.Empty() {
		cityName = stack.Pop().(*eval.StringResult).Val
	}

	if len(cityName) == 0 {
		return nil, fmt.Errorf("cityName empty")
	}

	fmt.Println("cityName: ", cityName)

	var err error
	result, err = salary.GetCitySalary(cityName)
	if err != nil {
		return nil, err
	}

	fmt.Println("result: ", cityName)

	return custresults.NewSalaryResult(result), nil
}

func NewFunctionGetSalary() (eval.Function, error) {
	fun := &FunctionGetSalary{}
	_ = parser.HookToken(fun.Symbol())
	return fun, nil
}
