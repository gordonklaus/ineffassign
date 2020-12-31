package main

import (
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func Test(t *testing.T) {
	results := analysistest.Run(t, analysistest.TestData(), Analyzer)
	if len(results) != 1 {
		t.Fatalf("Unexpected number of results (%d)", len(results))
	}

	result := results[0]
	if result.Err != nil {
		t.Fatalf("Unexpected error %s", result.Err)
	}
}
