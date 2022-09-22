package util

import "testing"

func TestZip(t *testing.T) {
	type args struct {
		zipPath   string
		rateBytes int64
		paths     []string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ZipWithRateLimit(tt.args.zipPath, tt.args.rateBytes, tt.args.paths...); (err != nil) != tt.wantErr {
				t.Errorf("ZipWithRateLimit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
