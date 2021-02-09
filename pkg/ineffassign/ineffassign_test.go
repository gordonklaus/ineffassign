package ineffassign

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	analysistest.Run(t, analysistest.TestData(), Analyzer)
}
