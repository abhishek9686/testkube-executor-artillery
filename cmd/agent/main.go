package main

import (
	"os"

	"github.com/abhishek9686/testkube-executor-artillery/pkg/runner"
	"github.com/kubeshop/testkube/pkg/executor/agent"
)

func main() {
	agent.Run(runner.NewArtilleryRunner(), os.Args)
}
