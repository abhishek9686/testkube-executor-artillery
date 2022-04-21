package runner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
)

type ArtilleryExecutionResult struct {
	Output string
	Result ArtilleryTestResult
}
type Mapping struct {
	RelativeAccuracy float64 `json:"relativeAccuracy"`
	Offset           float64 `json:"_offset"`
	Gamma            float64 `json:"gamma"`
	Multiplier       float64 `json:"_multiplier"`
	MinPossible      float64 `json:"minPossible"`
	MaxPossible      float64 `json:"maxPossible"`
}

type Store struct {
	ChunkSize float64   `json:"chunkSize"`
	Bins      []float64 `json:"bins"`
	Count     float64   `json:"count"`
	MinKey    float64   `json:"minKey"`
	MaxKey    float64   `json:"maxKey"`
	Offset    float64   `json:"offset"`
}

type HistogramMetrics struct {
	Mapping       Mapping `json:"mapping"`
	Store         Store   `json:"store"`
	NegativeStore Store   `json:"negativeStore"`
	ZeroCount     int     `json:"zeroCount"`
	Count         int     `json:"count"`
	Min           float64 `json:"min"`
	Max           float64 `json:"max"`
	Sum           float64 `json:"sum"`
}

type Summary struct {
	Min    float64 `json:"min"`
	Max    float64 `json:"max"`
	Count  float64 `json:"count"`
	P50    float64 `json:"p50"`
	Median float64 `json:"median"`
	P75    float64 `json:"p75"`
	P90    float64 `json:"p90"`
	P95    float64 `json:"p95"`
	P99    float64 `json:"p99"`
	P999   float64 `json:"p999"`
}
type ArtilleryTestResult struct {
	Aggregate struct {
		Counters struct {
			VusersCreatedByNameGetVotes                 int `json:"vusers.created_by_name.Get Votes"`
			VusersCreated                               int `json:"vusers.created"`
			HTTPRequests                                int `json:"http.requests"`
			HTTPCodes200                                int `json:"http.codes.200"`
			HTTPResponses                               int `json:"http.responses"`
			PluginsExpectOk                             int `json:"plugins.expect.ok"`
			PluginsExpectOkStatusCode                   int `json:"plugins.expect.ok.statusCode"`
			PluginsExpectOkContentType                  int `json:"plugins.expect.ok.contentType"`
			PluginsMetricsByEndpointGetRequestCodes200  int `json:"plugins.metrics-by-endpoint.getRequest.codes.200"`
			VusersFailed                                int `json:"vusers.failed"`
			VusersCompleted                             int `json:"vusers.completed"`
			VusersCreatedByNamePostVotes                int `json:"vusers.created_by_name.Post Votes"`
			PluginsMetricsByEndpointPostRequestCodes200 int `json:"plugins.metrics-by-endpoint.postRequest.codes.200"`
		} `json:"counters"`
		Histograms struct {
			HTTPResponseTime                                HistogramMetrics `json:"http.response_time"`
			PluginsMetricsByEndpointResponseTimeGetRequest  HistogramMetrics `json:"plugins.metrics-by-endpoint.response_time.getRequest"`
			VusersSessionLength                             HistogramMetrics `json:"vusers.session_length"`
			PluginsMetricsByEndpointResponseTimePostRequest HistogramMetrics `json:"plugins.metrics-by-endpoint.response_time.postRequest"`
		} `json:"histograms"`
		Rates struct {
			HTTPRequestRate float64 `json:"http.request_rate"`
		} `json:"rates"`
		HTTPRequestRate  float64 `json:"http.request_rate"`
		FirstCounterAt   float64 `json:"firstCounterAt"`
		FirstHistogramAt float64 `json:"firstHistogramAt"`
		LastCounterAt    float64 `json:"lastCounterAt"`
		LastHistogramAt  float64 `json:"lastHistogramAt"`
		FirstMetricAt    float64 `json:"firstMetricAt"`
		LastMetricAt     float64 `json:"lastMetricAt"`
		Period           float64 `json:"period"`
		Summaries        struct {
			HTTPResponseTime                                Summary `json:"http.response_time"`
			PluginsMetricsByEndpointResponseTimeGetRequest  Summary `json:"plugins.metrics-by-endpoint.response_time.getRequest"`
			VusersSessionLength                             Summary `json:"vusers.session_length"`
			PluginsMetricsByEndpointResponseTimePostRequest Summary `json:"plugins.metrics-by-endpoint.response_time.postRequest"`
		} `json:"summaries"`
	} `json:"aggregate"`
	Intermediate []struct {
		Counters struct {
			VusersCreatedByNamePostVotes                float64 `json:"vusers.created_by_name.Post Votes"`
			VusersCreated                               float64 `json:"vusers.created"`
			HTTPRequests                                float64 `json:"http.requests"`
			HTTPCodes200                                float64 `json:"http.codes.200"`
			HTTPResponses                               float64 `json:"http.responses"`
			PluginsExpectOk                             float64 `json:"plugins.expect.ok"`
			PluginsExpectOkStatusCode                   float64 `json:"plugins.expect.ok.statusCode"`
			PluginsMetricsByEndpointPostRequestCodes200 float64 `json:"plugins.metrics-by-endpoint.postRequest.codes.200"`
			VusersFailed                                float64 `json:"vusers.failed"`
			VusersCompleted                             float64 `json:"vusers.completed"`
			VusersCreatedByNameGetVotes                 float64 `json:"vusers.created_by_name.Get Votes"`
			PluginsExpectOkContentType                  float64 `json:"plugins.expect.ok.contentType"`
			PluginsMetricsByEndpointGetRequestCodes200  float64 `json:"plugins.metrics-by-endpoint.getRequest.codes.200"`
		} `json:"counters"`
		Histograms struct {
			HTTPResponseTime                                HistogramMetrics `json:"http.response_time"`
			PluginsMetricsByEndpointResponseTimePostRequest HistogramMetrics `json:"plugins.metrics-by-endpoint.response_time.postRequest"`
			VusersSessionLength                             HistogramMetrics `json:"vusers.session_length"`
			PluginsMetricsByEndpointResponseTimeGetRequest  HistogramMetrics `json:"plugins.metrics-by-endpoint.response_time.getRequest"`
		} `json:"histograms"`
		Rates struct {
			HTTPRequestRate float64 `json:"http.request_rate"`
		} `json:"rates"`
		HTTPRequestRate  float64 `json:"http.request_rate"`
		FirstCounterAt   float64 `json:"firstCounterAt"`
		FirstHistogramAt float64 `json:"firstHistogramAt"`
		LastCounterAt    float64 `json:"lastCounterAt"`
		LastHistogramAt  float64 `json:"lastHistogramAt"`
		FirstMetricAt    float64 `json:"firstMetricAt"`
		LastMetricAt     float64 `json:"lastMetricAt"`
		Period           string  `json:"period"`
		Summaries        struct {
			HTTPResponseTime                                Summary `json:"http.response_time"`
			PluginsMetricsByEndpointResponseTimePostRequest Summary `json:"plugins.metrics-by-endpoint.response_time.postRequest"`
			VusersSessionLength                             Summary `json:"vusers.session_length"`
			PluginsMetricsByEndpointResponseTimeGetRequest  Summary `json:"plugins.metrics-by-endpoint.response_time.getRequest"`
		} `json:"summaries"`
	} `json:"intermediate"`
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
	err = json.Unmarshal(data, &result.Result)
	if err != nil {
		return result, err
	}
	return result, nil
}

func MapTestSummaryToResults(artilleryResult ArtilleryExecutionResult) testkube.ExecutionResult {

	status := testkube.StatusPtr(testkube.PASSED_ExecutionStatus)
	if artilleryResult.Result.Aggregate.Counters.VusersFailed > 0 {
		status = testkube.StatusPtr(testkube.FAILED_ExecutionStatus)

	}
	result := testkube.ExecutionResult{
		Output:     artilleryResult.Output,
		OutputType: "text/plain",
		Status:     status,
	}
	return result

}
