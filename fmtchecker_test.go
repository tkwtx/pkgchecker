package fmtchecker_test

import (
	"log"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/tkwtx/fmtchecker"
	"golang.org/x/tools/go/analysis/analysistest"
)

func init() {
	if err := fmtchecker.Analyzer.Flags.Set("name", "foo"); err != nil {
		log.Fatal(err)
	}
}

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, fmtchecker.Analyzer, "a")
}
