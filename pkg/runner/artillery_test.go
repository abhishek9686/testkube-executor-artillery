package runner

import (
	"testing"

	"github.com/kubeshop/testkube/pkg/api/v1/testkube"
	"github.com/kubeshop/testkube/pkg/executor"
	"github.com/stretchr/testify/assert"
)

func TestRun(t *testing.T) {

	t.Run("runner should run test based on execution data", func(t *testing.T) {
		// given
		_, err := executor.Run("", "npm", "install", "-g", "artillery@latest")
		if err != nil {
			t.Errorf("npm install artillery error: %v", err)
		}
		runner := NewArtilleryRunner()
		repoURI := "https://github.com/abhishek9686/testkube-executor-artillery.git"
		result, err := runner.Run(testkube.Execution{
			Content: &testkube.TestContent{
				Type_: string(testkube.TestContentTypeGitFile),
				Repository: &testkube.Repository{
					Type_:  "git",
					Uri:    repoURI,
					Branch: "main",
					Path:   "examples/test.yaml",
				},
			},
		})

		// then
		assert.NoError(t, err)
		assert.Equal(t, *result.Status, testkube.PASSED_ExecutionStatus)

	})

}
