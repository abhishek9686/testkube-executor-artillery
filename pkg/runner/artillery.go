package runner

import (
	"fmt"
	"path/filepath"

	"github.com/kelseyhightower/envconfig"
	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/executor"
	"github.com/kubeshop/testkube/pkg/executor/content"
	"github.com/kubeshop/testkube/pkg/executor/output"
)

type Params struct {
	GitUsername string // RUNNER_GITUSERNAME
	GitToken    string // RUNNER_GITTOKEN
}

// NewRunner ...
func NewArtilleryRunner() *ArtilleryRunner {
	var params Params
	err := envconfig.Process("runner", &params)
	if err != nil {
		panic(err.Error())
	}
	return &ArtilleryRunner{
		Fetcher: content.NewFetcher(""),
		Params:  params,
	}
}

// ArtilleryRunner ...
type ArtilleryRunner struct {
	Params  Params
	Fetcher content.ContentFetcher
}

func (r *ArtilleryRunner) Run(execution testkube.Execution) (result testkube.ExecutionResult, err error) {
	// make some validation
	err = r.Validate(execution)
	if err != nil {
		return result, err
	}
	if r.Params.GitUsername != "" && r.Params.GitToken != "" {
		if execution.Content != nil && execution.Content.Repository != nil {
			execution.Content.Repository.Username = r.Params.GitUsername
			execution.Content.Repository.Token = r.Params.GitToken
		}
	}
	path, err := r.Fetcher.Fetch(execution.Content)
	if err != nil {
		return result, err
	}

	output.PrintEvent("created content path", path)

	// if execution.Content.IsFile() {
	// 	output.PrintEvent("using file", execution)
	// 	// TODO implement file based test content for string, git-file, file-uri
	// 	//      or remove if not used
	// }

	// if execution.Content.IsDir() {
	// 	output.PrintEvent("using dir", execution)
	// 	// TODO implement file based test content for git-dir
	// 	//      or remove if not used
	// }

	params := make([]string, 0, len(execution.Params))
	for key, value := range execution.Params {
		params = append(params, fmt.Sprintf("%s=%s", key, value))
	}
	testDir, _ := filepath.Split(execution.Content.Repository.Path)
	args := []string{"run", path}
	if len(params) != 0 {
		args = append(args, params...)
	}
	testReportFile := filepath.Join(testDir, "test-report.json")
	args = append(args, "-o", testReportFile)
	// append args from execution
	args = append(args, execution.Args...)

	// run executor here
	out, err := executor.Run(path, "artillery", args...)
	// error result should be returned if something is not ok
	if err != nil {
		return result.Err(fmt.Errorf("some test execution related error occured")), err
	}
	artilleryResult, err := r.GetArtilleryExecutionResult(testReportFile, out)

	if err != nil {
		return result.Err(fmt.Errorf("failed to get test execution results")), err
	}

	// return ExecutionResult
	return MapTestSummaryToResults(artilleryResult), nil
}
