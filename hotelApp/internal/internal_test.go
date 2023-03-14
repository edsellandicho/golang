package internal

import (
	"testing"
)

func Test_cook(t *testing.T) {
	type args struct {
		dish string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		//fmt.Println("Run unit test by edsel")
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cook(tt.args.dish)
		})
	}
}
