package report

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/Checkmarx/kics/test"
	"github.com/stretchr/testify/require"
)

func TestPrintASFFReport(t *testing.T) {
	tests := []struct {
		name     string
		caseTest jsonCaseTest
	}{
		{
			name: "asff report",
			caseTest: jsonCaseTest{
				summary:  test.SummaryMock,
				path:     filepath.Join(os.TempDir(), "testdir"),
				filename: "output",
			},
		},
		{
			name: "asff report",
			caseTest: jsonCaseTest{
				summary:  test.SummaryMockCritical,
				path:     filepath.Join(os.TempDir(), "testdir"),
				filename: "output2",
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := os.MkdirAll(test.caseTest.path, os.ModePerm); err != nil {
				t.Fatal(err)
			}

			err := PrintASFFReport(test.caseTest.path, test.caseTest.filename, test.caseTest.summary)
			require.NoError(t, err)

			require.FileExists(t, filepath.Join(test.caseTest.path, "asff-"+test.caseTest.filename+".json"))
			os.RemoveAll(test.caseTest.path)
		})
	}
}
