package camelot

import (
	"reflect"
	"testing"
)

func TestExtractTablesFromSplitDirectory(t *testing.T) {
	type args struct {
		splitDir string
	}
	tests := []struct {
		name string
		args args
		want DirectoryWithTables
	}{
		// TODO: Add test cases.
		{"case #1", args{"/Users/leow/GOMOD/go-dundocs/splitout/BukanLisan-41-60"}, DirectoryWithTables{[]SplitPDF{
			{
				AbsolutePath: "abc",
				RawOutput:    "def",
			}, {
				AbsolutePath: "123",
				RawOutput:    "456",
			},
		}}},
		{"case #2", args{"/Users/leow/GOMOD/go-dundocs/splitout/SOALAN MULUT (81-100)"}, DirectoryWithTables{[]SplitPDF{
			{
				AbsolutePath: "ABC",
				RawOutput:    "DEF",
			}, {
				AbsolutePath: "@@@@",
				RawOutput:    "!!!!",
			},
		}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ExtractTablesFromSplitDirectory(tt.args.splitDir); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractTablesFromSplitDirectory() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSplitPDF_SaveCSV(t *testing.T) {
	type fields struct {
		AbsolutePath string
		RawOutput    string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//sp := SplitPDF{
			//	AbsolutePath: tt.fields.AbsolutePath,
			//	RawOutput:    tt.fields.RawOutput,
			//}
		})
	}
}

func TestSplitPDF_SaveYAML(t *testing.T) {
	type fields struct {
		AbsolutePath string
		RawOutput    string
	}
	tests := []struct {
		name   string
		fields fields
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			//sp := SplitPDF{
			//	AbsolutePath: tt.fields.AbsolutePath,
			//	RawOutput:    tt.fields.RawOutput,
			//}

		})
	}
}

func Test_detectTablesInPDF(t *testing.T) {
	type args struct {
		absolutePDFPath string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := detectTablesInPDF(tt.args.absolutePDFPath); got != tt.want {
				t.Errorf("detectTablesInPDF() = %v, want %v", got, tt.want)
			}
		})
	}
}
