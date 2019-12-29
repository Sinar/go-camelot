package camelot

import "testing"

func Test_getMetadata(t *testing.T) {
	type args struct {
		pdfSourcePath string
	}
	tests := []struct {
		name string
		args args
	}{
		{"case #1", args{"testdata/par14sesi2-soalan-BukanLisan-385.pdf"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			getMetadata(tt.args.pdfSourcePath)
		})
	}
}
