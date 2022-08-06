package custresults

import (
	"fmt"
	"strings"

	"github.com/forfd8960/lingo-example/mapreduce/custtypes"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/eval"
	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/types"
)

type ListResult struct {
	Vals []int64
}

func (r ListResult) DeepCopy() eval.Result { return NewListResult(r.Vals) }
func (r ListResult) String() string {
	var sb = &strings.Builder{}

	fmt.Println("r.Vals: ", r.Vals)

	sb.WriteString("(")
	for _, v := range r.Vals[0 : len(r.Vals)-1] {
		sb.WriteString(fmt.Sprintf("%d ", v))
	}

	sb.WriteString(fmt.Sprintf("%d)", r.Vals[len(r.Vals)-1]))

	return sb.String()
}
func (r ListResult) Type() types.Type   { return custtypes.TypeListData }
func (r ListResult) Tidy()              {}
func (r ListResult) Value() interface{} { return r.Vals }
func (r *ListResult) SetValue(value interface{}) error {
	vals, ok := value.([]int64)
	if !ok {
		return fmt.Errorf("invalid type for []int64")
	}

	r.Vals = vals
	return nil
}

// NewListResult ...
func NewListResult(vals []int64) *ListResult {
	r := make([]int64, len(vals))
	copy(r, vals)

	return &ListResult{
		Vals: r,
	}
}
