package runner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

type ArtilleryExecutionResult struct {
	Output  string
	Summary ArtilleryTestSummary
}

type ArtilleryTestSummary struct {
	Aggregate struct {
		Counters struct {
			CoreVusersCreatedTotal     int `json:"core.vusers.created.total"`
			EngineHTTPRequests         int `json:"engine.http.requests"`
			EngineHTTPCodes200         int `json:"engine.http.codes.200"`
			EngineHTTPResponses        int `json:"engine.http.responses"`
			PluginsExpectOk            int `json:"plugins.expect.ok"`
			PluginsExpectOkStatusCode  int `json:"plugins.expect.ok.statusCode"`
			PluginsExpectOkContentType int `json:"plugins.expect.ok.contentType"`
			CoreVusersCompleted        int `json:"core.vusers.completed"`
			ErrorsECONNREFUSED         int `json:"errors.ECONNREFUSED"`
			CoreVusersFailed           int `json:"core.vusers.failed"`
		} `json:"counters"`
	} `json:"aggregate"`
}

// Validate checks if Execution has valid data in context of Cypress executor
// Cypress executor runs currently only based on cypress project
func (r *ArtilleryRunner) Validate(execution testkube.Execution) error {

	if execution.Content == nil {
		return fmt.Errorf("can't find any content to run in execution data: %+v", execution)
	}

	if execution.Content.Repository == nil {
		return fmt.Errorf("cypress executor handle only repository based tests, but repository is nil")
	}

	if execution.Content.Repository.Path == "" {
		return fmt.Errorf("can't find repository path in params, repo:%+v", execution.Content.Repository)
	}

	if execution.Content.Repository.Branch == "" {
		return fmt.Errorf("can't find branch in params, repo:%+v", execution.Content.Repository)
	}

	return nil
}

func (r *ArtilleryRunner) GetArtilleryExecutionResult(testReportFile string, out []byte) (ArtilleryExecutionResult, error) {
	result := ArtilleryExecutionResult{}
	result.Output = string(out)
	data, err := ioutil.ReadFile(testReportFile)
	if err != nil {
		return result, err
	}
	err = json.Unmarshal(data, &result.Summary)
	if err != nil {
		return result, err
	}
	return result, nil
}

func MapTestSummaryToResults(artilleryResult ArtilleryExecutionResult) testkube.ExecutionResult {

	status := testkube.StatusPtr(testkube.PASSED_ExecutionStatus)
	if artilleryResult.Summary.Aggregate.Counters.CoreVusersFailed > 0 {
		status = testkube.StatusPtr(testkube.FAILED_ExecutionStatus)

	}
	result := testkube.ExecutionResult{
		Output:     artilleryResult.Output,
		OutputType: "text/plain",
		Status:     status,
	}
	return result

}
