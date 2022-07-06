package remediation

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_GetFixs(t *testing.T) {

	filePath := "assets\\queries\\terraform\\alicloud\\ram_account_password_policy_not_required_symbols\\test\\negative1.tf"

	file1 := &File{
		FilePath:        filePath,
		Line:            1,
		Remediation:     "minimum_password_length = 14",
		RemediationType: "addition",
		SimilarityID:    "f282fa13cf5e4ffd4bbb0ee2059f8d0240edcd2ca54b3bb71633145d961de5ce",
	}

	file2 := &File{
		FilePath:        filePath,
		Line:            5,
		Remediation:     "require_symbols = true",
		RemediationType: "replacement",
		SimilarityID:    "87abbee5d0ec977ba193371c702dca2c040ea902d2e606806a63b66119ff89bc",
	}

	query1 := &Query{
		Files:   []File{*file1},
		QueryID: "a9dfec39-a740-4105-bbd6-721ba163c053",
	}

	query2 := &Query{
		Files:   []File{*file2},
		QueryID: "41a38329-d81b-4be4-aef4-55b2615d3282",
	}

	res := &Result{
		Queries: []Query{*query1, *query2},
	}

	replacement := &Remediation{
		Line:         5,
		Remediation:  "require_symbols = true",
		SimilarityID: "87abbee5d0ec977ba193371c702dca2c040ea902d2e606806a63b66119ff89bc",
		QueryID:      "41a38329-d81b-4be4-aef4-55b2615d3282",
	}

	addition := &Remediation{
		Line:         1,
		Remediation:  "minimum_password_length = 14",
		SimilarityID: "f282fa13cf5e4ffd4bbb0ee2059f8d0240edcd2ca54b3bb71633145d961de5ce",
		QueryID:      "a9dfec39-a740-4105-bbd6-721ba163c053",
	}

	var fix, fix2 Fix

	fix.Replacement = append(fix.Replacement, *replacement)
	fix.Addition = append(fix.Addition, *addition)

	want := make(map[string]interface{})
	want[filePath] = fix

	fix2.Replacement = append(fix2.Replacement, *replacement)
	want2 := make(map[string]interface{})
	want2[filePath] = fix2

	type args struct {
		res     *Result
		include []string
	}

	tests := []struct {
		name                       string
		args                       args
		selectedRemediationsNumber int
		want                       map[string]interface{}
	}{
		{
			name: "include all similarityID",
			args: args{
				res:     res,
				include: []string{"all"},
			},
			want:                       want,
			selectedRemediationsNumber: 2,
		},
		{
			name: "include specific similarityID",
			args: args{
				res:     res,
				include: []string{"87abbee5d0ec977ba193371c702dca2c040ea902d2e606806a63b66119ff89bc"},
			},
			want:                       want2,
			selectedRemediationsNumber: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &Summary{
				SelectedRemediationsNumber:   0,
				ActualRemediationsDoneNumber: 0,
			}
			if got := s.GetFixs(*tt.args.res, tt.args.include); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetFixs() = %v, want %v", got, tt.want)
			}
			require.Equal(t, s.SelectedRemediationsNumber, tt.selectedRemediationsNumber)
		})
	}
}
