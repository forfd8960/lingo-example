package custresults

import (
	"encoding/json"
	"fmt"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/eval"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/types"
)

type SalaryResult struct {
	Val map[string]interface{}
}

func (r SalaryResult) DeepCopy() eval.Result { return NewSalaryResult(r.Val) }
func (r SalaryResult) String() string {
	bs, err := json.Marshal(r.Val)
	if err != nil {
		return ""
	}

	return string(bs)
}
func (r SalaryResult) Type() types.Type   { return types.TypeDictionary }
func (r SalaryResult) Tidy()              {}
func (r SalaryResult) Value() interface{} { return r.Val }
func (r *SalaryResult) SetValue(value interface{}) error {
	val, ok := value.(map[string]interface{})
	if !ok {
		return fmt.Errorf("invalid type for map[string]interface{}")
	}

	r.Val = val
	return nil
}

// NewSalaryResult ...
func NewSalaryResult(value map[string]interface{}) *SalaryResult {
	r := map[string]interface{}{}
	for k, v := range value {
		r[k] = v
	}
	return &SalaryResult{
		Val: r,
	}
}
