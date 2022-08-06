package main

import (
	"log"

	"github.com/forfd8960/lingo-example/mapreduce/functions"

	"gitlab.com/gitlab-org/vulnerability-research/foss/lingo/eval"
)

func register(fn eval.Function, err error) {
	if err != nil {
		log.Fatalf("failed to create %s function %s:", fn.Symbol(), err.Error())
	}
	err = eval.HookFunction(fn)
	if err != nil {
		log.Fatalf("failed to hook bool function %s:", err.Error())
	}
}

func main() {
	register(functions.NewFunctionList())
	eval.RunLoop()
}
