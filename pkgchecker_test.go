package pkgchecker_test

import (
	"log"
	"testing"

	"github.com/gostaticanalysis/testutil"
	"github.com/tkwtx/pkgchecker"
	"golang.org/x/tools/go/analysis/analysistest"
)

func init() {
	if err := pkgchecker.Analyzer.Flags.Set("name", "foo"); err != nil {
		log.Fatal(err)
	}
}

// TestAnalyzer is a test for Analyzer.
func TestAnalyzer(t *testing.T) {
	testdata := testutil.WithModules(t, analysistest.TestData(), nil)
	analysistest.Run(t, testdata, pkgchecker.Analyzer, "a")
}
